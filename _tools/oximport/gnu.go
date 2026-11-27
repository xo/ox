package main

import (
	"strings"

	"github.com/xo/ox"
)

// parseGNU reads the help of a program that uses getopt_long or argp, such as
// du, cp and tar:
//
//	Usage: du [OPTION]... [FILE]...
//	  or:  du [OPTION]... --files0-from=F
//	Summarize device usage of the set of FILEs, recursively for directories.
//
//	Mandatory arguments to long options are mandatory for short options too.
//	  -0, --null
//	         end each output line with NUL, not newline
//	      --apparent-size    print apparent sizes rather than device usage
//
// argp, which tar uses, starts a section with a line such as
// ` Main operation mode:`. The text after the options is the footer.
func parseGNU(cmd *Command, help string) {
	lines := splitLines(help)
	var banner []string
	i := 0
	// the text before the first option
	for ; i < len(lines); i++ {
		line := lines[i]
		if _, ok := parseOptionLine(line, false); ok || isArgpSection(line) {
			break
		}
		switch {
		case strings.HasPrefix(line, "Usage: "):
			cmd.Spec = cutName(strings.TrimPrefix(line, "Usage: "), cmd.Name)
		case strings.HasPrefix(line, "  or:  ") && cmd.Spec != "":
			cmd.Spec += "\n" + line
		case line == "Examples:":
			var example []string
			for i+1 < len(lines) && strings.HasPrefix(lines[i+1], " ") {
				i++
				example = append(example, lines[i])
			}
			cmd.Example = trimLines(strings.Join(example, "\n"))
		default:
			banner = append(banner, line)
		}
	}
	cmd.Banner = trimLines(strings.Join(banner, "\n"))
	// the options
	r := &optionReader{}
	section := -1
	for ; i < len(lines); i++ {
		line := lines[i]
		switch {
		case isArgpSection(line):
			section, r.last = cmd.section(strings.TrimSuffix(strings.TrimSpace(line), ":")), nil
			continue
		case r.add(line, section):
			continue
		case indentOf(line) == 0:
			// the footer
			cmd.Footer = trimLines(strings.Join(lines[i:], "\n"))
			i = len(lines)
		}
	}
	cmd.Flags = r.flags
	for _, g := range cmd.Flags {
		typeGetopt(g)
	}
}

// isArgpSection returns true when line starts a section of argp help, such as
// ` Main operation mode:`.
func isArgpSection(line string) bool {
	return strings.HasPrefix(line, " ") &&
		!strings.HasPrefix(line, "  ") &&
		strings.HasSuffix(line, ":") &&
		!strings.HasPrefix(line, " -")
}

// typeGetopt sets the type of a flag of GNU or psql help from its
// placeholder.
func typeGetopt(g *Flag) {
	switch {
	case g.Name == "help" || g.Name == "version":
		g.Hook, g.Type = g.Name, ox.BoolT
	case g.Optional:
		// --backup[=CONTROL] takes a value or none
		g.Type = ox.StringT
	default:
		g.Type, g.Elem = placeholderType(g.Spec)
	}
}

// cutName removes the name of the program from the start of a usage line.
func cutName(s, name string) string {
	if rest, ok := strings.CutPrefix(s, name); ok && (rest == "" || rest[0] == ' ') {
		return strings.TrimSpace(rest)
	}
	return s
}
