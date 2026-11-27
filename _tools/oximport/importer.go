package main

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// maxDepth is the most levels of sub commands that the importer reads.
const maxDepth = 12

// Importer reads the command-line interface of a program.
type Importer struct {
	// Runner runs the program.
	Runner Runner
	// Format is the format of the help. When it is empty, the importer finds
	// it.
	Format Format
	// Values is true when the importer asks __complete for the values of each
	// flag.
	Values bool
	// Path is the path of commands to read, such as pr list. The importer
	// reads each command on the path, and each command below its end. It
	// reads only the name and the description of the other commands.
	Path []string
	// Depth is the number of levels of commands below the end of Path that
	// the importer reads. A negative depth reads every level.
	Depth int
	// Logf writes a message.
	Logf func(string, ...any)

	// complete is true when the program answers __complete.
	complete bool
}

// Import reads the program name.
func (imp *Importer) Import(ctx context.Context, name string) (*Command, error) {
	if imp.Logf == nil {
		imp.Logf = func(string, ...any) {}
	}
	root := &Command{Name: name, Section: -1}
	if err := imp.detect(ctx, root); err != nil {
		return nil, err
	}
	if err := imp.read(ctx, root, imp.Path, imp.Depth); err != nil {
		return nil, err
	}
	return root, nil
}

// detect finds the format of the help of the root command. It asks
// `__complete` only when the help does not have a format that cobra does not
// write, because a program that is not built with cobra runs with
// `__complete` as an argument.
func (imp *Importer) detect(ctx context.Context, root *Command) error {
	if imp.Format != "" {
		root.Format = imp.Format
		if imp.Format == FormatCobra {
			out, err := imp.Runner.Run(ctx, "__complete", "")
			if err == nil {
				_, imp.complete = parseComplete(out)
			}
		}
		return nil
	}
	help, err := imp.Runner.Run(ctx, "--help")
	if err != nil {
		return err
	}
	if format := detectFormat(string(help)); format != "" {
		root.Format = format
		return nil
	}
	out, err := imp.Runner.Run(ctx, "__complete", "")
	if err != nil {
		return err
	}
	if _, imp.complete = parseComplete(out); !imp.complete && !isCobra(string(help)) {
		return fmt.Errorf("%w: %s: pass --format", ErrUnknownFormat, root.Name)
	}
	root.Format = FormatCobra
	imp.Logf("%s: format %s, __complete %t", root.Name, root.Format, imp.complete)
	return nil
}

// detectFormat returns the format of help, or an empty string for the help
// of cobra, or for help that it does not know.
func detectFormat(help string) Format {
	lines := splitLines(help)
	switch {
	case isGoFlag(help):
		return FormatGoFlag
	case len(lines) != 0 && strings.HasPrefix(lines[0], "Usage: "):
		return FormatGNU
	}
	for i, line := range lines {
		if psqlSectionRE.MatchString(line) && isPsqlSection(lines, i) {
			return FormatPsql
		}
	}
	return ""
}

// psqlSectionRE matches the title of a section of psql, such as "Input and
// output options:".
var psqlSectionRE = regexp.MustCompile(`^[A-Z][a-z]+( [a-z]+)* options:$`)

// isCobra returns true when help has the sections of cobra.
func isCobra(help string) bool {
	lines := splitLines(help)
	return slices.Contains(lines, "Usage:") || slices.Contains(lines, "USAGE")
}

// read reads a command, and the commands below it.
func (imp *Importer) read(ctx context.Context, cmd *Command, path []string, depth int) error {
	args := append(cmd.Path(), "--help")
	imp.Logf("reading %s", strings.Join(cmd.Tree(), " "))
	help, err := imp.Runner.Run(ctx, args...)
	if err != nil {
		return err
	}
	// a command that passes its arguments to a plugin, as helm does, writes
	// the help of its parent for each sub command, without an end
	if len(cmd.Path()) > maxDepth || cmd.Parent != nil && cmd.Parent.help == string(help) {
		imp.Logf("%s: the help repeats, so the importer stops", strings.Join(cmd.Tree(), " "))
		cmd.Stub = true
		return nil
	}
	cmd.help = string(help)
	cmd.Format = cmd.root().Format
	home := imp.Runner.Home()
	switch cmd.Format {
	case FormatCobra:
		parseCobra(cmd, string(help), home)
	case FormatGNU:
		parseGNU(cmd, string(help))
	case FormatPsql:
		parsePsql(cmd, string(help))
	case FormatGoFlag:
		parseGoFlag(cmd, string(help), home)
	default:
		return fmt.Errorf("%w: %s", ErrUnknownFormat, cmd.Format)
	}
	if imp.complete {
		if err := imp.readComplete(ctx, cmd); err != nil {
			return err
		}
	}
	compactSections(cmd)
	for _, c := range cmd.Commands {
		c.Parent = cmd
		switch {
		case len(path) != 0 && c.Name == path[0]:
			err = imp.read(ctx, c, path[1:], depth)
		case len(path) != 0 || depth == 0:
			c.Stub = true
		default:
			err = imp.read(ctx, c, nil, depth-1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// readComplete joins what `__complete` gives with what the help gives.
func (imp *Importer) readComplete(ctx context.Context, cmd *Command) error {
	path := cmd.Path()
	// cobra lists the sub commands, and calls the completion func of the
	// arguments of the command, which can be slow. Ask only on the root and
	// on a command whose help lists sub commands.
	if cmd.Parent == nil || len(cmd.Commands) != 0 {
		out, err := imp.Runner.Run(ctx, append(append([]string{"__complete"}, path...), "")...)
		if err != nil {
			return err
		}
		if comp, ok := parseComplete(out); ok {
			joinCommands(cmd, comp)
		}
	}
	out, err := imp.Runner.Run(ctx, append(append([]string{"__complete"}, path...), "-")...)
	if err != nil {
		return err
	}
	if comp, ok := parseComplete(out); ok {
		joinFlags(cmd, comp)
	}
	if !imp.Values {
		return nil
	}
	for _, g := range cmd.Flags {
		if g.NoValue() || g.Inherited || cmd.Inherits(g.Name) {
			continue
		}
		out, err := imp.Runner.Run(ctx, append(append([]string{"__complete"}, path...), "--"+g.Name, "")...)
		if err != nil {
			return err
		}
		if comp, ok := parseComplete(out); ok {
			joinValues(cmd, g, comp)
		}
	}
	return nil
}

// root returns the root command.
func (cmd *Command) root() *Command {
	c := cmd
	for c.Parent != nil {
		c = c.Parent
	}
	return c
}

// compactSections removes the names of the sections that have no flag and no
// command, and numbers the sections again.
func compactSections(cmd *Command) {
	cmd.Sections = compact(cmd.Sections, func(i, j int) {
		for _, g := range cmd.Flags {
			if g.Section == i {
				g.Section = j
			}
		}
	}, func(i int) bool {
		return slices.ContainsFunc(cmd.Flags, func(g *Flag) bool { return g.Section == i })
	})
	cmd.CommandSections = compact(cmd.CommandSections, func(i, j int) {
		for _, c := range cmd.Commands {
			if c.Section == i {
				c.Section = j
			}
		}
	}, func(i int) bool {
		return slices.ContainsFunc(cmd.Commands, func(c *Command) bool { return c.Section == i })
	})
}

// compact keeps the names that used returns true for, and calls move for each
// index that changes.
func compact(names []string, move func(int, int), used func(int) bool) []string {
	var v []string
	for i, name := range names {
		if !used(i) {
			continue
		}
		if j := len(v); j != i {
			move(i, j)
		}
		v = append(v, name)
	}
	return v
}
