// Command gen parses and generates xo/ox style run entries for well-known,
// commmon commands `kubectl`, `podman`, `docker`, `helm`, `hugo`, `doctl`,
// `gh` based on their `cmd help ...` output.
package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"github.com/xo/ox"
)

func main() {
	args := &Args{}
	ox.RunContext(
		context.Background(),
		ox.Defaults(
			ox.Banner(`Parses and generates xo/ox style run entires for well-known, common commands.`),
		),
		ox.From(args),
		ox.Exec(run(args)),
	)
}

type Args struct {
	Verbose bool   `ox:"verbose,short:v"`
	Out     string `ox:"out directory,short:o,default:generated"`
	Command string `ox:"command to generate,short:c"`
}

func run(args *Args) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// set verbose logger
		if args.Verbose {
			logger = func(s string, v ...any) {
				fmt.Fprintf(os.Stderr, s+"\n", v...)
			}
		}
		var apps []string
		if args.Command != "" {
			apps = []string{args.Command}
		} else {
			apps = []string{
				//"docker",
				//"doctl",
				"gh",
				//"helm",
				//"hugo",
				//"kubectl",
				//"podman",
			}
		}
		for _, app := range apps {
			if err := args.gen(ctx, app); err != nil {
				return err
			}
		}
		return nil
	}
}

func (args *Args) gen(ctx context.Context, app string) error {
	exec, err := exec.LookPath(app)
	if err != nil {
		return fmt.Errorf("unable to find command %q: %w", app, err)
	}
	logger("generating %s (%s)", app, exec)
	root := &command{
		app:     app,
		exec:    exec,
		name:    app,
		section: -1,
	}
	if err := root.parse(ctx); err != nil {
		return err
	}
	root = root
	return nil
}

type command struct {
	app             string
	exec            string
	parent          *command
	name            string
	usage           string
	section         int
	banner          string
	example         string
	aliases         []string
	footer          string
	sections        []string
	commandSections []string
	flags           [][]string
	commands        []*command
}

func (cmd *command) parse(ctx context.Context) error {
	logger("run: %s %s --help", cmd.exec, strings.Join(cmd.names()[1:], " "))
	buf, err := runCommand(ctx, cmd.exec, append(cmd.names()[1:], "--help")...)
	if err != nil {
		return err
	}
	var next string
	var n int
	buf, next, n = cmd.firstSection(buf)
	if n == -1 {
		return nil
	}
	var s string
	for {
		section := next
		// read up to next section
		s, next, n = cmd.nextSection(buf, n)
		if next == "" {
			// at the end, so this is the footer -- split s by '\n\n'
			s, cmd.footer, _ = strings.Cut(s, "\n\n")
			cmd.footer = trimRight(cmd.footer)
		}
		// parse the section
		logger("  parsing %q: %q to %q", cmd.String(), section, next)
		switch sect := strings.ToLower(section); {
		case sect == "usage":
		case sect == "examples":
		case sect == "aliases":
		case sect == "learn more":
		case sect == "arguments":
		case strings.Contains(sect, "options"), strings.Contains(sect, "flags"):
		case strings.Contains(sect, "commands"):
			if err := cmd.parseCommands(ctx, section, s); err != nil {
				return fmt.Errorf("parsing commands: %w", err)
			}
		default:
			return fmt.Errorf("parsing unknown section %q", section)
		}
		// no next section, so bail
		if next == "" {
			break
		}
	}
	return nil
}

func (cmd *command) parseCommands(ctx context.Context, section, s string) error {
	// parse all lines
	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}
		v := strings.SplitN(line, " ", 2)
		if len(v) != 2 {
			return fmt.Errorf("%q: unable to parse command description: %q", cmd.String(), line)
		}
		name, usage := strings.TrimSuffix(strings.TrimSpace(v[0]), ":"), strings.TrimSpace(v[1])
		switch name {
		case "help", "completion", "version":
			continue
		}
		c := &command{
			app:     cmd.app,
			parent:  cmd,
			exec:    cmd.exec,
			name:    name,
			usage:   usage,
			section: -1,
		}
		logger("  command %q -- %q", c.String(), usage)
		if strings.ToLower(section) != "available commands" {
			cmd.commandSections = append(cmd.commandSections, section)
			cmd.section = len(cmd.commandSections) - 1
		}
		if err := c.parse(ctx); err != nil {
			return fmt.Errorf("%q: %w", name, err)
		}
	}
	return nil
}

func (cmd *command) names() []string {
	var v []string
	for c := cmd; c != nil; c = c.parent {
		v = append(v, c.name)
	}
	slices.Reverse(v)
	return v
}

func (cmd *command) String() string {
	return strings.Join(cmd.names(), " ")
}

func (cmd *command) firstSection(buf []byte) ([]byte, string, string, int) {
	// strip out usage

	return buf, "", "", -1
}

// nextSection gets the next section in buf.
func (cmd *command) nextSection(buf []byte, i int) (string, string, int) {
	m := cmd.sectRE().FindIndex(buf[i:])
	if m == nil {
		return string(buf[i:]), "", len(buf)
	}
	section := trimRight(string(buf[i : i+m[0]]))
	next := strings.TrimSuffix(trimRight(string(buf[i+m[0]:i+m[1]])), ":")
	return section, next, i + m[1]
}

func (cmd *command) capsSections() bool {
	switch cmd.app {
	case "gh":
		return true
	}
	return false
}

func (cmd *command) sectRE() *regexp.Regexp {
	if cmd.capsSections() {
		return sectionCapsRE
	}
	return sectionRE
}

func (cmd *command) usageRE() *regexp.Regexp {
	if cmd.capsSections() {
		return usageCapsRE
	}
	return usageRE
}

// parseType parses a type.
func (cmd *command) parseType(s string) (string, error) {
	return "", nil
}

// globalFlags is a map of global flags commands.
var globalFlags = map[string][]string{
	"kubectl": {"options"},
}

func runCommand(ctx context.Context, pathstr string, args ...string) ([]byte, error) {
	buf := new(bytes.Buffer)
	cmd := exec.CommandContext(ctx, pathstr, args...)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func trimRight(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func allNames(app string, names []string) string {
	return strings.Join(append([]string{app}, names...), " ")
}

var (
	// sectionRE matches sections in help output
	sectionRE = regexp.MustCompile(`(?m)^([A-Z][a-zA-Z0-9 ()]{3,50}):$`)
	// sectionCapsRE matches all caps sections.
	sectionCapsRE = regexp.MustCompile(`(?m)^([A-Z][A-Z0-9 ]{3,50})$`)
	// usageRE
	usageRE = regexp.MustCompile(`\nUsage:\s{1,10}([^\n]+)\n\n`)
	// usageCapsRE
	usageCapsRE = regexp.MustCompile(`\nUSAGE\s{1,10}([^\n]+)\n\n`)
)

var logger = func(string, ...any) {
}
