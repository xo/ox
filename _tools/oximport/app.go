package main

import (
	"strconv"
	"strings"
)

// app writes the app style: a struct for each command, whose doc comment
// gives the help of the command, and whose fields are the flags, with their
// descriptions in the doc comment of each field. It follows "Apps from a
// struct" in docs/DESIGN.md, and ox.App reads each section that it writes.
func (w *writer) app() {
	w.header(StyleApp)
	w.printf("import \"github.com/xo/ox\"\n\n")
	w.root.Walk(func(cmd *Command) {
		name := typeName(cmd, "Command")
		w.printf("// %s is %s.\n", name, strings.Join(cmd.Tree(), " "))
		if cmd.Parent != nil {
			w.docSection("Description", cmd.Usage)
		}
		w.docSection("Usage", strings.TrimSpace(strings.Join(cmd.Tree(), " ")+" "+cmd.Spec))
		w.docSection("Aliases", strings.Join(cmd.Aliases, "\n"))
		if cmd.Section >= 0 && cmd.Parent != nil {
			w.docSection("Section", strconv.Itoa(cmd.Section))
		}
		w.docSection("Banner", cmd.Banner)
		w.docSection("Examples", cmd.Example)
		w.docSection("Sections", strings.Join(cmd.Sections, "\n"))
		w.docSection("Command sections", strings.Join(cmd.CommandSections, "\n"))
		if cmd.Footer != defaultFooter(cmd) {
			w.docSection("Footer", cmd.Footer)
		}
		w.printf("type %s struct {\n", name)
		for i, f := range fields(cmd.Flags) {
			if i != 0 {
				w.printf("\n")
			}
			for _, line := range splitLines(f.flag.Usage) {
				w.printf("\t// %s\n", line)
			}
			if c := comment(f.flag); c != "" {
				w.printf("\t//\n\t// oximport: %s\n", c)
			}
			w.printf("\t%s %s %s\n", f.name, f.typ, structTag(append([]string{""}, f.opts...)))
		}
		w.printf("}\n\n")
		if len(cmd.Commands) != 0 {
			w.printf("// Commands returns the sub commands of %s.\n", strings.Join(cmd.Tree(), " "))
			w.printf("func (cmd *%s) Commands() []any {\n\treturn []any{\n", name)
			for _, c := range cmd.Commands {
				w.printf("\t\tnew(%s),\n", typeName(c, "Command"))
			}
			w.printf("\t}\n}\n\n")
		}
	})
	w.printf("func main() {\n\tox.AppFrom(new(App))\n}\n")
}

// docSection writes a section of a doc comment, with each line of s indented
// as a code block.
func (w *writer) docSection(title, s string) {
	if s == "" {
		return
	}
	w.printf("//\n// %s:\n//\n", title)
	for _, line := range splitLines(s) {
		if line == "" {
			w.printf("//\n")
		} else {
			w.printf("//\t%s\n", line)
		}
	}
}
