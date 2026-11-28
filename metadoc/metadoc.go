// Package metadoc writes the docs of an ox command tree, as a man page or as
// Markdown. It uses only the standard library, so it is in the module of ox.
// A program that does not import it does not link it, and an import changes
// nothing.
//
// A program adds the hidden command `__docs` with [Command]:
//
//	ox.Run(
//		ox.Usage("app", "an app"),
//		ox.Defaults(),
//		metadoc.Command(metadoc.ManPage(), metadoc.Markdown()),
//		// ...
//	)
//
// Then `app __docs man` writes the whole tree as one man page, and
// `app __docs markdown` writes it as one Markdown document. With
// `--dir DIR`, each writes one file for each command in the folder, such as
// `app.1` and `app-sub.1`, or `app.md` and `app-sub.md`.
//
// A program that keeps the writer out of its binary calls [Write],
// [WriteCommand] or [WriteDir] from a test or from `go generate`.
//
// The docs leave out each hidden command and hidden flag, each deprecated
// command and deprecated flag, and the `__docs` command. The
// [ox.CommandHelp] of a command with [ox.CommandHelp.Hidden] or
// [ox.CommandHelp.Deprecated] keeps them, as help does. The docs keep the
// help, version and completion commands. The docs show each default as the
// program wrote it, such as `$HOME/.apprc`, so the docs are the same on each
// computer.
package metadoc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"iter"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/xo/ox"
	"github.com/xo/ox/text"
)

// DefaultName is the name of the hidden command that [Command] adds.
var DefaultName = "__docs"

// special is the special value of the hidden command. ox keeps the command
// with this value out of help, suggestions and completion.
const special = "docs"

// Format is a format of the docs, from [ManPage] or [Markdown].
type Format struct {
	// name is the name of the sub command of the hidden command, such as
	// "man".
	name string
	// w writes the format.
	w writer
}

// writer writes the docs in one format.
type writer interface {
	// ext returns the extension of a file, such as ".1".
	ext() string
	// tree writes the page of root and of each command below it as one
	// document.
	tree(w io.Writer, root *page) error
	// page writes the page of one command, with links to its parent and to
	// its sub commands.
	page(w io.Writer, p *page) error
}

// Command is a root [ox.CommandOption] that adds the hidden command
// `__docs`, with a sub command for each format, such as `__docs man` and
// `__docs markdown`. With no format, it adds [ManPage] and [Markdown]. Each
// sub command writes the whole tree to standard out, as one document. With
// `--dir DIR`, it writes a file for each command in the folder, as
// [WriteDir] does.
//
// Command adds the command after each other option, so the help, version
// and completion of [ox.Defaults] do not change. It returns an error that
// wraps [ox.ErrExist] when the program has a command named `__docs`, and an
// error that wraps [ox.ErrInvalidValue] when two formats have the same name.
func Command(formats ...Format) ox.CommandOption {
	return ox.Extend("metadoc.Command", func(root *ox.Command) error {
		fs := formats
		if len(fs) == 0 {
			fs = []Format{ManPage(), Markdown()}
		}
		if root.Command(DefaultName) != nil {
			return fmt.Errorf("command %s: %w", DefaultName, ox.ErrExist)
		}
		docs, err := ox.NewCommand(
			ox.Parent(root),
			ox.Usage(DefaultName, text.DocCommandUsage),
			ox.Hidden(true),
			ox.Special(special),
		)
		if err != nil {
			return err
		}
		for _, f := range fs {
			if f.w == nil {
				return fmt.Errorf("%w: empty format", ox.ErrInvalidValue)
			}
			if docs.Command(f.name) != nil {
				return fmt.Errorf("%w: format %s is given twice", ox.ErrInvalidValue, f.name)
			}
			sub, err := newFormatCommand(docs, f)
			if err != nil {
				return err
			}
			docs.Commands = append(docs.Commands, sub)
		}
		// the hidden command gets a help flag when the root has one
		if root.FlagSpecial("hook:help") != nil {
			if err := ox.NewHelpFlag(docs); err != nil {
				return err
			}
			if err := ox.AddHelpFlag(docs); err != nil {
				return err
			}
		}
		root.Commands = append(root.Commands, docs)
		return nil
	})
}

// newFormatCommand creates the sub command of the hidden command for the
// format.
func newFormatCommand(docs *ox.Command, f Format) (*ox.Command, error) {
	var dir string
	return ox.NewCommand(
		ox.Parent(docs),
		ox.Usage(f.name, fmt.Sprintf(text.DocFormatUsage, f.name)),
		ox.Special(special+":"+f.name),
		ox.Flags().
			String("dir", text.DocDirFlagUsage, ox.Spec("DIR"), ox.Bind(&dir)),
		ox.Exec(func(ctx context.Context) error {
			c, _ := ox.Ctx(ctx)
			if dir != "" {
				return WriteDir(dir, c.Root, f)
			}
			return Write(c.Stdout, c.Root, f)
		}),
	)
}

