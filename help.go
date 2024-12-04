package ox

import (
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"

	"github.com/xo/ox/text"
)

// NewVersionFor adds a `--version` flag to a command, or hooks the command's
// flag with `Special == "hook:version"`.
func NewVersionFor(cmd *Command) error {
	if g := cmd.FlagSpecial("hook:version"); g != nil {
		g.Type, g.Def, g.NoArg, g.NoArgDef = HookT, DefaultVersion, true, ""
	} else {
		var opts []Option
		if cmd.Flag(text.VersionFlagShort, false, true) == nil {
			opts = append(opts, Short(text.VersionFlagShort))
		}
		cmd.Flags = cmd.Flags.Hook(text.VersionFlagName, text.VersionFlagDesc, DefaultVersion, opts...)
	}
	return nil
}

// NewHelpFor adds a `--help` flag to a command, or hooks the command's flag
// with `Special == "hook:help"`.
func NewHelpFor(cmd *Command, opts ...Option) error {
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
	// CommandSections are the command's sub command section names.
	CommandSections []string
	// NoAliases when true will not output the command's aliases.
	NoAliases bool
	// NoFlags when true will not output the command's flags.
	NoFlags bool
	// Sections are the command's flag section names.
	Sections []string
	// NoFooter when true will not output the command's footer.
	NoFooter bool
	// Footer is the command's footer.
	Footer string
}

// NewCommandHelp creates command help.
func NewCommandHelp(cmd *Command, opts ...Option) (*CommandHelp, error) {
	help := &CommandHelp{
		Command: cmd,
	}
	if err := applyOpts(help, opts...); err != nil {
		return nil, err
	}
	return help, nil
}

// SetContext sets the context on the help.
func (help *CommandHelp) SetContext(ctx *Context) {
	help.Context = ctx
}

