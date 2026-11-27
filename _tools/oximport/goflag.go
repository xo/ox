package main

import (
	"regexp"
	"strings"
)

// parseGoFlag reads the help that the flag package of Go writes:
//
//	usage: gofmt [flags] [path ...]
//	  -cpuprofile string
//	    	write cpu profile to this file
//	  -d	display diffs instead of rewriting files
//
// A name of one letter is a short name. Each other name is a long name.
func parseGoFlag(cmd *Command, help, home string) {
	lines := splitLines(help)
	var banner, footer []string
	var last *Flag
	for _, line := range lines {
		switch m := goFlagRE.FindStringSubmatch(line); {
		case m != nil:
			typ, spec := goFlagType(m[2])
			last = &Flag{
				Name:    m[1],
				Type:    typ,
				Spec:    spec,
				Usage:   m[3],
				Section: -1,
			}
			if len(m[1]) == 1 {
				last.Short = m[1]
			}
			cmd.Flags = append(cmd.Flags, last)
		case last != nil && strings.HasPrefix(line, "    \t"):
			last.appendUsage(strings.TrimPrefix(line, "    \t"))
		case len(cmd.Flags) != 0:
			last = nil
			footer = append(footer, line)
		case cmd.Spec == "" && goUsageRE.MatchString(line):
			m := goUsageRE.FindStringSubmatch(line)
			cmd.Spec = cutName(m[1], cmd.Name)
		default:
			banner = append(banner, line)
		}
	}
	for _, g := range cmd.Flags {
		usage, def := cutDefault(g.Usage)
		g.Usage = usage
		if def != "" {
			g.Default = cleanDefault(g, def, home)
		}
		if g.Name == "help" || g.Name == "h" {
			g.Hook = "help"
		}
	}
	cmd.Banner = trimLines(strings.Join(banner, "\n"))
	cmd.Footer = trimLines(strings.Join(footer, "\n"))
}

var (
	// goFlagRE matches the first line of a flag: `  -name type` or
	// `  -d<tab>description`.
	goFlagRE = regexp.MustCompile(`^  -([^\s-][^\s]*)(?: ([^\s]+))?(?:\t(.*))?$`)
	// goUsageRE matches a usage line, such as `usage: gofmt [flags] [path
	// ...]`.
	goUsageRE = regexp.MustCompile(`^[Uu]sage(?: of)?:? (.*?):?$`)
)

// isGoFlag returns true when help looks like the help of the flag package of
// Go.
func isGoFlag(help string) bool {
	for _, line := range splitLines(help) {
		if strings.HasPrefix(line, "    \t") {
			return true
		}
		if m := goFlagRE.FindStringSubmatch(line); m != nil && m[3] != "" {
			return true
		}
	}
	return false
}
