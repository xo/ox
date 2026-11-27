package ox

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

// docs are the doc comments of struct types, from [RegisterDoc].
var docs = make(map[reflect.Type]typeDoc)

// typeDoc is the doc comment of a struct type and of its fields.
type typeDoc struct {
	doc    string
	fields map[string]string
}

// RegisterDoc registers the doc comment of the struct type T, and the doc
// comment of each of its fields, by the name of the field. [App] and [From]
// read them. The tool `github.com/xo/ox/cmd/oxdoc` writes a call for each
// struct type of a program, from `go generate`, because Go keeps no doc
// comment in a program.
func RegisterDoc[T any](doc string, fields map[string]string) {
	docs[reflect.TypeFor[T]()] = typeDoc{doc: doc, fields: fields}
}

// fieldDoc returns the first paragraph of the doc comment of the field of
// typ, from [RegisterDoc].
func fieldDoc(typ reflect.Type, name string) string {
	s := strings.TrimSpace(docs[typ].fields[name])
	if i := strings.Index(s, "\n\n"); i != -1 {
		s = s[:i]
	}
	return s
}

// docSections are the sections of a doc comment that [App] reads.
var docSections = []string{
	"Description",
	"Usage",
	"Aliases",
	"Section",
	"Banner",
	"Examples",
	"Sections",
	"Command sections",
	"Footer",
}

// parseDoc returns the sections of a doc comment, by name. A section starts
// with a line that holds only its name and a colon, and it ends at the next
// section. One tab is removed from the start of each line.
func parseDoc(doc string) map[string]string {
	m := make(map[string]string)
	var name string
	var lines []string
	add := func() {
		for len(lines) != 0 && strings.TrimSpace(lines[0]) == "" {
			lines = lines[1:]
		}
		for len(lines) != 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
			lines = lines[:len(lines)-1]
		}
		if name != "" {
			m[name] = strings.Join(lines, "\n")
		}
	}
	for line := range strings.SplitSeq(doc, "\n") {
		if s, ok := strings.CutSuffix(line, ":"); ok && slices.Contains(docSections, s) {
			add()
			name, lines = s, nil
			continue
		}
		lines = append(lines, strings.TrimPrefix(line, "\t"))
	}
	add()
	return m
}

// appOption returns the option of [App] for v. prefix is the name of the
// type of the parent command, without the suffix `Command`.
func appOption(v any, prefix string) option {
	var opts []Option
	return option{
		name: "App",
		cmd: func(cmd *Command) error {
			var err error
			if opts, err = appOptions(cmd, v, prefix); err != nil {
				return err
			}
			return Apply(cmd, opts...)
		},
		post: func(cmd *Command) error {
			return ApplyPost(cmd, opts...)
		},
	}
}

// appOptions returns the options that build cmd from v. See [App].
func appOptions(cmd *Command, v any, prefix string) ([]Option, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer || val.IsNil() || val.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("%w: %T is not a *struct", ErrInvalidType, v)
	}
	typ := val.Elem().Type()
	m := parseDoc(docs[typ].doc)
	// the name is the word of the usage line at the depth of the command,
	// and the words after it are the spec
	depth := 0
	for c := cmd.Parent; c != nil; c = c.Parent {
		depth++
	}
	var name, spec string
	usage, _, _ := strings.Cut(m["Usage"], "\n")
	if words := strings.Fields(usage); len(words) > depth {
		name, spec = words[depth], strings.Join(words[depth+1:], " ")
	}
	// a sub command without a usage line takes the name of its type, without
	// the suffix Command and without the name of the type of its parent,
	// such as list for AppListCommand
	typeName := strings.TrimSuffix(typ.Name(), "Command")
	if name == "" && depth != 0 {
		s := typeName
		if t, ok := strings.CutPrefix(s, prefix); ok && t != "" {
			s = t
		}
		name = DefaultFlagNameMapper(s)
	}
	desc := m["Description"]
	opts := []Option{option{
		name: "App",
		cmd: func(cmd *Command) error {
			if name != "" {
				cmd.Name = name
			}
			if desc != "" {
				cmd.Usage = desc
			}
			return nil
		},
	}}
	if spec != "" {
		opts = append(opts, Spec(spec))
	}
	if s := m["Aliases"]; s != "" {
		opts = append(opts, Aliases(strings.Fields(s)...))
	}
	if s := m["Section"]; s != "" {
		section, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("%s: %w: section %q: %w", typ, ErrInvalidValue, s, err)
		}
		opts = append(opts, Section(section))
	}
	if s := m["Banner"]; s != "" {
		opts = append(opts, Banner(s))
	}
	if s := m["Examples"]; s != "" {
		opts = append(opts, Example(indent(s, "  ")))
	}
	if s := m["Footer"]; s != "" {
		opts = append(opts, Footer(s))
	}
	if s := m["Sections"]; s != "" {
		opts = append(opts, flagSections(strings.Split(s, "\n")))
	}
	if s := m["Command sections"]; s != "" {
		// at the level of a command, Sections names the sections of the sub
		// commands
		opts = append(opts, Sections(strings.Split(s, "\n")...))
	}
	// flags
	flags, args, err := appendFlags(nil, nil, val, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: cannot build flags: %w", typ, err)
	}
	if len(flags) != 0 || len(args) != 0 {
		opts = append(opts, option{
			name: "App",
			cmd: func(cmd *Command) error {
				if len(flags) != 0 {
					cmd.Flags = cmd.Flags.With(flags...)
				}
				addTagGroups(cmd, flags)
				return cmd.addArgs(args...)
			},
		})
	}
	// sub commands
	if c, ok := v.(interface{ Commands() []any }); ok {
		for _, sub := range c.Commands() {
			opts = append(opts, Sub(appOption(sub, typeName)))
		}
	}
	// exec, before and after
	for _, o := range []Option{appExec(v), appBefore(v), appAfter(v)} {
		if o != nil {
			opts = append(opts, o)
		}
	}
	return opts, nil
}