// WriteTo satisfies the [io.WriterTo] interface.
func (help *CommandHelp) WriteTo(w io.Writer) (int64, error) {
	var n int64
	var err error
	// banner
	if !help.NoBanner {
		banner := help.Banner
		if banner == "" {
			banner = help.Command.Name + " " + help.Command.Usage
		}
		banner = strings.TrimSpace(banner)
		if banner != "" {
			n, err = writeStrings(w, n, err, banner, "\n")
		}
	}
	// usage
	if !help.NoUsage {
		n, err = writeBreak(w, n, err)
		n, err = writeStrings(w, n, err, text.Usage, ":\n  ", strings.Join(help.Command.Tree(), " "))
		if !help.NoSpec {
			spec := help.Spec
			if spec == "" {
				var v []string
				if help.Command.Flags != nil && len(help.Command.Flags.Flags) > 0 {
					v = append(v, text.CommandFlagsSpec)
				}
				if len(help.Command.Commands) != 0 {
					v = append(v, text.CommandSubSpec)
				}
				// TODO: don't add args ...
				spec = strings.Join(append(v, text.CommandArgsSpec), " ")
			}
			if spec != "" {
				n, err = writeStrings(w, n, err, " ", spec)
			}
		}
		n, err = writeStrings(w, n, err, "\n")
	}
	// aliases
	if !help.NoAliases && len(help.Command.Aliases) != 0 {
		n, err = writeBreak(w, n, err)
		n, err = writeStrings(w, n, err, text.Aliases, ":\n  ", strings.Join(prepend(help.Command.Aliases, help.Command.Name), ", "), "\n")
	}
	// sub commands
	if !help.NoCommands && len(help.Command.Commands) != 0 {
		sections := make(map[int]string, len(help.CommandSections))
		for i, section := range help.CommandSections {
			sections[i] = section
		}
		// split between sections
		width, indexes := 0, make(map[int][]int)
		for i, c := range help.Command.Commands {
			indexes[c.Section] = append(indexes[c.Section], i)
			width = max(width, DefaultWidth(c.Name))
			if _, ok := sections[c.Section]; !ok {
				sections[c.Section] = text.Commands
			}
		}
		// write commands
		for _, section := range slices.Sorted(maps.Keys(indexes)) {
			n, err = writeBreak(w, n, err)
			n, err = writeStrings(w, n, err, sections[section], ":")
			for _, i := range indexes[section] {
				c := help.Command.Commands[i]
				n, err = writeStrings(w, n, err, "\n  ", c.Name, strings.Repeat(" ", width-len(c.Name)+2), DefaultWrap(c.Usage, width+4))
			}
			n, err = writeStrings(w, n, err, "\n")
		}
	}
	// flags
	if !help.NoFlags && help.Command.Flags != nil && len(help.Command.Flags.Flags) != 0 {
		sections := make(map[int]string, len(help.Sections))
		for i, section := range help.Sections {
			sections[i] = section
		}
		// split between sections
		width, indexes, specs, hasShort := 0, make(map[int][]int), make([]string, len(help.Command.Flags.Flags)), false
		for i, g := range help.Command.Flags.Flags {
			indexes[g.Section] = append(indexes[g.Section], i)
			specs[i] = spec(g)
			width = max(width, len(specs[i]))
			if _, ok := sections[g.Section]; !ok {
				sections[g.Section] = text.Flags
			}
			hasShort = hasShort || g.Short != ""
		}
		offset := 2
		if hasShort {
			offset = 10
		}
		// write flags
		for _, section := range slices.Sorted(maps.Keys(indexes)) {
			n, err = writeBreak(w, n, err)
			n, err = writeStrings(w, n, err, sections[section], ":")
			for _, i := range indexes[section] {
				s, g := "", help.Command.Flags.Flags[i]
				switch {
				case hasShort && g.Short == "":
					s = "    "
				case hasShort:
					s = "-" + g.Short + ", "
				}
				usage := g.Usage
				if g.Def != nil && g.Type != HookT && !g.NoArg {
					if def, err := help.Context.Expand(g.Def); err == nil {
						usage += " " + fmt.Sprintf(text.FlagDefault, def)
					}
				}
				n, err = writeStrings(w, n, err, "\n  ", s, "--", specs[i], strings.Repeat(" ", width-len(specs[i])+2), DefaultWrap(usage, width+offset))
			}
			n, err = writeBreak(w, n, err)
		}
	}
	// footer
	if !help.NoFooter {
		footer := help.Footer
		if footer == "" && len(help.Command.Commands) != 0 {
			footer = fmt.Sprintf(text.Footer, strings.Join(help.Command.Tree(), " "))
		}
		footer = strings.TrimSpace(footer)
		if footer != "" {
			n, err = writeBreak(w, n, err)
			n, err = writeStrings(w, n, err, footer, "\n")
		}
	}
	return n, err
}

// writeStrings writes the strings to w.
func writeStrings(w io.Writer, n int64, err error, strs ...string) (int64, error) {
	if err != nil {
		return n, err
	}
	var i int
	for _, s := range strs {
		if len(s) == 0 {
			continue
		}
		i, err = w.Write([]byte(s))
		n += int64(i)
		if err != nil {
			break
		}
	}
	return n, err
}

// writeBreak conditionally writes a break when n != 0.
func writeBreak(w io.Writer, n int64, err error) (int64, error) {
	if err != nil {
		return n, err
	}
	var i int
	if n != 0 {
		i, err = w.Write([]byte{'\n'})
	}
	return n + int64(i), err
}

// spec returns the spec text for the flag.
func spec(g *Flag) string {
	switch {
	case g.NoArg || g.Type == HookT:
		return g.Name
	case g.Spec != "":
		return g.Name + text.FlagSpecSpacer + g.Spec
	case g.Type == SliceT:
		return g.Name + text.FlagSpecSpacer + g.Elem.String()
	case g.Type == MapT:
		return g.Name + text.FlagSpecSpacer + g.MapKey.String() + "=" + g.Elem.String()
	}
	return g.Name + text.FlagSpecSpacer + g.Type.String()
}
