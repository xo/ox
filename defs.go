package ox

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/xo/ox/text"
)

// AddHelp recursively adds help for all sub commands on the command, copying
// settings selectively.
func AddHelp(cmd *Command) error {
	if len(cmd.Commands) == 0 {
		return nil
	}
	sort, commandSort, minDist := false, true, 0
	if help, ok := cmd.Help.(*CommandHelp); ok {
		sort, commandSort, minDist = help.Sort, help.CommandSort, help.MinDist
	}
	for _, c := range cmd.Commands {
		if err := NewHelpFlag(c, Sort(sort), CommandSort(commandSort), MinDist(minDist)); err != nil {
			return err
		}
		if err := AddHelp(c); err != nil {
			return err
		}
	}
	return nil
}

// NewVersion adds a `version` sub command to the command.
func NewVersion(cmd *Command, opts ...Option) error {
	return cmd.Sub(prepend(
		opts,
		Usage(text.VersionCommandName, text.VersionCommandDesc),
		Special(`version`),
		Exec(func(ctx context.Context) error {
			c, _ := Ctx(ctx)
			_ = DefaultVersion(c)
			return ErrExit
		}),
	)...)
}

// NewVersionFlag adds a `--version` flag to the command, or hooks the command's
// flag with `Special == "hook:version"`.
func NewVersionFlag(cmd *Command, opts ...Option) error {
	if g := cmd.FlagSpecial("hook:version"); g != nil {
		g.Type, g.Def, g.NoArg, g.NoArgDef = HookT, DefaultVersion, true, ""
	} else {
		if cmd.Flag(text.VersionFlagShort, false, true) == nil {
			opts = append(opts, Short(text.VersionFlagShort))
		}
		cmd.Flags = cmd.Flags.Hook(text.VersionFlagName, text.VersionFlagDesc, DefaultVersion, opts...)
	}
	return nil
}

