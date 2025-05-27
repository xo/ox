package ox

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/xo/ox/text"
)

// AddHelpFlag recursively adds a `--help` flag for all commands in the command
// tree, copying the command's [CommandHelp.Sort], [CommandHelp.CommandSort],
// and [CommandHelp.MaxDist] settings.
func AddHelpFlag(cmd *Command) error {
	if len(cmd.Commands) == 0 {
		return nil
	}
	var opts []Option
	if help, ok := cmd.Help.(*CommandHelp); ok {
		opts = append(
			opts,
			Sort(help.Sort),
			CommandSort(help.CommandSort),
			MaxDist(help.MaxDist),
		)
	}
	for _, c := range cmd.Commands {
		if err := NewHelpFlag(c, opts...); err != nil {
			return err
		}
		if err := AddHelpFlag(c); err != nil {
			return err
		}
	}
	return nil
}

// NewVersion adds a `version` sub command to the command.
func NewVersion(cmd *Command, opts ...Option) error {
	const special = `version`
	if cmd.SubSpecial(special) != nil {
		return nil
	}
	return cmd.Sub(prepend(
		opts,
		Usage(text.VersionCommandName, fmt.Sprintf(text.VersionCommandUsage, cmd.Name)),
		Banner(fmt.Sprintf(text.VersionCommandBanner, cmd.Name)),
		Special(special),
		Exec(func(ctx context.Context) error {
			c, _ := Ctx(ctx)
			return DefaultVersion(c)
		}),
	)...)
}

// NewVersionFlag adds a `--version` flag to the command, or hooks the command's
// flag with `Special == "hook:version"`.
func NewVersionFlag(cmd *Command, opts ...Option) error {
	const special = `hook:version`
	if g := cmd.FlagSpecial(special); g == nil {
		if cmd.Flag(text.VersionFlagShort, false, true) == nil {
			opts = append(opts, Short(text.VersionFlagShort))
		}
		cmd.Flags = cmd.Flags.Hook(
			text.VersionFlagName,
			text.VersionFlagUsage,
			DefaultVersion,
			append(opts, Special(special))...,
		)
	} else {
		g.Type, g.Def, g.NoArg, g.NoArgDef = HookT, DefaultVersion, true, ""
	}
	return nil
}

// NewHelp adds a `help` sub command to the command.
func NewHelp(cmd *Command, opts ...Option) error {
	const special = `help`
	if cmd.SubSpecial(special) != nil {
		return nil
	}
	return cmd.Sub(prepend(
		opts,
		Usage(text.HelpCommandName, text.HelpCommandUsage),
		Banner(fmt.Sprintf(text.HelpCommandBanner, cmd.RootName())),
		Special(special),
		Exec(func(ctx context.Context, args []string) error {
			c, _ := Ctx(ctx)
			_, _ = cmd.Lookup(args...).HelpContext(c).WriteTo(c.Stdout)
			return ErrExit
		}),
		Option(OnErrContinue),
	)...)
}

// NewHelpFlag adds a `--help` flag to the command, or hooks the command's flag
// with `Special == "hook:help"`.
func NewHelpFlag(cmd *Command, opts ...Option) error {
	const special = `hook:help`
	var err error
	if cmd.Help != nil {
		if help, ok := cmd.Help.(*CommandHelp); ok {
			err = Apply(help, opts...)
		}
	} else {
		cmd.Help, err = NewCommandHelp(cmd, opts...)
	}
	if err != nil {
		return err
	}
	f := func(ctx *Context) error {
		_, _ = cmd.HelpContext(ctx).WriteTo(ctx.Stdout)
		return ErrExit
	}
	if g := cmd.FlagSpecial(special); g == nil {
		if cmd.Flag(text.HelpFlagShort, false, true) == nil {
			opts = append(opts, Short(text.HelpFlagShort))
		}
		cmd.Flags = cmd.Flags.Hook(
			text.HelpFlagName,
			text.HelpFlagUsage,
			f,
			append(opts, Special(special))...,
		)
	} else {
		g.Type, g.Def, g.NoArg, g.NoArgDef = HookT, f, true, ""
	}
	return nil
}

