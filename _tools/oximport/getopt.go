package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// maxOptionIndent is the largest indent of an option line. A line that is
// indented more is a line of a description.
const maxOptionIndent = 8

// optionLine is a line of help that starts with one or more option names,
// such as `-c, --command=COMMAND    run only single command`.
type optionLine struct {
	// longs are the long names, without dashes.
	longs []string
	// shorts are the short names, without the dash.
	shorts []string
	// note is the text after a short name, such as `("one")`.
	note string
	// spec is the placeholder of the value.
	spec string
	// optional is true when the value is optional, as `--name[=ARG]`.
	optional bool
	// desc is the text after the names.
	desc string
	// indent is the number of spaces before the first name.
	indent int
}

// parseOptionLine reads an option line. When spaceArg is true, as in pflag
// help, a single space separates a long name from its placeholder, and two
// spaces separate the names from the description. Otherwise, as in GNU help,
// only `=` starts a placeholder.
func parseOptionLine(line string, spaceArg bool) (*optionLine, bool) {
	s := strings.TrimLeft(line, " ")
	opt := &optionLine{indent: len(line) - len(s)}
	if opt.indent > maxOptionIndent || !strings.HasPrefix(s, "-") {
		return nil, false
	}
	i, ok := 0, false
	for {
		switch {
		case strings.HasPrefix(s[i:], "--"):
			i, ok = opt.scanLong(s, i, spaceArg)
		case strings.HasPrefix(s[i:], "-"):
			i, ok = opt.scanShort(s, i)
		}
		if !ok {
			return nil, false
		}
		if !strings.HasPrefix(s[i:], ", -") {
			break
		}
		i += 2
	}
	// the names end at a space, a tab or the end of the line
	if i < len(s) && s[i] != ' ' && s[i] != '\t' {
		return nil, false
	}
	opt.desc = strings.TrimSpace(s[i:])
	return opt, true
}

// scanLong reads the long name at s[i:], such as `--name`, `--name=ARG` or
// `--name[=ARG]`, and its placeholder. It returns the index after them.
func (opt *optionLine) scanLong(s string, i int, spaceArg bool) (int, bool) {
	j := i + 2
	for j < len(s) && !strings.ContainsRune(" \t=[,", rune(s[j])) {
		j++
	}
	if j == i+2 {
		return 0, false
	}
	opt.longs = append(opt.longs, s[i+2:j])
	i = j
	switch {
	case strings.HasPrefix(s[i:], "[="):
		k := strings.IndexByte(s[i:], ']')
		if k == -1 {
			return 0, false
		}
		opt.spec, opt.optional = s[i+2:i+k], true
		return i + k + 1, true
	case strings.HasPrefix(s[i:], "="):
		k := i + 1
		for k < len(s) && !strings.ContainsRune(" ,\t", rune(s[k])) {
			k++
		}
		if spec := s[i+1 : k]; spec != "" {
			opt.spec = spec
		}
		return k, true
	case spaceArg && strings.HasPrefix(s[i:], " ") && !strings.HasPrefix(s[i:], "  "):
		k := i + 1
		for k < len(s) && s[k] != ' ' {
			k++
		}
		// a placeholder is followed by two spaces or the end of the line
		if k == len(s) || strings.HasPrefix(s[k:], "  ") {
			opt.spec = s[i+1 : k]
			return k, true
		}
	}
	return i, true
}

// scanShort reads the short name at s[i:], such as `-c`, and the note after
// it, such as `("one")`. It returns the index after them.
func (opt *optionLine) scanShort(s string, i int) (int, bool) {
	if i+1 >= len(s) {
		return 0, false
	}
	r, n := utf8.DecodeRuneInString(s[i+1:])
	if unicode.IsSpace(r) || r == ',' || r == '-' {
		return 0, false
	}
	opt.shorts = append(opt.shorts, string(r))
	i += 1 + n
	// the note after a short name, such as -1 ("one")
	if strings.HasPrefix(s[i:], " (") {
		if k := strings.IndexByte(s[i:], ')'); k != -1 {
			opt.note, i = s[i+1:i+k+1], i+k+1
		}
	}
	return i, true
}

// optionReader reads the option lines of a help and the lines of their
// descriptions.
type optionReader struct {
	// spaceArg is passed to parseOptionLine.
	spaceArg bool
	// flags are the flags that the reader read.
	flags []*Flag
	// last is the flag whose description continues on the next line.
	last *Flag
	// indent is the indent of the option line of last.
	indent int
}

// add reads a line, and returns false when the line is not an option line or
// a line of a description. An empty line ends the description.
func (r *optionReader) add(line string, section int) bool {
	if strings.TrimSpace(line) == "" {
		r.last = nil
		return true
	}
	if opt, ok := parseOptionLine(line, r.spaceArg); ok {
		r.last, r.indent = r.addOption(opt, section), opt.indent
		return true
	}
	if r.last == nil || indentOf(line) <= r.indent {
		return false
	}
	r.last.appendUsage(strings.TrimSpace(line))
	return true
}

// addOption adds the flag of an option line, or merges it with the flag of
// the same name, as du does with `--time` and `--time=WORD`, and psql with
// `--help[=options]` and `--help=commands`.
func (r *optionReader) addOption(opt *optionLine, section int) *Flag {
	name := ""
	switch {
	case len(opt.longs) != 0:
		name = opt.longs[0]
	case len(opt.shorts) != 0:
		name = opt.shorts[0]
	}
	for _, g := range r.flags {
		if g.Name != name {
			continue
		}
		switch {
		case isValue(opt.spec):
			// --help=commands
			if g.Optional && isValue(g.Spec) && len(g.Values) == 0 {
				g.Values = append(g.Values, g.Spec)
			}
			g.Values, g.Exact = append(g.Values, opt.spec), true
		case g.Spec == "" && opt.spec != "":
			// --time and --time=WORD
			g.Spec, g.Optional = opt.spec, true
			g.appendUsage(opt.desc)
		default:
			g.appendUsage(opt.desc)
		}
		return g
	}
	g := &Flag{
		Name:     name,
		Spec:     opt.spec,
		Optional: opt.optional,
		Note:     opt.note,
		Usage:    opt.desc,
		Section:  section,
	}
	if len(opt.shorts) != 0 {
		g.Short = opt.shorts[0]
	}
	for _, s := range opt.longs[min(1, len(opt.longs)):] {
		g.addAlias(s)
	}
	for _, s := range opt.shorts[min(1, len(opt.shorts)):] {
		g.addAlias(s)
	}
	r.flags = append(r.flags, g)
	return g
}

// isValue returns true when spec is a value, such as commands in
// `--help=commands`, and not a placeholder, such as FILENAME.
func isValue(spec string) bool {
	if spec == "" {
		return false
	}
	for _, c := range spec {
		if !unicode.IsLower(c) && c != '-' && c != '_' {
			return false
		}
	}
	return true
}

// indentOf returns the number of spaces before the text of line. A tab counts
// as 8 spaces.
func indentOf(line string) int {
	n := 0
	for _, c := range line {
		switch c {
		case ' ':
			n++
		case '\t':
			n += 8
		default:
			return n
		}
	}
	return n
}

// splitLines splits s into lines, without the line breaks at the end.
func splitLines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
}
