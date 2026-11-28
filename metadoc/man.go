package metadoc

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xo/ox"
	"github.com/xo/ox/text"
)

// ManOption is an option of [ManPage].
type ManOption func(*man)

// ManPage returns the man page format. Its sub command of `__docs` is
// `man`. The section is 1, the date comes from [Date], and the source is
// the name and the version of the root command.
func ManPage(opts ...ManOption) Format {
	m := &man{section: 1}
	for _, o := range opts {
		o(m)
	}
	return Format{name: "man", w: m}
}

// Section is a [ManPage] option to set the section of the manual, such as 1
// for a program or 8 for a command of the administrator. The section is the
// extension of each file of [WriteDir], such as `app.8`.
func Section(section int) ManOption {
	return func(m *man) {
		m.section = section
	}
}

// Date is a [ManPage] option to set the date of the page. Without it, the
// date is the time of `$SOURCE_DATE_EPOCH`, in seconds since 1970 in UTC, or
// today in UTC. A value of `$SOURCE_DATE_EPOCH` that is not a number is an
// error. The page writes the date as `2006-01-02`.
func Date(date time.Time) ManOption {
	return func(m *man) {
		m.date = date
	}
}

// Source is a [ManPage] option to set the source of the page. Without it,
// the source is the name of the root command and its [ox.Command.Version],
// such as `app v1.2.0`. The version of the build info is not used, so that
// each build writes the same page.
func Source(source string) ManOption {
	return func(m *man) {
		m.source = source
	}
}

// Manual is a [ManPage] option to set the title of the manual. Without it,
// `man` writes the title of the section, such as "General Commands Manual".
func Manual(manual string) ManOption {
	return func(m *man) {
		m.manual = manual
	}
}

// man writes a man page in the roff format of man(7).
type man struct {
	section int
	date    time.Time
	source  string
	manual  string
}

// ext satisfies the writer interface.
func (m *man) ext() string {
	return "." + strconv.Itoa(m.section)
}

// tree satisfies the writer interface.
func (m *man) tree(w io.Writer, root *page) error {
	var sb strings.Builder
	if err := m.header(&sb, root); err != nil {
		return err
	}
	m.top(&sb, root, false)
	if len(root.subs) != 0 {
		m.heading(&sb, text.DocCommands)
		for p := range root.all() {
			if p != root {
				m.entry(&sb, p)
			}
		}
	}
	m.bottom(&sb, root)
	_, err := io.WriteString(w, sb.String())
	return err
}

// page satisfies the writer interface.
func (m *man) page(w io.Writer, p *page) error {
	var sb strings.Builder
	if err := m.header(&sb, p); err != nil {
		return err
	}
	m.top(&sb, p, true)
	if len(p.subs) != 0 {
		m.heading(&sb, text.DocCommands)
		for _, sub := range p.subs {
			_, _ = sb.WriteString(".TP\n")
			_, _ = sb.WriteString(roffBold(sub.title()))
			_, _ = sb.WriteString("\n")
			roffLines(&sb, sub.usage)
		}
	}
	m.bottom(&sb, p)
	// see also
	var refs []string
	if p.parent != nil {
		refs = append(refs, roffBold(p.parent.name())+"("+strconv.Itoa(m.section)+")")
	}
	for _, sub := range p.subs {
		refs = append(refs, roffBold(sub.name())+"("+strconv.Itoa(m.section)+")")
	}
	if len(refs) != 0 {
		m.heading(&sb, text.DocSeeAlso)
		_, _ = sb.WriteString(strings.Join(refs, ",\n"))
		_, _ = sb.WriteString("\n")
	}
	_, err := io.WriteString(w, sb.String())
	return err
}

