package main

import (
	"regexp"
	"slices"
	"strings"
	"unicode"

	"github.com/xo/ox"
)

// helpSection is a section of help: a title on a line of its own, and the
// indented lines after it.
type helpSection struct {
	// title is the title, without the colon.
	title string
	// lines are the lines of the section, with their indent.
	lines []string
	// afterUsage is true when the section comes after the usage section.
	afterUsage bool
}

// parseCobra reads the help that cobra writes. It reads the standard template:
//
//	List releases.
//
//	Usage:
//	  helm list [flags]
//
//	Aliases:
//	  list, ls
//
//	Flags:
//	  -a, --all            show all releases (default true)
//	  -o, --output format  prints the output in the format (default table)
//
//	Global Flags:
//	      --debug          enable verbose output
//
// It also reads the help of gh, which writes the titles in capitals without a
// colon (USAGE, FLAGS, INHERITED FLAGS), and the help of kubectl, which
// writes each flag with its default, and its description on the next line:
//
//	Options:
//	    -A, --all-namespaces=false:
//		If present, list the requested object(s) across all namespaces.
func parseCobra(cmd *Command, help, home string) {
	pre, sections, post := splitSections(splitLines(help))
	var banner, footer []string
	banner = append(banner, pre...)
	for _, sect := range sections {
		title := strings.ToLower(sect.title)
		switch {
		case title == "usage":
			readUsage(cmd, sect.lines)
		case title == "aliases":
			readAliases(cmd, sect.lines)
		case title == "examples" || title == "example":
			cmd.Example = strings.TrimRight(strings.Join(sect.lines, "\n"), "\n ")
		case strings.Contains(title, "help topics"):
		case strings.Contains(title, "flags") || strings.Contains(title, "options"):
			if readFlagSection(cmd, sect, home) {
				break
			}
			fallthrough
		case isCommandSection(sect):
			readCommands(cmd, sect)
		case sect.afterUsage:
			footer = append(footer, "", sect.title)
			footer = append(footer, sect.lines...)
		default:
			banner = append(banner, "", sect.title)
			banner = append(banner, sect.lines...)
		}
	}
	footer = append(footer, "")
	footer = append(footer, post...)
	cmd.Banner = trimLines(strings.Join(banner, "\n"))
	cmd.Footer = trimLines(strings.Join(footer, "\n"))
}

// splitSections splits the lines of help into the text before the first
// section, the sections, and the text after the last section.
func splitSections(lines []string) ([]string, []*helpSection, []string) {
	var pre, post []string
	var sections []*helpSection
	var sect *helpSection
	afterUsage := false
	for i, line := range lines {
		switch {
		case isSectionTitle(lines, i):
			sect = &helpSection{
				title:      strings.TrimSuffix(line, ":"),
				afterUsage: afterUsage,
			}
			sections = append(sections, sect)
			afterUsage = afterUsage || strings.EqualFold(sect.title, "usage")
			post = nil
		case sect == nil:
			pre = append(pre, line)
		case line == "" || indentOf(line) != 0:
			if post != nil {
				post = append(post, line)
			} else {
				sect.lines = append(sect.lines, line)
			}
		default:
			// text at the left edge after a section
			post = append(post, line)
		}
	}
	// a section ends at its last indented line
	for _, sect := range sections {
		sect.lines = splitLines(trimLines(strings.Join(sect.lines, "\n")))
	}
	return pre, sections, post
}

// isSectionTitle returns true when line i is the title of a section: a short
// line at the left edge that ends with a colon, or that is in capitals as gh
// writes it, and that an indented line follows.
func isSectionTitle(lines []string, i int) bool {
	line := lines[i]
	if line == "" || indentOf(line) != 0 || len(line) > 60 || !unicode.IsUpper(rune(line[0])) {
		return false
	}
	if i+1 >= len(lines) || indentOf(lines[i+1]) == 0 {
		return false
	}
	if title, ok := strings.CutSuffix(line, ":"); ok {
		return !strings.Contains(title, "  ") && !strings.Contains(title, ":")
	}
	return capsRE.MatchString(line)
}

// capsRE matches a title in capitals, such as INHERITED FLAGS.
var capsRE = regexp.MustCompile(`^[A-Z][A-Z0-9 ]+[A-Z0-9]$`)

// readUsage reads the usage lines of a command.
func readUsage(cmd *Command, lines []string) {
	tree := strings.Join(cmd.Tree(), " ")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		spec, ok := strings.CutPrefix(line, tree)
		if !ok {
			// a usage line that names the command another way
			n := len(cmd.Tree())
			if f := strings.Fields(line); len(f) > n {
				spec = " " + strings.Join(f[n:], " ")
			} else {
				spec = ""
			}
		}
		if cmd.Spec == "" {
			cmd.Spec = strings.TrimSpace(spec)
		} else {
			cmd.Spec += "\n  " + line
		}
	}
}