// NewComp adds a `completion` sub command to the command.
func NewComp(cmd *Command, opts ...Option) error {
	const special = `comp`
	if cmd.SubSpecial(special) != nil {
		return nil
	}
	var noDescriptions bool
	// base command
	comp, err := NewCommand(prepend(
		opts,
		Parent(cmd),
		Usage(text.CompCommandName, fmt.Sprintf(text.CompCommandUsage, text.CompCommandAnyShellDesc)),
		Banner(fmt.Sprintf(text.CompCommandBanner, text.CompCommandAnyShellDesc)),
		Special(special),
	)...)
	if err != nil {
		return err
	}
	// add help
	if comp.Help, err = NewCommandHelp(
		comp,
		Banner(fmt.Sprintf(text.CompCommandBanner, cmd.RootName(), text.CompCommandAnyShellDesc)),
	); err != nil {
		return err
	}
	// load shells from templates
	txt, tpl, err := loadTemplates(DefaultCompTemplates)
	if err != nil {
		return err
	}
	// build completion commands
	for _, shell := range slices.Sorted(maps.Keys(txt)) {
		templ, ok := tpl[shell]
		if !ok {
			continue
		}
		// create `completion <shell>` command
		sub, err := NewCommand(
			Parent(comp),
			Usage(shell, fmt.Sprintf(text.CompCommandUsage, shell)),
			Flags().
				Bool(
					text.CompCommandFlagNoDescriptionsName,
					text.CompCommandFlagNoDescriptionsUsage,
					Special(special+`:no-desc`),
					Bind(&noDescriptions),
				),
			Exec(func(ctx context.Context) error {
				c, _ := Ctx(ctx)
				return DefaultCompWrite(c, cmd, noDescriptions, shell, templ)
			}),
			Special(special+`:`+shell),
		)
		if err != nil {
			return err
		}
		var banner string
		if t, ok := txt[shell]; ok && t != "" {
			banner = fmt.Sprintf(t, cmd.RootName(), strings.Join(sub.Tree(), " "))
		} else {
			banner = fmt.Sprintf(text.CompCommandBanner, cmd.RootName(), shell)
		}
		if sub.Help, err = NewCommandHelp(sub, Banner(banner)); err != nil {
			return err
		}
		comp.Commands = append(comp.Commands, sub)
	}
	cmd.Commands = append(cmd.Commands, comp)
	return nil
}

// NewCompFlags adds `--completion-script-<type>` flags to a command, or
// hooking any existing flags with `Special == "hook:comp:<type>"`.
func NewCompFlags(cmd *Command, _ ...Option) error {
	const special = `comp`
	// load shells from templates
	txt, tpl, err := loadTemplates(DefaultCompTemplates)
	if err != nil {
		return err
	}
	noDescriptions := false
	// add `--no-descriptions` flag if not defined
	if cmd.Flag(text.CompCommandFlagNoDescriptionsName, true, len(text.CompCommandFlagNoDescriptionsName) == 1) == nil {
		cmd.Flags = cmd.Flags.Bool(
			text.CompCommandFlagNoDescriptionsName,
			text.CompCommandFlagNoDescriptionsUsage,
			Hidden(true),
			Special(special+`:no-desc`),
			Bind(&noDescriptions),
		)
	}
	// add flags
	for _, shell := range slices.Sorted(maps.Keys(txt)) {
		templ, ok := tpl[shell]
		if !ok {
			continue
		}
		f := func(ctx *Context) error {
			return DefaultCompWrite(ctx, cmd, noDescriptions, shell, templ)
		}
		special := `hook:comp:` + shell
		if g := cmd.FlagSpecial(special); g != nil {
			g.Type, g.Def, g.NoArg, g.NoArgDef, g.Hidden, g.Special = HookT, f, true, "", true, special
		} else {
			cmd.Flags = cmd.Flags.Hook(
				fmt.Sprintf(text.CompFlagName, shell),
				fmt.Sprintf(text.CompFlagUsage, shell),
				f,
				Hidden(true),
				Special(special),
			)
		}
	}
	return nil
}

// CommandHelp is command help.
type CommandHelp struct {
	// Context is the evaluation context.
	Context *Context
	// Command is the target command.
	Command *Command
	// NoBanner when true will not output the command's banner.
	NoBanner bool
	// Banner is the command's banner.
	Banner string
	// NoUsage when true will not output the command's usage.
	NoUsage bool
	// NoSpec when true will not output the command's usage spec.
	NoSpec bool
	// Spec is the command's usage spec.
	Spec string
	// NoCommands when true will not output the command's sub commands.
	NoCommands bool
	// CommandSort when true will sort commands.
	CommandSort bool
	// CommandSections are the command's sub command section names.
	CommandSections []string
	// NoAliases when true will not output the command's aliases.
	NoAliases bool
	// NoExample when true will not output the command's examples.
	NoExample bool
	// Example is the command's examples.
	Example string
	// NoFlags when true will not output the command's flags.
	NoFlags bool
	// Sort when true will sort flags.
	Sort bool
	// Sections are the command's flag section names.
	Sections []string
	// NoFooter when true will not output the command's footer.
	NoFooter bool
	// Footer is the command's footer.
	Footer string
	// Hidden includes hidden commands/flags.
	Hidden bool
	// Deprecated includes deprecated commands/flags.
	Deprecated bool
	// MaxDist is the maximum Levenshtein distance for suggestions.
	MaxDist int
}