// header writes the `.TH` line of the page.
func (m *man) header(sb *strings.Builder, p *page) error {
	date, err := m.when()
	if err != nil {
		return err
	}
	source := m.source
	if source == "" {
		name, version := p.root()
		source = strings.TrimSpace(name + " " + version)
	}
	args := []string{strings.ToUpper(p.name()), strconv.Itoa(m.section), date, source}
	if m.manual != "" {
		args = append(args, m.manual)
	}
	_, _ = sb.WriteString(".TH")
	for _, s := range args {
		_, _ = sb.WriteString(" ")
		_, _ = sb.WriteString(roffArg(s))
	}
	_, _ = sb.WriteString("\n")
	return nil
}

// when returns the date of the page.
func (m *man) when() (string, error) {
	if !m.date.IsZero() {
		return m.date.Format(time.DateOnly), nil
	}
	if s := os.Getenv("SOURCE_DATE_EPOCH"); s != "" {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return "", fmt.Errorf("%w: SOURCE_DATE_EPOCH %q: %w", ox.ErrInvalidValue, s, err)
		}
		return time.Unix(n, 0).UTC().Format(time.DateOnly), nil
	}
	return time.Now().UTC().Format(time.DateOnly), nil
}

// top writes the sections of the page before its sub commands: the name,
// the synopsis, the description, the options and the arguments. When global
// is true, it writes the flags that the command inherits.
func (m *man) top(sb *strings.Builder, p *page, global bool) {
	m.heading(sb, text.DocName)
	_, _ = sb.WriteString(roff(p.name()))
	if p.usage != "" {
		_, _ = sb.WriteString(` \- `)
		_, _ = sb.WriteString(roff(oneLine(p.usage)))
	}
	_, _ = sb.WriteString("\n")
	m.heading(sb, text.DocSynopsis)
	roffSynopsis(sb, p)
	if p.banner != "" {
		m.heading(sb, text.DocDescription)
		roffParagraphs(sb, p.banner)
	}
	if len(p.flags) != 0 {
		m.heading(sb, text.DocOptions)
		roffFlags(sb, p.flags)
	}
	if global && len(p.global) != 0 {
		m.heading(sb, text.DocGlobalOptions)
		roffFlags(sb, p.global)
	}
	if len(p.args) != 0 {
		m.heading(sb, text.Arguments)
		roffArgs(sb, p.args)
	}
}

// bottom writes the sections of the page after its sub commands: the
// examples and the footer.
func (m *man) bottom(sb *strings.Builder, p *page) {
	if p.example != "" {
		m.heading(sb, text.Examples)
		roffExample(sb, p.example)
	}
	if p.footer != "" {
		m.heading(sb, text.DocNotes)
		roffParagraphs(sb, p.footer)
	}
}

// entry writes a sub command in the section of the commands of the whole
// tree, with its own flags, arguments and examples.
func (m *man) entry(sb *strings.Builder, p *page) {
	_, _ = sb.WriteString(".SS ")
	_, _ = sb.WriteString(roffArg(p.title()))
	_, _ = sb.WriteString("\n")
	if p.usage != "" {
		roffParagraphs(sb, p.usage)
	}
	roffSynopsis(sb, p)
	if p.banner != "" && p.banner != p.usage {
		roffParagraphs(sb, p.banner)
	}
	if len(p.flags) != 0 {
		roffLabel(sb, text.DocOptions)
		roffFlags(sb, p.flags)
	}
	if len(p.args) != 0 {
		roffLabel(sb, text.Arguments)
		roffArgs(sb, p.args)
	}
	if p.example != "" {
		roffLabel(sb, text.Examples)
		roffExample(sb, p.example)
	}
	if p.footer != "" {
		roffParagraphs(sb, p.footer)
	}
}

// heading writes the heading of a section, in upper case.
func (m *man) heading(sb *strings.Builder, name string) {
	_, _ = sb.WriteString(".SH ")
	_, _ = sb.WriteString(roffArg(strings.ToUpper(name)))
	_, _ = sb.WriteString("\n")
}

// roffSynopsis writes the synopsis of the page: the names of the command and
// its parents in bold, then the spec.
func roffSynopsis(sb *strings.Builder, p *page) {
	_, _ = sb.WriteString(".PP\n")
	_, _ = sb.WriteString(roffBold(p.title()))
	if p.spec != "" {
		_, _ = sb.WriteString(" ")
		_, _ = sb.WriteString(roff(oneLine(p.spec)))
	}
	_, _ = sb.WriteString("\n")
}