// Write writes the docs of cmd and of each command below it to w, as one
// document: one man page, or one Markdown document. The page of cmd comes
// first, then a section that holds each command below it.
func Write(w io.Writer, cmd *ox.Command, format Format) error {
	if format.w == nil {
		return fmt.Errorf("%w: empty format", ox.ErrInvalidValue)
	}
	return format.w.tree(w, build(cmd, nil))
}

// WriteCommand writes the page of cmd to w, with references to the pages of
// its parent and of its sub commands, as [WriteDir] writes each file.
func WriteCommand(w io.Writer, cmd *ox.Command, format Format) error {
	if format.w == nil {
		return fmt.Errorf("%w: empty format", ox.ErrInvalidValue)
	}
	var parent *page
	if cmd.Parent != nil {
		parent = &page{cmd: cmd.Parent, tree: cmd.Parent.Tree(), usage: cmd.Parent.Usage}
	}
	return format.w.page(w, build(cmd, parent))
}

// WriteDir writes a file for cmd and for each command below it in dir. The
// name of a file is the names of the command and its parents, joined with
// `-`, and the extension of the format, such as `app-sub.1` or `app-sub.md`.
// Each file refers to the files of the parent and of the sub commands.
// WriteDir creates dir when it does not exist, and writes each file with
// the mode 0o644. It returns an error that wraps [ox.ErrInvalidValue], and
// writes no file, when the name of a file is not a name in dir, such as a
// name with a `/`.
func WriteDir(dir string, cmd *ox.Command, format Format) error {
	if format.w == nil {
		return fmt.Errorf("%w: empty format", ox.ErrInvalidValue)
	}
	root := build(cmd, nil)
	var pages []*page
	for p := range root.all() {
		if name := p.file(format.w.ext()); !isLocalName(name) {
			return fmt.Errorf("%w: file name %q is not in the folder", ox.ErrInvalidValue, name)
		}
		pages = append(pages, p)
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, p := range pages {
		var buf bytes.Buffer
		if err := format.w.page(&buf, p); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, p.file(format.w.ext())), buf.Bytes(), 0o644); err != nil {
			return err
		}
	}
	return nil
}