// NewCommandHelp creates command help based on the passed options.
func NewCommandHelp(cmd *Command, opts ...Option) (*CommandHelp, error) {
	help := &CommandHelp{
		Command:     cmd,
		CommandSort: true,
	}
	if err := Apply(help, opts...); err != nil {
		return nil, err
	}
	return help, nil
}

// SetContext sets the context on the command help.
func (help *CommandHelp) SetContext(ctx *Context) {
	help.Context = ctx
}

// WriteTo satisfies the [io.WriterTo] interface.
func (help *CommandHelp) WriteTo(w io.Writer) (int64, error) {
	wr := NewStringWriter(w)
	for _, f := range []func(StringWriter){
		help.AddBanner,
		help.AddUsage,
		help.AddAliases,
		help.AddExample,
		help.AddCommands,
		help.AddFlags,
		help.AddFooter,
	} {
		f(wr)
	}
	_, _ = wr.WriteString("\n")
	return int64(wr.Len()), nil
}

// AddBanner adds the command's banner.
func (help *CommandHelp) AddBanner(w StringWriter) {
	if help.NoBanner {
		return
	}
	banner := help.Banner
	if banner == "" {
		banner = help.Command.Name + " " + help.Command.Usage
	}
	if strings.TrimSpace(banner) == "" {
		return
	}
	addBreak(w)
	_, _ = w.WriteString(strings.TrimSpace(banner))
}

// AddUsage adds the command's usage.
func (help *CommandHelp) AddUsage(w StringWriter) {
	if help.NoUsage {
		return
	}
	addBreak(w)
	_, _ = w.WriteString(text.Usage)
	_, _ = w.WriteString(":\n  ")
	_, _ = w.WriteString(strings.Join(help.Command.Tree(), " "))
	if help.NoSpec {
		return
	}
	spec := help.Spec
	if spec == "" {
		var v []string
		if help.Command.Flags != nil && len(help.Command.Flags.Flags) > 0 {
			v = append(v, text.CommandFlagsSpec)
		}
		if len(help.Command.Commands) != 0 {
			v = append(v, text.CommandSubSpec)
		}
		// TODO: better output when arg validation/completion has been set
		spec = strings.Join(append(v, text.CommandArgsSpec), " ")
	}
	if strings.TrimSpace(spec) != "" {
		_, _ = w.WriteString(" ")
		_, _ = w.WriteString(spec)
	}
}

// AddAliases adds the command's aliases.
func (help *CommandHelp) AddAliases(w StringWriter) {
	if help.NoAliases || len(help.Command.Aliases) == 0 {
		return
	}
	addBreak(w)
	_, _ = w.WriteString(text.Aliases)
	_, _ = w.WriteString(":\n  ")
	_, _ = w.WriteString(strings.Join(prepend(help.Command.Aliases, help.Command.Name), ", "))
}

// AddExample adds the command's example.
func (help *CommandHelp) AddExample(w StringWriter) {
	if help.NoExample || strings.TrimSpace(help.Example) == "" {
		return
	}
	addBreak(w)
	_, _ = w.WriteString(text.Examples)
	_, _ = w.WriteString(":\n")
	_, _ = w.WriteString(help.Example)
}