// NewHelp adds a `help` sub command to the command.
func NewHelp(cmd *Command, opts ...Option) error {
	return cmd.Sub(prepend(
		opts,
		Usage(text.HelpCommandName, text.HelpCommandDesc),
		Special(`help`),
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
	var err error
	if cmd.Help, err = NewCommandHelp(cmd, opts...); err != nil {
		return err
	}
	f := func(ctx *Context) error {
		_, _ = cmd.HelpContext(ctx).WriteTo(ctx.Stdout)
		return ErrExit
	}
	if g := cmd.FlagSpecial("hook:help"); g != nil {
		g.Type, g.Def, g.NoArg, g.NoArgDef = HookT, f, true, ""
	} else {
		var opts []Option
		if cmd.Flag(text.HelpFlagShort, false, true) == nil {
			opts = append(opts, Short(text.HelpFlagShort))
		}
		cmd.Flags = cmd.Flags.Hook(text.HelpFlagName, text.HelpFlagDesc, f, opts...)
	}
	return nil
}

// NewComp adds a `completion` sub command to the command.
func NewComp(cmd *Command, opts ...Option) error {
	tpls := cmd.Comp
	if tpls == nil {
		tpls = templates
	}
	// base command
	comp, err := NewCommand(prepend(
		opts,
		Parent(cmd),
		Usage(text.CompCommandName, fmt.Sprintf(text.CompDesc, text.CompCommandAShellDesc)),
		Special("comp"),
	)...)
	if err != nil {
		return err
	}
	// add help
	if comp.Help, err = NewCommandHelp(
		comp,
		Banner(fmt.Sprintf(text.CompCommandBanner, cmd.RootName(), text.CompCommandAShellDesc)),
	); err != nil {
		return err
	}
	// load shells from templates
	txt, tpl, err := loadTemplates(tpls)
	if err != nil {
		return err
	}
	// build completion commands
	for _, shell := range slices.Sorted(maps.Keys(txt)) {
		templ, ok := tpl[shell]
		if !ok {
			continue
		}
		sub, err := NewCommand(
			Parent(comp),
			Usage(shell, ""),
			Special("comp:"+shell),
			Exec(func(ctx context.Context) error {
				c, _ := Ctx(ctx)
				_, _ = fmt.Fprintf(c.Stdout, templ, cmd.RootName())
				return nil
			}),
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
	tpls := cmd.Comp
	if tpls == nil {
		tpls = templates
	}
	// load shells from templates
	txt, tpl, err := loadTemplates(tpls)
	if err != nil {
		return err
	}
	// add flags
	for _, shell := range slices.Sorted(maps.Keys(txt)) {
		templ, ok := tpl[shell]
		if !ok {
			continue
		}
		special := "hook:comp:" + shell
		f := func(ctx *Context) error {
			_, _ = fmt.Fprintf(ctx.Stdout, templ, cmd.RootName())
			return ErrExit
		}
		if g := cmd.FlagSpecial(special); g != nil {
			g.Type, g.Def, g.NoArg, g.NoArgDef, g.Hidden, g.Special = HookT, f, true, "", true, special
		} else {
			cmd.Flags = cmd.Flags.Hook(
				fmt.Sprintf(text.CompFlagName, shell),
				fmt.Sprintf(text.CompDesc, shell),
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
	// MinDist is the minimum Levenshtein distance for suggestions.
	MinDist int
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
	sb := new(strings.Builder)
	for _, f := range []func(*strings.Builder){
		help.AddBanner,
		help.AddUsage,
		help.AddAliases,
		help.AddExample,
		help.AddCommands,
		help.AddFlags,
		help.AddFooter,
	} {
		f(sb)
	}
	n, err := w.Write(append([]byte(sb.String()), '\n'))
	return int64(n), err
}

// AddBanner adds the command's banner.
func (help *CommandHelp) AddBanner(sb *strings.Builder) {
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
	addBreak(sb)
	_, _ = sb.WriteString(strings.TrimSpace(banner))
}

// AddUsage adds the command's usage.
func (help *CommandHelp) AddUsage(sb *strings.Builder) {
	if help.NoUsage {
		return
	}
	addBreak(sb)
	_, _ = sb.WriteString(text.Usage)
	_, _ = sb.WriteString(":\n  ")
	_, _ = sb.WriteString(strings.Join(help.Command.Tree(), " "))
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
		_, _ = sb.WriteString(" ")
		_, _ = sb.WriteString(spec)
	}
}

// AddAliases adds the command's aliases.
func (help *CommandHelp) AddAliases(sb *strings.Builder) {
	if help.NoAliases || len(help.Command.Aliases) == 0 {
		return
	}
	addBreak(sb)
	_, _ = sb.WriteString(text.Aliases)
	_, _ = sb.WriteString(":\n  ")
	_, _ = sb.WriteString(strings.Join(prepend(help.Command.Aliases, help.Command.Name), ", "))
}

// AddExample adds the command's example.
func (help *CommandHelp) AddExample(sb *strings.Builder) {
	if help.NoExample || strings.TrimSpace(help.Example) == "" {
		return
	}
	addBreak(sb)
	_, _ = sb.WriteString(text.Examples)
	_, _ = sb.WriteString(":\n")
	_, _ = sb.WriteString(help.Example)
}

// AddCommands adds the command's sub commands.
func (help *CommandHelp) AddCommands(sb *strings.Builder) {
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
	addBreak(sb)
	// write commands
	for j, section := range slices.Sorted(maps.Keys(indexes)) {
		if j != 0 {
			_, _ = sb.WriteString("\n\n")
		}
		_, _ = sb.WriteString(sections[section])
		_, _ = sb.WriteString(":")
		for _, i := range indexes[section] {
			c := commands[i]
			_, _ = sb.WriteString("\n  ")
			_, _ = sb.WriteString(c.Name)
			_, _ = sb.WriteString(strings.Repeat(" ", width-len(c.Name)+2))
			_, _ = sb.WriteString(DefaultWrap(c.Usage, width+4))
		}
	}
}

// AddFlags adds the command's flags.
func (help *CommandHelp) AddFlags(sb *strings.Builder) {
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
	addBreak(sb)
	// write flags
	for j, section := range slices.Sorted(maps.Keys(indexes)) {
		if j != 0 {
			_, _ = sb.WriteString("\n\n")
		}
		_, _ = sb.WriteString(sections[section])
		_, _ = sb.WriteString(":")
		for _, i := range indexes[section] {
			_, _ = sb.WriteString("\n  ")
			g := flags[i]
			switch {
			case hasShort && g.Short == "":
				_, _ = sb.WriteString("    ")
			case hasShort:
				_, _ = sb.WriteString("-")
				_, _ = sb.WriteString(g.Short)
				_, _ = sb.WriteString(", ")
			}
			_, _ = sb.WriteString("--")
			_, _ = sb.WriteString(specs[i])
			_, _ = sb.WriteString(strings.Repeat(" ", width-len(specs[i])+2))
			usage := g.Usage
			if g.Def != nil && g.Type != HookT && !g.NoArg {
				if def, err := help.Context.Expand(g.Def); err == nil {
					usage += " " + fmt.Sprintf(text.FlagDefault, def)
				}
			}
			_, _ = sb.WriteString(DefaultWrap(usage, width+offset))
		}
	}
}

// AddFooter adds the command's footer.
func (help *CommandHelp) AddFooter(sb *strings.Builder) {
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
	addBreak(sb)
	_, _ = sb.WriteString(footer)
}

// addBreak conditionally adds a section break.
func addBreak(sb *strings.Builder) {
	if sb.Len() != 0 {
		_, _ = sb.WriteString("\n\n")
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

// templates are the embedded completion templates.
//
//go:embed bash.txt
//go:embed bash.bash
//go:embed fish.txt
//go:embed fish.fish
//go:embed zsh.txt
//go:embed zsh.zsh
//go:embed powershell.txt
//go:embed powershell.ps1
var templates embed.FS
