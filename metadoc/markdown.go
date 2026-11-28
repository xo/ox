package metadoc

import (
	"io"
	"strings"
	"unicode"

	"github.com/xo/ox"
	"github.com/xo/ox/text"
)

// MarkdownOption is an option of [Markdown].
type MarkdownOption func(*markdown)

// Markdown returns the Markdown format. Its sub command of `__docs` is
// `markdown`.
func Markdown(opts ...MarkdownOption) Format {
	m := new(markdown)
	for _, o := range opts {
		o(m)
	}
	return Format{name: "markdown", w: m}
}

// FrontMatter is a [Markdown] option to set a func that returns the text
// before the docs of a command, such as the front matter of Hugo. The func
// gets the command of each file of [WriteDir], the command of
// [WriteCommand], and the root of [Write]. The text is written as it is.
func FrontMatter(f func(cmd *ox.Command) string) MarkdownOption {
	return func(m *markdown) {
		m.front = f
	}
}

// Link is a [Markdown] option to set a func that returns the target of a
// link to the page of a command, in the files of [WriteDir] and the page of
// [WriteCommand]. Without it, the target is the name of the file, such as
// `app-sub.md`. [Write] links to the heading of each command, and does not
// use it.
func Link(f func(cmd *ox.Command) string) MarkdownOption {
	return func(m *markdown) {
		m.link = f
	}
}

// markdown writes Markdown, as CommonMark reads it.
type markdown struct {
	front func(*ox.Command) string
	link  func(*ox.Command) string
}

// ext satisfies the writer interface.
func (m *markdown) ext() string {
	return ".md"
}

// tree satisfies the writer interface.
func (m *markdown) tree(w io.Writer, root *page) error {
	var sb strings.Builder
	m.frontMatter(&sb, root)
	m.top(&sb, root, 1, false)
	if len(root.subs) != 0 {
		mdHeading(&sb, 2, text.DocCommands)
		var links []mdLink
		for p := range root.all() {
			if p != root {
				links = append(links, mdLink{p.title(), "#" + anchor(p.title()), p.usage})
			}
		}
		mdLinks(&sb, links)
		for p := range root.all() {
			if p != root {
				m.top(&sb, p, 3, false)
				m.bottom(&sb, p, 4)
			}
		}
	}
	m.bottom(&sb, root, 2)
	// end with one line break
	_, err := io.WriteString(w, strings.TrimRight(sb.String(), "\n")+"\n")
	return err
}

// page satisfies the writer interface.
func (m *markdown) page(w io.Writer, p *page) error {
	var sb strings.Builder
	m.frontMatter(&sb, p)
	m.top(&sb, p, 1, true)
	if len(p.subs) != 0 {
		mdHeading(&sb, 2, text.DocCommands)
		links := make([]mdLink, 0, len(p.subs))
		for _, sub := range p.subs {
			links = append(links, mdLink{sub.title(), m.target(sub), sub.usage})
		}
		mdLinks(&sb, links)
	}
	m.bottom(&sb, p, 2)
	var links []mdLink
	if p.parent != nil {
		links = append(links, mdLink{p.parent.title(), m.target(p.parent), p.parent.usage})
	}
	for _, sub := range p.subs {
		links = append(links, mdLink{sub.title(), m.target(sub), sub.usage})
	}
	if len(links) != 0 {
		mdHeading(&sb, 2, text.DocSeeAlso)
		mdLinks(&sb, links)
	}
	// end with one line break
	_, err := io.WriteString(w, strings.TrimRight(sb.String(), "\n")+"\n")
	return err
}

// frontMatter writes the front matter of the page.
func (m *markdown) frontMatter(sb *strings.Builder, p *page) {
	if m.front == nil {
		return
	}
	if s := m.front(p.cmd); s != "" {
		_, _ = sb.WriteString(s)
		if !strings.HasSuffix(s, "\n") {
			_, _ = sb.WriteString("\n")
		}
		_, _ = sb.WriteString("\n")
	}
}

// target returns the target of a link to the page.
func (m *markdown) target(p *page) string {
	if m.link != nil {
		return m.link(p.cmd)
	}
	return p.file(m.ext())
}

// top writes the heading of the page at the level, then the usage, the
// synopsis, the description, the options and the arguments, one level
// below. When global is true, it writes the flags that the command
// inherits.
func (m *markdown) top(sb *strings.Builder, p *page, level int, global bool) {
	mdHeading(sb, level, p.title())
	if p.usage != "" {
		mdParagraph(sb, p.usage)
	}
	mdHeading(sb, level+1, text.DocSynopsis)
	synopsis := p.title()
	if p.spec != "" {
		synopsis += " " + oneLine(p.spec)
	}
	mdCode(sb, synopsis)
	if p.banner != "" && p.banner != p.usage {
		mdHeading(sb, level+1, text.DocDescription)
		mdParagraph(sb, p.banner)
	}
	if len(p.flags) != 0 {
		mdHeading(sb, level+1, text.DocOptions)
		mdFlags(sb, p.flags)
	}
	if global && len(p.global) != 0 {
		mdHeading(sb, level+1, text.DocGlobalOptions)
		mdFlags(sb, p.global)
	}
	if len(p.args) != 0 {
		mdHeading(sb, level+1, text.Arguments)
		for _, a := range p.args {
			mdItem(sb, codeSpan(a.Name), a.Usage, "")
		}
		_, _ = sb.WriteString("\n")
	}
}

// bottom writes the examples and the footer of the page at the level.
func (m *markdown) bottom(sb *strings.Builder, p *page, level int) {
	if p.example != "" {
		mdHeading(sb, level, text.Examples)
		mdCode(sb, p.example)
	}
	if p.footer != "" {
		mdHeading(sb, level, text.DocNotes)
		mdParagraph(sb, p.footer)
	}
}