// roffLabel writes a label in bold, such as `Options:`, in a paragraph of
// its own.
func roffLabel(sb *strings.Builder, label string) {
	_, _ = sb.WriteString(".PP\n")
	_, _ = sb.WriteString(roffBold(label + ":"))
	_, _ = sb.WriteString("\n")
}

// roffFlags writes each flag as a tagged paragraph.
func roffFlags(sb *strings.Builder, flags []flagDoc) {
	for _, f := range flags {
		_, _ = sb.WriteString(".TP\n")
		for i, name := range f.names {
			if i != 0 {
				_, _ = sb.WriteString(", ")
			}
			_, _ = sb.WriteString(roffBold(name))
		}
		_, _ = sb.WriteString(roff(f.arg))
		_, _ = sb.WriteString("\n")
		roffLines(sb, f.usage)
		for _, val := range f.values {
			_, _ = sb.WriteString(".TP\n")
			_, _ = sb.WriteString(roffBold(val.Name))
			_, _ = sb.WriteString("\n")
			roffLines(sb, val.Usage)
		}
	}
}

// roffArgs writes each positional argument as a tagged paragraph.
func roffArgs(sb *strings.Builder, args []ox.HelpLine) {
	for _, a := range args {
		_, _ = sb.WriteString(".TP\n")
		_, _ = sb.WriteString(`\fI`)
		_, _ = sb.WriteString(roff(a.Name))
		_, _ = sb.WriteString(`\fR`)
		_, _ = sb.WriteString("\n")
		roffLines(sb, a.Usage)
	}
}

// roffExample writes the example text as it is, without filling its lines.
func roffExample(sb *strings.Builder, example string) {
	_, _ = sb.WriteString(".PP\n.RS 4\n.nf\n")
	_, _ = sb.WriteString(roff(example))
	_, _ = sb.WriteString("\n.fi\n.RE\n")
}

// roffParagraphs writes the text as paragraphs. An empty line starts a new
// paragraph.
func roffParagraphs(sb *strings.Builder, s string) {
	for para := range strings.SplitSeq(s, "\n\n") {
		if para = strings.Trim(para, "\n"); para == "" {
			continue
		}
		_, _ = sb.WriteString(".PP\n")
		_, _ = sb.WriteString(roff(para))
		_, _ = sb.WriteString("\n")
	}
}

// roffLines writes the text as the body of a tagged paragraph. An empty line
// is a space between the lines.
func roffLines(sb *strings.Builder, s string) {
	if s = strings.Trim(s, "\n"); s == "" {
		return
	}
	for line := range strings.SplitSeq(s, "\n") {
		if strings.TrimSpace(line) == "" {
			_, _ = sb.WriteString(".sp\n")
			continue
		}
		_, _ = sb.WriteString(roff(line))
		_, _ = sb.WriteString("\n")
	}
}

// roffBold returns the text in bold.
func roffBold(s string) string {
	return `\fB` + roff(s) + `\fR`
}

// roffArg returns the text as an argument of a request, such as the title
// of a section, in double quotes. A hyphen stays as it is, and a backslash
// becomes `\e`.
func roffArg(s string) string {
	return `"` + strings.NewReplacer(`\`, `\e`, `"`, `\(dq`).Replace(oneLine(s)) + `"`
}

// roff escapes the text for a man page. A backslash becomes `\e`, a hyphen
// becomes `\-`, so that a flag such as `--dir` is copied as it is, and a
// line that starts with a `.` or a `'`, which roff reads as a request, starts
// with `\&`.
func roff(s string) string {
	s = strings.NewReplacer(`\`, `\e`, `-`, `\-`).Replace(s)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, ".") || strings.HasPrefix(line, "'") {
			lines[i] = `\&` + line
		}
	}
	return strings.Join(lines, "\n")
}

// oneLine returns the text with each line break as a space.
func oneLine(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
