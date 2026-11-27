package main

import (
	"slices"
	"strings"

	"github.com/xo/ox"
)

// Format is the format of the help that a program writes.
type Format string

// Formats.
const (
	// FormatCobra is the help of cobra, with pflag flags (`--flag type`).
	FormatCobra Format = "cobra"
	// FormatGNU is the help of GNU getopt_long and argp, as du, cp and tar
	// write it.
	FormatGNU Format = "gnu"
	// FormatPsql is the help of psql.
	FormatPsql Format = "psql"
	// FormatGoFlag is the help of the flag package of Go (`-flag value`).
	FormatGoFlag Format = "goflag"
)

// Command is a command that the importer read.
type Command struct {
	// Name is the name of the command.
	Name string
	// Aliases are the other names of the command.
	Aliases []string
	// Usage is the one line description from the list of commands of the
	// parent.
	Usage string
	// Banner is the text above the usage.
	Banner string
	// Spec is the text after the command path on the usage line.
	Spec string
	// Example is the text of the examples section.
	Example string
	// Footer is the text after the last section.
	Footer string
	// Section is the index of the command in CommandSections of the parent,
	// or -1.
	Section int
	// CommandSections are the names of the sections of the sub commands.
	CommandSections []string
	// Sections are the names of the sections of the flags.
	Sections []string
	// Flags are the flags.
	Flags []*Flag
	// Commands are the sub commands.
	Commands []*Command
	// Parent is the parent command.
	Parent *Command
	// Format is the format of the help.
	Format Format
	// Stub is true when the importer did not read the help of the command.
	Stub bool

	// help is the help that the importer read.
	help string
}

// Path returns the names of the command and its parents, without the root.
func (cmd *Command) Path() []string {
	var v []string
	for c := cmd; c != nil && c.Parent != nil; c = c.Parent {
		v = append(v, c.Name)
	}
	slices.Reverse(v)
	return v
}

// Tree returns the names of the command and all its parents.
func (cmd *Command) Tree() []string {
	var v []string
	for c := cmd; c != nil; c = c.Parent {
		v = append(v, c.Name)
	}
	slices.Reverse(v)
	return v
}

// Flag returns the flag with the long or short name.
func (cmd *Command) Flag(name string) *Flag {
	for _, g := range cmd.Flags {
		if g.Name == name || g.Short == name || slices.Contains(g.Aliases, name) {
			return g
		}
	}
	return nil
}

// Command returns the sub command with the name.
func (cmd *Command) Command(name string) *Command {
	for _, c := range cmd.Commands {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// Inherits returns true when a parent of the command has the flag name.
func (cmd *Command) Inherits(name string) bool {
	for c := cmd.Parent; c != nil; c = c.Parent {
		if g := c.Flag(name); g != nil && !g.Inherited {
			return true
		}
	}
	return false
}

// Walk calls f for the command and each sub command, parents first.
func (cmd *Command) Walk(f func(*Command)) {
	f(cmd)
	for _, c := range cmd.Commands {
		c.Walk(f)
	}
}

// section returns the index of the flag section name, and adds it when
// needed.
func (cmd *Command) section(name string) int {
	if i := slices.Index(cmd.Sections, name); i != -1 {
		return i
	}
	cmd.Sections = append(cmd.Sections, name)
	return len(cmd.Sections) - 1
}

// commandSection returns the index of the command section name, and adds it
// when needed.
func (cmd *Command) commandSection(name string) int {
	if i := slices.Index(cmd.CommandSections, name); i != -1 {
		return i
	}
	cmd.CommandSections = append(cmd.CommandSections, name)
	return len(cmd.CommandSections) - 1
}

// Flag is a flag that the importer read.
type Flag struct {
	// Name is the long name, without dashes. A flag that has only a short
	// name has the short name here too.
	Name string
	// Aliases are the other long and short names.
	Aliases []string
	// Short is the short name, without the dash.
	Short string
	// Note is the text after the short name, such as `("one")` in psql.
	Note string
	// Type is the type of the value.
	Type ox.Type
	// Elem is the type of an element of a slice, an array or a map.
	Elem ox.Type
	// Spec is the placeholder of the value, such as FILENAME.
	Spec string
	// Optional is true when the value is optional, as `--name[=ARG]`.
	Optional bool
	// Default is the default value.
	Default string
	// Usage is the description. It keeps the line breaks of the help.
	Usage string
	// Section is the index of the flag in Sections of the command, or -1.
	Section int
	// Values are the values that the help or __complete lists.
	Values []string
	// Exact is true when the description lists each of Values, so that the
	// values are the only values that the program allows.
	Exact bool
	// Path is "file" or "dir" when __complete completes a path.
	Path string
	// Exts are the extensions that __complete gives for a file.
	Exts []string
	// Inherited is true when the help lists the flag as a flag of a parent.
	Inherited bool
	// Required is true when __complete lists the flag as required.
	Required bool
	// Hook is "help" or "version" for a flag that writes the help or the
	// version.
	Hook string

	// typeUnknown is true when only __complete lists the flag, so that the
	// importer does not know its type.
	typeUnknown bool
	// kubectl is true when the flag was read from a line such as
	// `--chunk-size=500:`, where kubectl writes the default in place of the
	// placeholder.
	kubectl bool
}

// NoValue returns true when the flag takes no value.
func (g *Flag) NoValue() bool {
	return g.Type == ox.BoolT || g.Type == ox.CountT || g.Hook != ""
}

// names returns each name of the flag.
func (g *Flag) names() []string {
	v := []string{g.Name}
	if g.Short != "" && g.Short != g.Name {
		v = append(v, g.Short)
	}
	return append(v, g.Aliases...)
}

// addAlias adds a long or short name to the flag.
func (g *Flag) addAlias(name string) {
	if name == "" || slices.Contains(g.names(), name) {
		return
	}
	g.Aliases = append(g.Aliases, name)
}

// appendUsage appends a line of description to the flag.
func (g *Flag) appendUsage(s string) {
	switch {
	case s == "":
	case g.Usage == "":
		g.Usage = s
	default:
		g.Usage += "\n" + s
	}
}

// trimLines trims the space at the end of each line of s, and the empty lines
// at the start and the end of s.
func trimLines(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	for len(lines) != 0 && lines[0] == "" {
		lines = lines[1:]
	}
	for len(lines) != 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return strings.Join(lines, "\n")
}