// mdLink is a link in a list.
type mdLink struct {
	name, target, usage string
}

// mdLinks writes a list of links, each with its usage.
func mdLinks(sb *strings.Builder, links []mdLink) {
	for _, l := range links {
		mdItem(sb, "["+mdEscape(l.name)+"]("+mdTarget(l.target)+")", l.usage, "")
	}
	_, _ = sb.WriteString("\n")
}

// mdFlags writes a list of flags. The values of a flag are a list in its
// item.
func mdFlags(sb *strings.Builder, flags []flagDoc) {
	for _, f := range flags {
		mdItem(sb, codeSpan(f.spec()), f.usage, "")
		for _, val := range f.values {
			mdItem(sb, codeSpan(val.Name), val.Usage, "  ")
		}
	}
	_, _ = sb.WriteString("\n")
}

// mdItem writes an item of a list: the name, then a colon and the usage.
// Each line of the usage after the first is indented, so that it stays in
// the item.
func mdItem(sb *strings.Builder, name, usage, indent string) {
	_, _ = sb.WriteString(indent)
	_, _ = sb.WriteString("- ")
	_, _ = sb.WriteString(name)
	if usage = strings.Trim(usage, "\n"); usage != "" {
		_, _ = sb.WriteString(": ")
		for i, line := range strings.Split(mdEscape(usage), "\n") {
			if i != 0 {
				_, _ = sb.WriteString("\n")
				if line != "" {
					_, _ = sb.WriteString(indent + "  ")
				}
			}
			_, _ = sb.WriteString(line)
		}
	}
	_, _ = sb.WriteString("\n")
}

// mdHeading writes a heading at the level.
func mdHeading(sb *strings.Builder, level int, s string) {
	_, _ = sb.WriteString(strings.Repeat("#", level))
	_, _ = sb.WriteString(" ")
	_, _ = sb.WriteString(mdEscape(oneLine(s)))
	_, _ = sb.WriteString("\n\n")
}

// mdParagraph writes the text as paragraphs.
func mdParagraph(sb *strings.Builder, s string) {
	_, _ = sb.WriteString(mdEscape(strings.Trim(s, "\n")))
	_, _ = sb.WriteString("\n\n")
}

// mdCode writes the text as a block of code, with a fence that is longer
// than each run of backticks in the text.
func mdCode(sb *strings.Builder, s string) {
	fence := strings.Repeat("`", max(3, longestRun(s, '`')+1))
	_, _ = sb.WriteString(fence)
	_, _ = sb.WriteString("\n")
	_, _ = sb.WriteString(s)
	_, _ = sb.WriteString("\n")
	_, _ = sb.WriteString(fence)
	_, _ = sb.WriteString("\n\n")
}

// codeSpan returns the text as a code span, with a run of backticks that is
// longer than each run of backticks in the text.
func codeSpan(s string) string {
	s = oneLine(s)
	n := longestRun(s, '`') + 1
	fence := strings.Repeat("`", n)
	if strings.HasPrefix(s, "`") || strings.HasSuffix(s, "`") {
		s = " " + s + " "
	}
	return fence + s + fence
}

// longestRun returns the length of the longest run of c in s.
func longestRun(s string, c rune) int {
	var n, longest int
	for _, r := range s {
		if r != c {
			n = 0
			continue
		}
		n++
		longest = max(longest, n)
	}
	return longest
}

// mdEscape escapes the text for Markdown. Each character that marks text,
// such as `*`, `_`, a backtick, `<`, `>`, `[`, `]` and `|`, gets a
// backslash. So does a `#`, a `-`, a `+` or a `=` at the start of a line,
// and the `.` or the `)` of a number at the start of a line, so that the
// line does not start a heading or a list.
func mdEscape(s string) string {
	var sb strings.Builder
	for i, line := range strings.Split(s, "\n") {
		if i != 0 {
			_, _ = sb.WriteString("\n")
		}
		start := len(line) - len(strings.TrimLeft(line, " \t"))
		_, _ = sb.WriteString(line[:start])
		line = line[start:]
		switch {
		case line == "":
		case strings.ContainsRune("#-+=", rune(line[0])):
			_, _ = sb.WriteString(`\`)
		default:
			// a number and a . or a ) start an ordered list
			if j := strings.IndexFunc(line, func(r rune) bool { return !unicode.IsDigit(r) }); j > 0 && (line[j] == '.' || line[j] == ')') {
				_, _ = sb.WriteString(line[:j])
				_, _ = sb.WriteString(`\`)
				line = line[j:]
			}
		}
		for _, r := range line {
			if strings.ContainsRune("\\`*_<>[]|&", r) {
				_, _ = sb.WriteString(`\`)
			}
			_, _ = sb.WriteRune(r)
		}
	}
	return sb.String()
}

// mdTarget returns the target of a link, in angle brackets when it has a
// space or a parenthesis.
func mdTarget(s string) string {
	if strings.ContainsAny(s, " ()<>") {
		return "<" + strings.NewReplacer("<", `\<`, ">", `\>`).Replace(s) + ">"
	}
	return s
}

// anchor returns the anchor of a heading, as GitHub makes it: the text in
// lower case, without each character that is not a letter, a digit, a
// space, a `-` or a `_`, and with a `-` for each space.
func anchor(s string) string {
	var sb strings.Builder
	for _, r := range strings.ToLower(s) {
		switch {
		case r == ' ':
			_, _ = sb.WriteRune('-')
		case r == '-', r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			_, _ = sb.WriteRune(r)
		}
	}
	return sb.String()
}
