package main

import (
	"strings"
)

// parsePsql reads the help of psql:
//
//	psql is the PostgreSQL interactive terminal.
//
//	Usage:
//	  psql [OPTION]... [DBNAME [USERNAME]]
//
//	General options:
//	  -c, --command=COMMAND    run only single command (SQL or internal) and exit
//	  -v, --set=, --variable=NAME=VALUE
//	                           set psql variable NAME to VALUE
//	  -1 ("one"), --single-transaction
//	                           execute as a single transaction (if non-interactive)
//	  -?, --help[=options]     show this help, then exit
//	      --help=commands      list backslash commands, then exit
//
// The text before "Usage:" is the banner, and the text after the last option
// is the footer. A description that goes on the next line belongs to the
// option above it.
func parsePsql(cmd *Command, help string) {
	lines := splitLines(help)
	var banner []string
	i := 0
	for ; i < len(lines) && lines[i] != "Usage:"; i++ {
		banner = append(banner, lines[i])
	}
	cmd.Banner = trimLines(strings.Join(banner, "\n"))
	// the usage lines
	for i++; i < len(lines) && strings.HasPrefix(lines[i], "  "); i++ {
		spec := cutName(strings.TrimSpace(lines[i]), cmd.Name)
		if cmd.Spec == "" {
			cmd.Spec = spec
		} else {
			cmd.Spec += "\n  " + strings.TrimSpace(lines[i])
		}
	}
	// the sections of options
	r := &optionReader{}
	section := -1
	for ; i < len(lines); i++ {
		line := lines[i]
		switch {
		case isPsqlSection(lines, i):
			section, r.last = cmd.section(strings.TrimSuffix(line, ":")), nil
		case r.add(line, section):
		case indentOf(line) == 0:
			cmd.Footer = trimLines(strings.Join(lines[i:], "\n"))
			i = len(lines)
		}
	}
	cmd.Flags = r.flags
	for _, g := range cmd.Flags {
		typeGetopt(g)
	}
}

// isPsqlSection returns true when line i starts a section of options, such as
// "General options:". The line after it is an option.
func isPsqlSection(lines []string, i int) bool {
	line := lines[i]
	if line == "" || indentOf(line) != 0 || !strings.HasSuffix(line, ":") || i+1 >= len(lines) {
		return false
	}
	_, ok := parseOptionLine(lines[i+1], false)
	return ok
}