// isLocalName returns true when name is the name of a file in a folder, and
// not a path.
func isLocalName(name string) bool {
	return name != "" &&
		filepath.IsLocal(name) &&
		filepath.Base(name) == name &&
		!strings.ContainsAny(name, `/\`)
}

// page is the docs of a command, from the data of its help.
type page struct {
	// cmd is the command.
	cmd *ox.Command
	// tree is the names of the command and its parents, from the root.
	tree []string
	// usage is the usage of the command.
	usage string
	// spec is the usage spec of the command, such as `[flags] [args]`.
	spec string
	// banner is the banner of the command, or its usage.
	banner string
	// flags are the flags of the command.
	flags []flagDoc
	// global are the flags that the command inherits from its parents.
	global []flagDoc
	// args are the positional arguments of the command.
	args []ox.HelpLine
	// example is the example text of the command.
	example string
	// footer is the footer of the command.
	footer string
	// parent is the page of the parent, or nil.
	parent *page
	// subs are the pages of the sub commands.
	subs []*page
}

// flagDoc is a flag of a page.
type flagDoc struct {
	// names are the names of the flag, such as `-v` and `--verbose`.
	names []string
	// arg is the spec after the last name, such as ` FILE`.
	arg string
	// usage is the usage of the flag, with its default and its allowed
	// values.
	usage string
	// values are the values of [ox.HelpValue], such as `--help=commands`.
	values []ox.HelpLine
}

// title returns the names of the command and its parents, joined with
// spaces, such as `app sub`.
func (p *page) title() string {
	return strings.Join(p.tree, " ")
}

// name returns the name of the page, such as `app-sub`.
func (p *page) name() string {
	return strings.Join(p.tree, "-")
}

// file returns the file name of the page, with the extension.
func (p *page) file(ext string) string {
	return p.name() + ext
}

// root returns the name and the version of the root of the command.
func (p *page) root() (string, string) {
	c := p.cmd
	for c.Parent != nil {
		c = c.Parent
	}
	return c.Name, c.Version
}

// all returns an iterator over the page and each page below it, in order.
func (p *page) all() iter.Seq[*page] {
	return func(yield func(*page) bool) {
		p.walk(yield)
	}
}

// walk calls yield with the page and each page below it, until yield
// returns false.
func (p *page) walk(yield func(*page) bool) bool {
	if !yield(p) {
		return false
	}
	for _, sub := range p.subs {
		if !sub.walk(yield) {
			return false
		}
	}
	return true
}

// build builds the page of cmd and of each command below it.
func build(cmd *ox.Command, parent *page) *page {
	help := helpOf(cmd)
	p := &page{
		cmd:    cmd,
		tree:   cmd.Tree(),
		usage:  strings.TrimSpace(cmd.Usage),
		parent: parent,
	}
	if !help.NoSpec {
		p.spec = help.SpecString()
	}
	if !help.NoBanner {
		if p.banner = strings.TrimSpace(help.Banner); p.banner == "" {
			p.banner = p.usage
		}
	}
	if !help.NoFlags {
		var own []*ox.Flag
		if cmd.Flags != nil {
			own = cmd.Flags.Flags
		}
		p.flags = newFlags(help, own)
		p.global = newFlags(help, inherited(cmd))
	}
	for _, a := range cmd.Positional {
		name := a.Flag.Spec
		if name == "" {
			name = a.Flag.Name
		}
		p.args = append(p.args, ox.HelpLine{Name: name, Usage: flagUsage(help, a.Flag)})
	}
	p.args = append(p.args, help.Args...)
	if !help.NoExample {
		p.example = strings.TrimRight(help.Example, " \t\n")
	}
	if !help.NoFooter {
		p.footer = strings.TrimSpace(help.Footer)
	}
	for _, c := range subs(cmd, help) {
		p.subs = append(p.subs, build(c, p))
	}
	return p
}

// helpOf returns a copy of the help of the command without a context, so
// that no default is expanded.
func helpOf(cmd *ox.Command) *ox.CommandHelp {
	if help, ok := cmd.Help.(*ox.CommandHelp); ok {
		h := *help
		h.Context = nil
		return &h
	}
	return &ox.CommandHelp{Command: cmd, CommandSort: true}
}

// subs returns the sub commands of cmd that the docs show.
func subs(cmd *ox.Command, help *ox.CommandHelp) []*ox.Command {
	// a method of the command builds the sub commands that wait for their
	// first read, such as the shells of completion
	_ = cmd.Command("")
	v := make([]*ox.Command, 0, len(cmd.Commands))
	for _, c := range cmd.Commands {
		switch {
		case c.Special == special,
			c.Hidden && !help.Hidden,
			c.Deprecated && !help.Deprecated:
			continue
		}
		v = append(v, c)
	}
	if help.CommandSort {
		slices.SortStableFunc(v, func(a, b *ox.Command) int {
			return strings.Compare(a.Name, b.Name)
		})
	}
	return v
}

// inherited returns the flags that cmd inherits from its parents, from the
// nearest parent first. A local flag of a parent is not inherited, and a
// flag with the name of a nearer flag is hidden by it.
func inherited(cmd *ox.Command) []*ox.Flag {
	seen := make(map[string]bool)
	if cmd.Flags != nil {
		for _, g := range cmd.Flags.Flags {
			seen[g.Name] = true
		}
	}
	var flags []*ox.Flag
	for c := cmd.Parent; c != nil; c = c.Parent {
		if c.Flags == nil {
			continue
		}
		for _, g := range c.Flags.Flags {
			if g.Local || seen[g.Name] {
				continue
			}
			seen[g.Name] = true
			flags = append(flags, g)
		}
	}
	return flags
}

// newFlags returns the flags of a page, without the hidden and the
// deprecated flags that help leaves out.
func newFlags(help *ox.CommandHelp, all []*ox.Flag) []flagDoc {
	var flags []*ox.Flag
	for _, g := range all {
		if g.Hidden && !help.Hidden || g.Deprecated && !help.Deprecated {
			continue
		}
		flags = append(flags, g)
	}
	if help.Sort {
		slices.SortStableFunc(flags, func(a, b *ox.Flag) int {
			return strings.Compare(a.Name, b.Name)
		})
	}
	v := make([]flagDoc, 0, len(flags))
	for _, g := range flags {
		var shorts, longs []string
		if g.Short != "" {
			shorts = append(shorts, "-"+g.Short)
		}
		// a negated form shows before the name, as help writes it, such as
		// --[no-]debug
		long := "--" + g.Name
		if negs := help.Command.NegatedNames(g); len(negs) != 0 {
			long = "--[" + strings.TrimSuffix(negs[0], g.Name) + "]" + g.Name
		}
		longs = append(longs, long)
		for _, s := range g.Aliases {
			if utf8.RuneCountInString(s) == 1 {
				shorts = append(shorts, "-"+s)
			} else {
				longs = append(longs, "--"+s)
			}
		}
		f := flagDoc{
			names: append(shorts, longs...),
			arg:   strings.TrimPrefix(g.SpecString(), g.Name),
			usage: flagUsage(help, g),
		}
		for _, val := range g.HelpValues {
			f.values = append(f.values, ox.HelpLine{Name: "--" + g.Name + "=" + val.Name, Usage: val.Usage})
		}
		v = append(v, f)
	}
	return v
}

// flagUsage returns the usage of the flag, as help writes it, with its
// allowed values.
func flagUsage(help *ox.CommandHelp, g *ox.Flag) string {
	usage := help.FlagUsage(g)
	if len(g.Allowed) != 0 {
		values := make([]string, len(g.Allowed))
		for i, val := range g.Allowed {
			values[i] = fmt.Sprint(val)
		}
		usage += " " + fmt.Sprintf(text.DocAllowed, strings.Join(values, ", "))
	}
	return strings.TrimSpace(usage)
}

// spec returns the names of the flag joined with `, `, and its spec.
func (f flagDoc) spec() string {
	return strings.Join(f.names, ", ") + f.arg
}