// appExec returns the [Exec] option of the Exec method of v, or nil when v has
// no Exec method that [Exec] takes.
func appExec(v any) Option {
	switch x := v.(type) {
	case interface {
		Exec(context.Context, []string) error
	}:
		return Exec(x.Exec)
	case interface {
		Exec(context.Context, []string)
	}:
		return Exec(x.Exec)
	case interface{ Exec(context.Context) error }:
		return Exec(x.Exec)
	case interface{ Exec(context.Context) }:
		return Exec(x.Exec)
	case interface{ Exec([]string) error }:
		return Exec(x.Exec)
	case interface{ Exec([]string) }:
		return Exec(x.Exec)
	case interface{ Exec() error }:
		return Exec(x.Exec)
	case interface{ Exec() }:
		return Exec(x.Exec)
	}
	return nil
}

// appBefore returns the [Before] option of the Before method of v, or nil
// when v has no Before method that [Before] takes.
func appBefore(v any) Option {
	switch x := v.(type) {
	case interface {
		Before(context.Context, []string) error
	}:
		return Before(x.Before)
	case interface {
		Before(context.Context, []string)
	}:
		return Before(x.Before)
	case interface{ Before(context.Context) error }:
		return Before(x.Before)
	case interface{ Before(context.Context) }:
		return Before(x.Before)
	case interface{ Before([]string) error }:
		return Before(x.Before)
	case interface{ Before([]string) }:
		return Before(x.Before)
	case interface{ Before() error }:
		return Before(x.Before)
	case interface{ Before() }:
		return Before(x.Before)
	}
	return nil
}

// appAfter returns the [After] option of the After method of v, or nil when v
// has no After method that [After] takes.
func appAfter(v any) Option {
	switch x := v.(type) {
	case interface {
		After(context.Context, []string) error
	}:
		return After(x.After)
	case interface {
		After(context.Context, []string)
	}:
		return After(x.After)
	case interface{ After(context.Context) error }:
		return After(x.After)
	case interface{ After(context.Context) }:
		return After(x.After)
	case interface{ After([]string) error }:
		return After(x.After)
	case interface{ After([]string) }:
		return After(x.After)
	case interface{ After() error }:
		return After(x.After)
	case interface{ After() }:
		return After(x.After)
	}
	return nil
}

// flagSections returns an option that sets the names of the flag sections in
// the help of a command.
func flagSections(sections []string) Option {
	return option{
		name: "App",
		cmd: func(*Command) error {
			return nil
		},
		post: func(cmd *Command) error {
			if cmd.Help == nil {
				cmd.Help, _ = NewCommandHelp(cmd)
			}
			if help, ok := cmd.Help.(*CommandHelp); ok {
				help.Sections = sections
			}
			return nil
		},
	}
}

// indent removes the indent that each line of s has in common, and adds
// prefix to each line that is not empty. gofmt changes the indent of a code
// block in a doc comment, so the help sets its own.
func indent(s, prefix string) string {
	lines := strings.Split(s, "\n")
	var common string
	first := true
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		ws := line[:len(line)-len(strings.TrimLeft(line, " \t"))]
		switch {
		case first:
			common, first = ws, false
		default:
			for !strings.HasPrefix(ws, common) {
				common = common[:len(common)-1]
			}
		}
	}
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			lines[i] = ""
			continue
		}
		lines[i] = prefix + strings.TrimPrefix(line, common)
	}
	return strings.Join(lines, "\n")
}
