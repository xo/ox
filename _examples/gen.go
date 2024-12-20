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
				"docker",
				"doctl",
				"gh",
				"helm",
				"hugo",
				"kubectl",
				"podman",
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
	logger("run: %s help %s", cmd.exec, strings.Join(cmd.names()[1:], " "))
	buf, err := runCommand(ctx, cmd.exec, append([]string{"help"}, cmd.names()[1:]...)...)
	if err != nil {
		return err
	}
	next, n := cmd.firstSection(buf)
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
		typ, err := cmd.parseSect(section)
		if err != nil {
			return err
		}
		logger("  section %q --> %s", section, typ)
		switch typ {
		case sectNone:
		case sectCommands:
			if err := cmd.parseCommands(ctx, section, s); err != nil {
				return fmt.Errorf("parsing commands: %w", err)
			}
		}
		// no next section, so bail
		if next == "" {
			break
		}
	}
	return nil
}

func (cmd *command) parseCommands(ctx context.Context, sect, s string) error {
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
		switch {
		case name == "help", name == "completion", name == "version", cmd.skipCommand(name):
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
		if strings.ToLower(sect) != "available commands" {
			cmd.commandSections = append(cmd.commandSections, sect)
			cmd.section = len(cmd.commandSections) - 1
		}
		if err := c.parse(ctx); err != nil {
			return fmt.Errorf("%q: %w", name, err)
		}
	}
	return nil
}

func (cmd *command) skipCommand(name string) bool {
	switch cmd.String() + " " + name {
	case "docker context create":
		return true
	}
	return false
}

func (cmd *command) parseSect(section string) (sectType, error) {
	switch sect := strings.ToLower(section); {
	case sect == "usage":
		return sectUsage, nil
	case sect == "examples", sect == "example":
		return sectExamples, nil
	case sect == "aliases":
		return sectAliases, nil
	case strings.Contains(sect, "options"), strings.Contains(sect, "flags"):
		return sectFlags, nil
	case strings.Contains(sect, "commands"):
		return sectCommands, nil
	default:
		switch cmd.app {
		case "docker":
			return cmd.parseSectDocker(sect)
		case "doctl":
			return cmd.parseSectDoctl(sect)
		case "gh":
			return cmd.parseSectGh(sect)
		}
	}
	return sectNone, fmt.Errorf("unknown section %q", section)
}

func (cmd *command) parseSectDocker(sect string) (sectType, error) {
	switch sect {
	case "docker endpoint config":
		return sectNone, nil
	}
	return sectNone, fmt.Errorf("unknown doctl section %q", sect)
}

func (cmd *command) parseSectDoctl(sect string) (sectType, error) {
	switch {
	case strings.Contains(sect, "manage"),
		strings.Contains(sect, "configure"),
		strings.Contains(sect, "view"):
		return sectCommands, nil
	}
	return sectNone, fmt.Errorf("unknown doctl section %q", sect)
}

func (cmd *command) parseSectGh(sect string) (sectType, error) {
	switch sect {
	case "arguments", "environment variables", "json fields", "help topics":
		return sectNone, nil
	case "learn more":
		return sectFooter, nil
	}
	return sectNone, fmt.Errorf("unknown doctl section %q", sect)
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

func (cmd *command) firstSection(buf []byte) (string, int) {
	const expstr = "Experimental: "
	offset := 0
	if bytes.HasPrefix(buf, []byte(expstr)) {
		offset = len(expstr)
	}
	var re *regexp.Regexp
	switch cmd.app {
	case "helm", "doctl", "hugo", "podman":
		re = regexp.MustCompile(`(?m)^(Usage):\n`)
	default:
		re = cmd.sectRE()
	}
	m := re.FindSubmatchIndex(buf[offset:])
	if m == nil {
		return "", -1
	}
	cmd.banner = trimRight(string(buf[offset:m[2]]))
	return string(buf[offset+m[2] : offset+m[3]]), offset + m[1]
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

type sectType int

const (
	sectNone sectType = iota
	sectUsage
	sectAliases
	sectCommands
	sectExamples
	sectFlags
	sectFooter
)

func (typ sectType) String() string {
	switch typ {
	case sectNone:
		return "none"
	case sectUsage:
		return "usage"
	case sectAliases:
		return "aliases"
	case sectCommands:
		return "commands"
	case sectExamples:
		return "examples"
	case sectFlags:
		return "flags"
	case sectFooter:
		return "footer"
	}
	return fmt.Sprintf("sectType(%d)", int(typ))
}

var (
	// sectionRE matches sections in help output
	sectionRE = regexp.MustCompile(`(?m)^([A-Z][a-zA-Z0-9 ()]{3,50}):`)
	// sectionCapsRE matches all caps sections.
	sectionCapsRE = regexp.MustCompile(`(?m)^([A-Z][A-Z0-9 ]{3,50})$`)
)

var logger = func(string, ...any) {
}