// readAliases reads the aliases of a command, as `list, ls` or as `gh pr ls`.
func readAliases(cmd *Command, lines []string) {
	for _, line := range lines {
		for s := range strings.SplitSeq(line, ",") {
			f := strings.Fields(s)
			if len(f) == 0 {
				continue
			}
			if name := f[len(f)-1]; name != cmd.Name && !slices.Contains(cmd.Aliases, name) {
				cmd.Aliases = append(cmd.Aliases, name)
			}
		}
	}
}

// commandRE matches a line in a list of commands, such as `  ls   show the
// releases` or `  auth:   Authenticate gh`.
var commandRE = regexp.MustCompile(`^ {1,4}([A-Za-z0-9][A-Za-z0-9_.-]*)(?:(?::\s*| {2,})(.*))?$`)

// looseCommandRE matches a line in a section of commands whose name and
// description have one space between them.
var looseCommandRE = regexp.MustCompile(`^ {1,4}([A-Za-z0-9][A-Za-z0-9_.-]*):? +(.*)$`)

// isCommandSection returns true when the title of a section names commands,
// or when each line of the section is a command.
func isCommandSection(sect *helpSection) bool {
	if strings.Contains(strings.ToLower(sect.title), "command") && !strings.Contains(strings.ToLower(sect.title), "help") {
		return true
	}
	n := 0
	for _, line := range sect.lines {
		switch {
		case line == "":
		case commandRE.MatchString(line):
			n++
		default:
			return false
		}
	}
	return n != 0
}

// readCommands reads a section of commands.
func readCommands(cmd *Command, sect *helpSection) {
	section := -1
	switch strings.ToLower(sect.title) {
	case "available commands", "commands":
	default:
		section = cmd.commandSection(sect.title)
	}
	for _, line := range sect.lines {
		m := commandRE.FindStringSubmatch(line)
		if m == nil {
			// a name as long as the column, as `reserved-ip-action Display`
			m = looseCommandRE.FindStringSubmatch(line)
		}
		if m == nil || cmd.Command(m[1]) != nil {
			continue
		}
		cmd.Commands = append(cmd.Commands, &Command{
			Name:    m[1],
			Usage:   strings.TrimSpace(m[2]),
			Section: section,
			Parent:  cmd,
		})
	}
}

// readFlagSection reads a section of flags, and returns false when the
// section holds no flags.
func readFlagSection(cmd *Command, sect *helpSection, home string) bool {
	title := strings.ToLower(sect.title)
	inherited := strings.Contains(title, "global") || strings.Contains(title, "inherited")
	section := -1
	switch title {
	case "flags", "options":
	default:
		section = cmd.section(sect.title)
	}
	r := &optionReader{spaceArg: true}
	start := len(cmd.Flags)
	r.flags = cmd.Flags
	for _, line := range sect.lines {
		// kubectl writes -A, --all-namespaces=false:
		s := strings.TrimSpace(line)
		kubectl := strings.HasPrefix(s, "-") && strings.HasSuffix(s, ":") && strings.Contains(s, "=")
		if kubectl {
			line = strings.TrimSuffix(line, ":")
		}
		r.spaceArg = !kubectl
		if r.add(line, section) && kubectl && r.last != nil {
			r.last.kubectl = true
		}
	}
	cmd.Flags = r.flags
	if len(cmd.Flags) == start {
		return false
	}
	for _, g := range cmd.Flags[start:] {
		g.Inherited = inherited
		typePflag(g, home)
	}
	return true
}

// typePflag sets the type and the default of a flag of cobra help.
func typePflag(g *Flag, home string) {
	if g.Name == "help" || g.Name == "version" {
		g.Hook, g.Type, g.Spec = g.Name, ox.BoolT, ""
		g.Usage, _ = cutDefault(g.Usage)
		return
	}
	if g.kubectl {
		// kubectl writes the default where pflag writes the placeholder
		var def string
		g.Type, g.Elem, def = kubectlType(g.Spec)
		g.Spec, g.Usage = "", strings.TrimSpace(g.Usage)
		g.Default = cleanDefault(g, def, home)
		if g.Type == ox.BoolT && g.Default == "false" {
			g.Default = ""
		}
		return
	}
	word, noArg, optional := strings.Cut(g.Spec, "[=")
	g.Type, g.Elem, g.Spec = pflagType(word)
	if optional {
		g.Optional, g.Default = true, unquote(strings.TrimSuffix(noArg, "]"))
	}
	usage, def := cutDefault(g.Usage)
	g.Usage = usage
	if def != "" && !g.Optional {
		g.Default = cleanDefault(g, def, home)
	}
}
