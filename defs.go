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
	"sync"

	"github.com/xo/ox/text"
)

// AddHelpFlag recursively adds a `--help` flag for all commands in the command
// tree, copying the command's [CommandHelp.Sort], [CommandHelp.CommandSort],
// [CommandHelp.MaxDist] and [CommandHelp.Column] settings, and
// [CommandHelp.InheritedFlags] when it is true.
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
			HelpColumn(help.Column),
		)
		if help.InheritedFlags {
			opts = append(opts, InheritedFlags(true))
		}
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
		// complete the path of a command, such as help sub <tab>
		CompFunc(func(ctx *Context, toComplete string) ([]Completion, CompDirective, error) {
			comps, _ := cmd.Lookup(ctx.Args...).CompCommands(toComplete, false, false)
			return comps, CompNoFileComp | CompKeepOrder | CompNoFilter, nil
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
	// the shells wait for their first read (B55)
	comp.lazy = &lazySubs{f: func(comp *Command) error {
		return addCompShells(cmd, comp)
	}}
	cmd.Commands = append(cmd.Commands, comp)
	return nil
}

// addCompShells adds a `completion <shell>` command to comp, the completion
// command of cmd, for each shell. When comp has a help flag, each shell gets
// one too, as [AddHelpFlag] gives it.
func addCompShells(cmd, comp *Command) error {
	const special = `comp`
	txt, tpl, err := compTemplates()
	if err != nil {
		return err
	}
	var noDescriptions bool
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
	if comp.FlagSpecial("hook:help") != nil {
		return AddHelpFlag(comp)
	}
	return nil
}

// NewCompFlags adds `--completion-script-<type>` flags to a command, or
// hooking any existing flags with `Special == "hook:comp:<type>"`.
func NewCompFlags(cmd *Command, _ ...Option) error {
	const special = `comp`
	// load shells from templates
	txt, tpl, err := compTemplates()
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
	// InheritedFlags when true writes the flags that the command inherits
	// from its parents, in a section of their own. See [InheritedFlags].
	InheritedFlags bool
	// Args are the names and the usage of the command's positional
	// arguments. See [ArgUsage].
	Args []HelpLine
	// Column is the column where each description starts. See [HelpColumn].
	Column int
	// NoFooter when true will not output the command's footer.
	NoFooter bool
	// Footer is the command's footer.
	Footer string
	// Hidden includes hidden commands/flags.
	Hidden bool
	// Deprecated includes deprecated commands/flags.
	Deprecated bool
	// MaxDist is the maximum Levenshtein distance for suggestions. 0 means
	// [DefaultMaxDist], and a negative distance turns off the distance match.
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
		help.AddArgs,
		help.AddAliases,
		help.AddExample,
		help.AddCommands,
		help.AddFlags,
		help.AddInheritedFlags,
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
	if spec := help.SpecString(); strings.TrimSpace(spec) != "" {
		_, _ = w.WriteString(" ")
		_, _ = w.WriteString(spec)
	}
}

// SpecString returns the usage spec of the command, which help writes after
// the names of the command and its parents, such as `[flags] [command]
// [args]`. It is [CommandHelp.Spec], or the spec of the flags, the sub
// commands and the arguments of the command.
func (help *CommandHelp) SpecString() string {
	if help.Spec != "" {
		return help.Spec
	}
	var v []string
	if help.Command.Flags != nil && len(help.Command.Flags.Flags) > 0 {
		v = append(v, text.CommandFlagsSpec)
	}
	if help.Command.hasSubs() {
		v = append(v, text.CommandSubSpec)
	}
	switch cmd := help.Command; {
	case len(cmd.Positional) != 0:
		for _, a := range cmd.Positional {
			v = append(v, a.SpecString())
		}
	case cmd.noArgs:
		// ValidArgs(0, 0) allows no argument
	case len(cmd.ValidArgs) != 0:
		// the values of ValidArgs, with ... when more than one is allowed
		arg := fmt.Sprintf(text.CommandArgValuesSpec, strings.Join(cmd.ValidArgs, "|"))
		if cmd.maxArgs != 1 {
			arg += text.CommandArgsMore
		}
		v = append(v, arg)
	default:
		v = append(v, text.CommandArgsSpec)
	}
	return strings.Join(v, " ")
}

// AddArgs adds the usage of the command's positional arguments: each
// argument of [Command.Positional], with its default, then each line of
// [CommandHelp.Args].
func (help *CommandHelp) AddArgs(w StringWriter) {
	args := make([]HelpLine, 0, len(help.Command.Positional)+len(help.Args))
	for _, a := range help.Command.Positional {
		args = append(args, HelpLine{Name: a.name(), Usage: help.FlagUsage(a.Flag)})
	}
	args = append(args, help.Args...)
	if len(args) == 0 {
		return
	}
	width := DefaultCommandWidth
	for _, a := range args {
		width = max(width, DefaultWidth(a.Name))
	}
	addBreak(w)
	_, _ = w.WriteString(text.Arguments)
	_, _ = w.WriteString(":")
	for _, a := range args {
		_, _ = w.WriteString("\n  ")
		if help.Column > 0 {
			writeColumn(w, a.Name, a.Usage, help.Column)
			continue
		}
		_, _ = w.WriteString(a.Name)
		_, _ = w.WriteString(strings.Repeat(" ", width-DefaultWidth(a.Name)+2))
		DefaultWrapTo(w, a.Usage, DefaultWrapWidth, width+4)
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
	if help.NoCommands || len(help.Command.subs()) == 0 {
		return
	}
	// determine commands
	var commands []*Command
	for _, c := range help.Command.subs() {
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
			if help.Column > 0 {
				writeColumn(w, c.Name, c.Usage, help.Column)
				continue
			}
			_, _ = w.WriteString(c.Name)
			_, _ = w.WriteString(strings.Repeat(" ", width-len(c.Name)+2))
			DefaultWrapTo(w, c.Usage, DefaultWrapWidth, width+4)
		}
	}
}

// AddFlags adds the command's flags.
func (help *CommandHelp) AddFlags(w StringWriter) {
	if help.NoFlags || help.Command.Flags == nil {
		return
	}
	help.writeFlags(w, help.Command.Flags.Flags, help.Sections, true)
}

// AddInheritedFlags adds the flags that the command inherits from its
// parents, in a section of their own, when [CommandHelp.InheritedFlags] is
// true.
func (help *CommandHelp) AddInheritedFlags(w StringWriter) {
	if help.NoFlags || !help.InheritedFlags {
		return
	}
	help.writeFlags(w, help.Command.inheritedFlags(), []string{text.GlobalFlags}, false)
}

// AddFooter adds the command's footer.
func (help *CommandHelp) AddFooter(w StringWriter) {
	if help.NoFooter {
		return
	}
	footer := help.Footer
	if footer == "" && help.Command.hasSubs() {
		footer = fmt.Sprintf(text.Footer, strings.Join(help.Command.Tree(), " "))
	}
	if strings.TrimSpace(footer) == "" {
		return
	}
	addBreak(w)
	_, _ = w.WriteString(footer)
}

// FlagUsage returns the usage of the flag that help writes, with the text of
// its default value unless [Flag.NoDefault] is set. A required flag gets the
// text of [text.FlagRequired], and not its default, because a default does
// not count as a value. [CommandHelp.Context] expands the default. Without a
// context, the default is as the program wrote it, such as `$HOME/.apprc`.
func (help *CommandHelp) FlagUsage(g *Flag) string {
	if g.Required {
		return g.Usage + " " + text.FlagRequired
	}
	if g.Def == nil || g.Type == HookT || g.NoArg || g.NoDefault {
		return g.Usage
	}
	def, err := help.Context.Expand(g.Def)
	if err != nil {
		return g.Usage
	}
	// a default that is not a string, such as Default(3)
	s, err := asString[string](def)
	if err != nil {
		s = fmt.Sprint(def)
	}
	if g.Home {
		s = homePath(s)
	}
	return g.Usage + " " + fmt.Sprintf(text.FlagDefault, s)
}

// writeFlags writes the flags that help shows. When sections is true, each
// flag goes in its section, which takes its name from names. When it is
// false, every flag goes in one section, with the name names[0].
func (help *CommandHelp) writeFlags(w StringWriter, all []*Flag, names []string, sections bool) {
	// determine flags
	var flags []*Flag
	for _, g := range all {
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
	sectionNames := make(map[int]string, len(names))
	for i, section := range names {
		sectionNames[i] = section
	}
	var specs, negs []string
	prefixes := help.Command.negatePrefixes()
	i, width, indexes, hasShort, hasPlus := 0, DefaultFlagWidth, make(map[int][]int), false, false
	for _, g := range flags {
		section := 0
		if sections {
			section = g.Section
		}
		// a negated form shows before the name, such as --[no-]debug
		var neg string
		if v := g.negations(prefixes); len(v) != 0 {
			neg = fmt.Sprintf(text.FlagNegatedSpec, v[0])
		}
		negs = append(negs, neg)
		specs = append(specs, neg+g.SpecString())
		indexes[section] = append(indexes[section], i)
		width = max(width, len(specs[i]))
		for _, v := range g.HelpValues {
			width = max(width, len(g.Name)+1+len(v.Name))
		}
		if _, ok := sectionNames[section]; !ok {
			sectionNames[section] = text.Flags
		}
		hasShort = hasShort || g.Short != ""
		hasPlus = hasPlus || g.AllowPlus
		i++
	}
	// the short column holds `-x, `, and `-x, +x, ` when a flag has a plus
	// form (D52)
	offset, shortCol := 2, ""
	switch {
	case hasPlus:
		offset, shortCol = 14, "        "
	case hasShort:
		offset, shortCol = 10, "    "
	}
	addBreak(w)
	// write flags
	for j, section := range slices.Sorted(maps.Keys(indexes)) {
		if j != 0 {
			_, _ = w.WriteString("\n\n")
		}
		_, _ = w.WriteString(sectionNames[section])
		_, _ = w.WriteString(":")
		for _, i := range indexes[section] {
			_, _ = w.WriteString("\n  ")
			g := flags[i]
			if help.Column > 0 {
				writeColumn(w, columnNames(g, negs[i], hasShort), help.FlagUsage(g), help.Column)
			} else {
				switch {
				case g.Short == "":
					_, _ = w.WriteString(shortCol)
				case g.AllowPlus:
					_, _ = w.WriteString("-" + g.Short + ", +" + g.Short + ", ")
				default:
					_, _ = w.WriteString("-" + g.Short + ", ")
					_, _ = w.WriteString(shortCol[4:])
				}
				_, _ = w.WriteString("--")
				_, _ = w.WriteString(specs[i])
				_, _ = w.WriteString(strings.Repeat(" ", width-len(specs[i])+2))
				DefaultWrapTo(w, help.FlagUsage(g), DefaultWrapWidth, width+offset)
			}
			help.addValues(w, g, shortCol, width, offset)
		}
	}
}

// addValues writes the values of [Flag.HelpValues] on lines of their own,
// such as `--help=commands`, in the layout of the flag. indent is the blank
// short column.
func (help *CommandHelp) addValues(w StringWriter, g *Flag, indent string, width, offset int) {
	for _, v := range g.HelpValues {
		spec := g.Name + "=" + v.Name
		_, _ = w.WriteString("\n  ")
		if help.Column > 0 {
			writeColumn(w, indent+"--"+spec, v.Usage, help.Column)
			continue
		}
		_, _ = w.WriteString(indent)
		_, _ = w.WriteString("--")
		_, _ = w.WriteString(spec)
		_, _ = w.WriteString(strings.Repeat(" ", width-len(spec)+2))
		DefaultWrapTo(w, v.Usage, DefaultWrapWidth, width+offset)
	}
}

// columnNames returns the names of the flag for the layout of [HelpColumn],
// such as `-v, --set=, --variable=NAME=VALUE`. The short names come first,
// then the long name and the long aliases. Each long name but the last ends
// with the spacer of [text.FlagSpecSpacer] without spaces, and the last one
// ends with the spec. neg is the spec of the negated form of the flag, such
// as `[no-]`, which goes before the name.
func columnNames(g *Flag, neg string, hasShort bool) string {
	var shorts, longs []string
	if g.Short != "" {
		shorts = append(shorts, g.Short)
	}
	longs = append(longs, neg+g.Name)
	for _, s := range g.Aliases {
		if len(s) == 1 {
			shorts = append(shorts, s)
		} else {
			longs = append(longs, s)
		}
	}
	var sb strings.Builder
	if hasShort && len(shorts) == 0 {
		_, _ = sb.WriteString("    ")
	}
	for _, s := range shorts {
		_, _ = sb.WriteString("-")
		_, _ = sb.WriteString(s)
		_, _ = sb.WriteString(", ")
		if g.AllowPlus {
			_, _ = sb.WriteString("+")
			_, _ = sb.WriteString(s)
			_, _ = sb.WriteString(", ")
		}
	}
	arg := g.specArg()
	for i, s := range longs {
		_, _ = sb.WriteString("--")
		_, _ = sb.WriteString(s)
		if i == len(longs)-1 {
			_, _ = sb.WriteString(arg)
			break
		}
		if arg != "" {
			_, _ = sb.WriteString(strings.TrimSpace(text.FlagSpecSpacer))
		}
		_, _ = sb.WriteString(", ")
	}
	return sb.String()
}

// writeColumn writes the names and the usage of a flag or a command for the
// layout of [HelpColumn]. The caller writes the indent of 2 spaces before the
// names. When the names do not leave 2 spaces before the column, the usage
// starts on the next line. Each line of the usage starts at the column.
func writeColumn(w StringWriter, names, usage string, column int) {
	_, _ = w.WriteString(names)
	if n := column - 2 - DefaultWidth(names); n >= 2 {
		_, _ = w.WriteString(strings.Repeat(" ", n))
	} else {
		_, _ = w.WriteString("\n")
		_, _ = w.WriteString(strings.Repeat(" ", column))
	}
	for i, line := range strings.Split(usage, "\n") {
		if i != 0 {
			_, _ = w.WriteString("\n")
			if line != "" {
				_, _ = w.WriteString(strings.Repeat(" ", column))
			}
		}
		_, _ = w.WriteString(line)
	}
}

// addBreak conditionally adds a section break.
func addBreak(w StringWriter) {
	if w.Len() != 0 {
		_, _ = w.WriteString("\n\n")
	}
}

// compTemplates returns the completion templates of [DefaultCompTemplates].
// It reads the embedded templates once, and reuses them, because each
// command with completion asks for them when it is built. The maps are only
// read. A program that sets DefaultCompTemplates to its own files gets them
// read each time.
func compTemplates() (map[string]string, map[string]string, error) {
	if DefaultCompTemplates == templates {
		t := embeddedTemplates()
		return t.txt, t.tpl, t.err
	}
	return loadTemplates(DefaultCompTemplates)
}

// compTemplateSet is the result of [loadTemplates].
type compTemplateSet struct {
	txt, tpl map[string]string
	err      error
}

// embeddedTemplates reads the embedded completion templates once.
var embeddedTemplates = sync.OnceValue(func() compTemplateSet {
	txt, tpl, err := loadTemplates(templates)
	return compTemplateSet{txt, tpl, err}
})

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
			_ = f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		switch ext := filepath.Ext(name); ext {
		case ".txt":
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
	// CompNoFilter indicates that ox does not keep only the completions that
	// start with the word to complete. A [CompFunc] returns it. ox does not
	// pass it to the shell.
	CompNoFilter
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
	has(&sb, dir, CompNoFilter, "CompNoFilter")
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

// NewActiveHelp creates a line of ActiveHelp, which a [CompFunc] returns
// beside its completions. The completion scripts show it as help, as cobra's
// scripts do. The environment variable of
// [DefaultCompActiveHelpSuffix], such as APP_ACTIVE_HELP=0, turns it off.
func NewActiveHelp(help string) Completion {
	return Completion{Name: DefaultCompActiveHelpMarker + help}
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

// fileWriter wraps writing to a [os.File].
type fileWriter struct {
	w *os.File
	b bool
}

// WriteString satisfies the [StringWriter] interface.
func (w *fileWriter) WriteString(s string) (int, error) {
	w.b = w.b || s != ""
	return w.w.WriteString(s)
}

// Len satisfies the [StringWriter] interface.
func (w *fileWriter) Len() int {
	if w.b {
		return 1
	}
	return 0
}

// stringWriter waps a [io.Writer] with WriteString.
type stringWriter struct {
	w io.Writer
	b bool
}

// WriteString satisfies the [StringWriter] interface.
func (w *stringWriter) WriteString(s string) (int, error) {
	w.b = w.b || s != ""
	return w.w.Write([]byte(s))
}

// Len satisfies the [StringWriter] interface.
func (w *stringWriter) Len() int {
	if w.b {
		return 1
	}
	return 0
}

// has builds a string for a.
func has[T inti | uinti](w StringWriter, a, b T, s string) {
	if a&b != 0 {
		if w.Len() != 0 {
			_, _ = w.WriteString("|")
		}
		_, _ = w.WriteString(s)
	}
}

// templates are the embedded completion templates.
//
//go:embed comp/*
var templates embed.FS