// AddCommands adds the command's sub commands.
func (help *CommandHelp) AddCommands(w StringWriter) {
	if help.NoCommands || len(help.Command.Commands) == 0 {
		return
	}
	// determine commands
	var commands []*Command
	for _, c := range help.Command.Commands {
		if c.Hidden && !help.Hidden || c.Deprecated && !help.Deprecated {
			continue
		}
		commands = append(commands, c)
	}
	switch {
	case len(commands) == 0:
		return
	case help.CommandSort:
		sort.Slice(commands, func(i, j int) bool {
			return commands[i].Name < commands[j].Name
		})
	}
	// split between sections
	sections := make(map[int]string, len(help.CommandSections))
	for i, section := range help.CommandSections {
		sections[i] = section
	}
	i, width, indexes := 0, DefaultCommandWidth, make(map[int][]int)
	for _, c := range commands {
		indexes[c.Section] = append(indexes[c.Section], i)
		width = max(width, DefaultWidth(c.Name))
		if _, ok := sections[c.Section]; !ok {
			sections[c.Section] = text.Commands
		}
		i++
	}
	addBreak(w)
	// write commands
	for j, section := range slices.Sorted(maps.Keys(indexes)) {
		if j != 0 {
			_, _ = w.WriteString("\n\n")
		}
		_, _ = w.WriteString(sections[section])
		_, _ = w.WriteString(":")
		for _, i := range indexes[section] {
			c := commands[i]
			_, _ = w.WriteString("\n  ")
			_, _ = w.WriteString(c.Name)
			_, _ = w.WriteString(strings.Repeat(" ", width-len(c.Name)+2))
			DefaultWrapTo(w, c.Usage, DefaultWrapWidth, width+4)
		}
	}
}

// AddFlags adds the command's flags.
func (help *CommandHelp) AddFlags(w StringWriter) {
	if help.NoFlags || help.Command.Flags == nil || len(help.Command.Flags.Flags) == 0 {
		return
	}
	// determine flags
	var flags []*Flag
	for _, g := range help.Command.Flags.Flags {
		if g.Hidden && !help.Hidden || g.Deprecated && !help.Deprecated {
			continue
		}
		flags = append(flags, g)
	}
	switch {
	case len(flags) == 0:
		return
	case help.Sort:
		sort.Slice(flags, func(i, j int) bool {
			return flags[i].Name < flags[j].Name
		})
	}
	// split between sections
	sections := make(map[int]string, len(help.Sections))
	for i, section := range help.Sections {
		sections[i] = section
	}
	var specs []string
	i, width, indexes, hasShort := 0, DefaultFlagWidth, make(map[int][]int), false
	for _, g := range flags {
		specs = append(specs, g.SpecString())
		indexes[g.Section] = append(indexes[g.Section], i)
		width = max(width, len(specs[i]))
		if _, ok := sections[g.Section]; !ok {
			sections[g.Section] = text.Flags
		}
		hasShort = hasShort || g.Short != ""
		i++
	}
	offset := 2
	if hasShort {
		offset = 10
	}
	addBreak(w)
	// write flags
	for j, section := range slices.Sorted(maps.Keys(indexes)) {
		if j != 0 {
			_, _ = w.WriteString("\n\n")
		}
		_, _ = w.WriteString(sections[section])
		_, _ = w.WriteString(":")
		for _, i := range indexes[section] {
			_, _ = w.WriteString("\n  ")
			g := flags[i]
			switch {
			case hasShort && g.Short == "":
				_, _ = w.WriteString("    ")
			case hasShort:
				_, _ = w.WriteString("-")
				_, _ = w.WriteString(g.Short)
				_, _ = w.WriteString(", ")
			}
			_, _ = w.WriteString("--")
			_, _ = w.WriteString(specs[i])
			_, _ = w.WriteString(strings.Repeat(" ", width-len(specs[i])+2))
			usage := g.Usage
			if g.Def != nil && g.Type != HookT && !g.NoArg {
				if def, err := help.Context.Expand(g.Def); err == nil {
					usage += " " + fmt.Sprintf(text.FlagDefault, def)
				}
			}
			DefaultWrapTo(w, usage, DefaultWrapWidth, width+offset)
		}
	}
}

// AddFooter adds the command's footer.
func (help *CommandHelp) AddFooter(w StringWriter) {
	if help.NoFooter {
		return
	}
	footer := help.Footer
	if footer == "" && len(help.Command.Commands) != 0 {
		footer = fmt.Sprintf(text.Footer, strings.Join(help.Command.Tree(), " "))
	}
	if strings.TrimSpace(footer) == "" {
		return
	}
	addBreak(w)
	_, _ = w.WriteString(footer)
}

// addBreak conditionally adds a section break.
func addBreak(w StringWriter) {
	if w.Len() != 0 {
		_, _ = w.WriteString("\n\n")
	}
}

