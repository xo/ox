package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/xo/ox"
)

// The directives of cobra that the importer reads. See
// [ox.CompDirective].
const (
	// compFilterFileExt says that the completions are file extensions.
	compFilterFileExt = 8
	// compFilterDirs says that the value is a directory.
	compFilterDirs = 16
)

// completion is the output of `prog __complete <args> <partial>`.
type completion struct {
	// names are the completions.
	names []string
	// usages are the descriptions of the completions.
	usages []string
	// directive is the directive on the last line.
	directive int
}

// parseComplete reads the output of `prog __complete`. Each line is `name` or
// `name<tab>description`, and the last line is `:<directive>`. It returns
// false when the output is not from cobra.
func parseComplete(out []byte) (*completion, bool) {
	lines := splitLines(string(out))
	n := len(lines)
	for n != 0 && strings.TrimSpace(lines[n-1]) == "" {
		n--
	}
	if n == 0 || !strings.HasPrefix(lines[n-1], ":") {
		return nil, false
	}
	directive, err := strconv.Atoi(lines[n-1][1:])
	if err != nil {
		return nil, false
	}
	comp := &completion{directive: directive}
	for _, line := range lines[:n-1] {
		if line == "" || strings.HasPrefix(line, "_activeHelp_") {
			continue
		}
		name, usage, _ := strings.Cut(line, "\t")
		comp.names = append(comp.names, name)
		comp.usages = append(comp.usages, usage)
	}
	return comp, true
}

// joinCommands adds to cmd each sub command that __complete lists and the
// help does not. It skips help, and completion on the root, because ox adds
// its own.
func joinCommands(cmd *Command, comp *completion) {
	for i, name := range comp.names {
		switch {
		case strings.HasPrefix(name, "-"),
			name == "help",
			name == "completion" && cmd.Parent == nil:
			continue
		}
		if c := cmd.Command(name); c != nil {
			if c.Usage == "" {
				c.Usage = comp.usages[i]
			}
			continue
		}
		cmd.Commands = append(cmd.Commands, &Command{
			Name:    name,
			Usage:   comp.usages[i],
			Section: -1,
			Parent:  cmd,
		})
	}
	// ox adds its own help and completion
	var v []*Command
	for _, c := range cmd.Commands {
		if c.Name != "help" && (c.Name != "completion" || cmd.Parent != nil) {
			v = append(v, c)
		}
	}
	cmd.Commands = v
}

// joinFlags joins the flags that `__complete <path> -` lists with the flags of
// the help. cobra lists `--name`, then `-x` when the flag has a short name.
// When a command has required flags that are not set, cobra lists only those,
// and not --help, so the importer marks them as required.
func joinFlags(cmd *Command, comp *completion) {
	required := !slices.Contains(comp.names, "--help") && len(comp.names) != 0 && cmd.Flag("help") != nil
	var last *Flag
	for i, name := range comp.names {
		switch {
		case strings.HasPrefix(name, "--"):
			name = strings.TrimSuffix(strings.TrimPrefix(name, "--"), "=")
			if last = cmd.Flag(name); last == nil {
				last = &Flag{
					Name:        name,
					Type:        ox.StringT,
					Usage:       comp.usages[i],
					Section:     -1,
					typeUnknown: true,
				}
				if name == "help" || name == "version" {
					last.Hook, last.Type = name, ox.BoolT
				}
				cmd.Flags = append(cmd.Flags, last)
			}
			last.Required = last.Required || required
		case strings.HasPrefix(name, "-") && last != nil && len(name) > 1:
			if short := name[1:]; last.Short == "" && cmd.Flag(short) == nil {
				last.Short = short
			}
		}
	}
}

// joinValues sets the values of a flag from `__complete <path> --name ""`.
// When the help does not give the type of the flag, and cobra completes the
// sub commands after it, the flag takes no value, and it is a bool.
func joinValues(cmd *Command, g *Flag, comp *completion) {
	if g.typeUnknown && len(comp.names) != 0 && !slices.ContainsFunc(comp.names, func(name string) bool {
		return cmd.Command(name) == nil && name != "help" && name != "completion"
	}) {
		g.Type, g.typeUnknown = ox.BoolT, false
		return
	}
	switch {
	case comp.directive&compFilterFileExt != 0:
		g.Path, g.Exts = "file", comp.names
		return
	case comp.directive&compFilterDirs != 0:
		g.Path = "dir"
		return
	case len(comp.names) == 0 || len(g.Values) != 0:
		return
	}
	g.Values, g.Exact = comp.names, true
	for _, v := range comp.names {
		if !strings.Contains(g.Usage, v) {
			g.Exact = false
		}
	}
}