// loadTemplates loads templates.
func loadTemplates(tpls fs.FS) (map[string]string, map[string]string, error) {
	txt, tpl := make(map[string]string), make(map[string]string)
	err := fs.WalkDir(tpls, ".", func(name string, d fs.DirEntry, err error) error {
		switch {
		case err != nil:
			return err
		case d.IsDir():
			return nil
		}
		f, err := tpls.Open(name)
		if err != nil {
			return err
		}
		buf, err := io.ReadAll(f)
		if err != nil {
			f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		switch ext := filepath.Ext(name); {
		case ext == ".txt":
			txt[strings.TrimSuffix(filepath.Base(name), ".txt")] = string(buf)
		default:
			tpl[strings.TrimSuffix(filepath.Base(name), ext)] = string(buf)
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return txt, tpl, nil
}

// CompDirective is a bit map representing the different behaviors the
// shell can be instructed to have once completions have been provided.
//
// Cribbed from cobra for compatibility purposes.
type CompDirective int

// Comp directives.
//
// Cribbed from cobra for compatibility purposes.
const (
	// CompError indicates an error occurred and completions should be ignored.
	CompError CompDirective = 1 << iota
	// CompNoSpace indicates that the shell should not add a space after the
	// completion even if there is a single completion provided.
	CompNoSpace
	// CompNoFileComp indicates that the shell should not provide file
	// completion even when no completion is provided.
	CompNoFileComp
	// CompFilterFileExt indicates that the provided completions should be used
	// as file extension filters.
	CompFilterFileExt
	// CompFilterDirs indicates that only directory names should be provided in
	// file completion. To request directory names within another directory,
	// the returned completions should specify the directory within which to
	// search. The BashCompSubdirsInDir annotation can be used to obtain the
	// same behavior but only for flags.
	CompFilterDirs
	// CompKeepOrder indicates that the shell should preserve the order in
	// which the completions are provided.
	CompKeepOrder
	// CompDefault indicates to let the shell perform its default behavior
	// after completions have been provided.
	CompDefault CompDirective = 0
)

// String satisfies the [fmt.Stringer] interface.
func (dir CompDirective) String() string {
	var sb strings.Builder
	has(&sb, dir, CompError, "CompError")
	has(&sb, dir, CompNoSpace, "CompNoSpace")
	has(&sb, dir, CompNoFileComp, "CompNoFileComp")
	has(&sb, dir, CompFilterFileExt, "CompFilterFileExt")
	has(&sb, dir, CompFilterDirs, "CompFilterDirs")
	has(&sb, dir, CompKeepOrder, "CompKeepOrder")
	if s := sb.String(); s != "" {
		return s
	}
	return "CompDefault"
}

// Completion is a completion.
type Completion struct {
	Name  string
	Usage string
	Dist  int
}

// NewCompletion creates a completion.
func NewCompletion(name, usage string, dist int) Completion {
	if i := strings.Index(usage, "\n"); i != -1 {
		usage = usage[:i]
	}
	return Completion{
		Name:  strings.TrimSpace(name),
		Usage: strings.TrimSpace(usage),
		Dist:  dist,
	}
}

// StringWriter extends the [io.StringWriter] interface with a Len() method,
// allowing both [strings.Builder] and [bytes.Buffer] to be used when
// generating help output.
type StringWriter interface {
	io.StringWriter
	Len() int
}

// NewStringWriter converts a writer to a [StringWriter], or wraps the writer
// when necessary. Wraps (or directly returns) [os.File], [strings.Builder],
// [bytes.Buffer], or other [io.Writer].
func NewStringWriter(w io.Writer) StringWriter {
	switch x := w.(type) {
	case *os.File:
		return &fileWriter{w: x}
	case *strings.Builder:
		return x
	case *bytes.Buffer:
		return x
	case StringWriter:
		return x
	}
	return &stringWriter{w: w}
}

// fileWriter wraps writing to a file.
type fileWriter struct {
	w *os.File
	n bool
}

// WriteString satisfies the [StringWriter] interface.
func (w *fileWriter) WriteString(s string) (int, error) {
	w.n = w.n || s != ""
	return w.w.WriteString(s)
}

// Len satisfies the [StringWriter] interface.
func (w *fileWriter) Len() int {
	if w.n {
		return 1
	}
	return 0
}

type stringWriter struct {
	w io.Writer
	n bool
}

func (w *stringWriter) WriteString(s string) (int, error) {
	w.n = w.n || s != ""
	return w.w.Write([]byte(s))
}

// Len satisfies the [StringWriter] interface.
func (w *stringWriter) Len() int {
	if w.n {
		return 1
	}
	return 0
}

// has builds a string for a.
func has[T inti | uinti](w StringWriter, a, b T, s string) {
	if a&b != 0 {
		if w.Len() != 0 {
			w.WriteString("|")
		}
		w.WriteString(s)
	}
}

// templates are the embedded completion templates.
//
//go:embed comp/*
var templates embed.FS
