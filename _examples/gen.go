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
		for i, app := range apps {
			if i != 0 {
				logger("\n\n\n")
			}
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
	spec            string
	example         string
	aliases         []string
	footer          string
	sections        []string
	commandSections []string
	commands        []*command
	flags           []flag
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
			if footer := cmd.footer; footer != "" {
				if len(footer) > 30 {
					footer = footer[:30] + "..."
				}
				logger("  footer: %q", footer)
			}
		}
		// parse the section
		logger("  parsing %q: %q to %q", cmd.String(), section, next)
		typ, err := cmd.parseSect(section)
		if err != nil {
			return err
		}
		logger("  process %s %q", typ, section)
		switch typ {
		case sectNone:
		case sectUsage:
			spec, _, _ := strings.Cut(strings.TrimSpace(s), "\n")
			spec = strings.TrimSpace(strings.TrimPrefix(spec, cmd.String()))
			logger("    spec: %q", spec)
			cmd.spec = spec
		case sectAliases:
			for _, str := range strings.Split(strings.TrimSpace(s), ",") {
				if str = strings.TrimSpace(strings.TrimPrefix(str, cmd.String())); str != "" {
					cmd.aliases = append(cmd.aliases, str)
				}
			}
			logger("    aliases: %v", cmd.aliases)
		case sectExample:
			cmd.example = trimRight(s)
		case sectFooter:
			cmd.footer += trimRight(s)
		case sectFlags:
			if err := cmd.parseFlags(section, s); err != nil {
				return fmt.Errorf("parsing flags: %w", err)
			}
		case sectCommands:
			if err := cmd.parseCommands(ctx, section, s); err != nil {
				return fmt.Errorf("parsing commands: %w", err)
			}
		}
		if next == "" {
			break
		}
	}
	return nil
}

func (cmd *command) parseFlags(sect, s string) error {
	m := flagRE.FindAllStringSubmatch(s, -1)
	if m == nil {
		logger("  section %q has no flags!", sect)
		return nil
	}
	for _, v := range m {
		short, name, typstr, desc := v[1], v[2], strings.TrimSuffix(v[3], ":"), strings.TrimSpace(v[4])
		shortstr := ""
		if short != "" {
			shortstr = " (-" + short + ")"
		}
		logger("    flag --%s%s (%s): %q", name, shortstr, typstr, desc)
		if name == "help" || name == "version" {
			continue
		}
		if err := cmd.addFlag(sect, name, short, typstr, desc); err != nil {
			return err
		}
	}
	return nil
}

func (cmd *command) addFlag(sect, name, short, typstr, desc string) error {
	const defaultprefix = "(default "
	typ, key, elem, dflt, spec, err := cmd.parseFlagType(sect, name, typstr, desc)
	if err != nil {
		return err
	}
	if i := strings.LastIndex(desc, defaultprefix); i != -1 && strings.HasSuffix(desc, ")") {
		dflt = strings.TrimSpace(unquote(desc[i+len(defaultprefix) : len(desc)-1]))
		desc = desc[:i]
	}
	if dflt != "" {
		c := &ox.Context{
			Root: &ox.Command{
				Name: cmd.app,
			},
		}
		for _, s := range []string{
			"$APPNAME",
			"$HOME",
			"$USER",
			"$CONFIG",
			"$APPCONFIG",
			"$CACHE",
			"$APPCACHE",
			"$NUMCPU",
			"$ARCH",
			"$OS",
		} {
			val, err := c.Expand(s)
			if err != nil {
				panic(err)
			}
			if strings.HasPrefix(dflt, val) {
				dflt = strings.TrimPrefix(dflt, val)
			}
		}
	}
	desc = strings.TrimSpace(strings.ReplaceAll(desc, "\n", " "))
	cmd.flags = append(cmd.flags, flag{
		name:  name,
		short: short,
		typ:   typ,
		key:   key,
		elem:  elem,
		desc:  desc,
		dflt:  dflt,
		spec:  spec,
	})
	return nil
}

var flagRE = regexp.MustCompile(`(?im)^\s{1,}(?:-(.),\s+)?--([^\s=]+)[ =]([^\s]+)?\s+(.*)$`)

func (cmd *command) parseFlagType(sect, name, typstr, desc string) (ox.Type, ox.Type, ox.Type, string, string, error) {
	if typstr == "''" || typstr == `""` {
		return ox.StringT, "", "", "", "", nil
	}
	u := unquote(typstr)
	switch u {
	case "":
		return ox.BoolT, "", "", "", "", nil
	case "false", "true":
		return ox.BoolT, "", "", "", u, nil
	case "[]", "list":
		return ox.SliceT, "", ox.StringT, "", "", nil
	case "stringToString", "map":
		return ox.MapT, ox.StringT, ox.StringT, "", "", nil
	case "int", "uint", "duration", "string", "float32", "float64", "UUID", "uuid", "URL", "url", "date", "int32", "uint16", "uint32":
		return ox.Type(strings.ToLower(u)), "", "", "", "", nil
	case "byte", "port":
		return ox.UintT, "", "", "", u, nil
	case "number":
		return ox.IntT, "", "", "", u, nil
	case "float", "decimal":
		return ox.Float32T, "", "", "", u, nil
	case "ip":
		return ox.AddrT, "", "", "", "", nil
	case "0.0.0.0", "127.0.0.1":
		return ox.AddrT, "", "", u, "", nil
	case "ipNet":
		return ox.CIDRT, "", "", "", "", nil
	case "node-addr":
		return ox.AddrPortT, "", "", "", "", nil
	case "10.116.0.0/20", "10.244.0.0/16", "10.245.0.0/16":
		return ox.CIDRT, "", "", u, "", nil
	case "Y", "N", "Yes", "No":
		return ox.BoolT, "", "", u, "", nil
	case "[localhost]":
		return ox.SliceT, "", ox.StringT, "localhost", "", nil
	case
		"format",
		"postRendererString",
		"postRendererArgsSlice",
		"path/to/file.yaml",
		"path/to/file.json",
		"PATH",
		"argfile.conf",
		"host:ip",
		"[username[:password]]",
		"pathname",
		"path",
		"type",
		"feature",
		"version",
		"OS/ARCH[/VARIANT]",
		"strategy",
		"command",
		"<number><unit>",
		"containerGID:hostGID:length",
		"variant",
		"VARIANT",
		"DEVICE_NAME:WEIGHT",
		"<number>[<unit>]",
		"Format",
		"ARCH",
		"OS",
		"choice",
		"tmpfs",
		"preset",
		"image",
		"containerUID:hostUID:length",
		"Credential",
		"Pathname",
		"annotation",
		"architecture",
		"Variant",
		"subject",
		"ulimit",
		"filter",
		"gpu-request",
		"mount",
		"network",
		"credential-spec",
		"pref",
		"secret",
		"config",
		"pem-file",
		"external-ca",
		"name:test-domain-1",
		"Certificate",
		"topic:permission",
		"index:permission",
		"application",
		"is_enabled:true",
		"target_protocol:http,target_port:80",
		"node-pool",
		"size",
		"tag",
		"enum",
		"option",
		"FINGERPRINT",
		"schedule",
		"386734086,391669331",
		"entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306",
		"ip:1.2.3.4,cidr:1.2.0.0/16",
		"protocol:http,port:80,path:/index.html,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3",
		"12,33",
		"215,378",
		"protocol:http,",
		"IP",
		"SLUG",
		"ID",
		"file",
		"type:value",
		"type:cookies,",
		"https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ",
		"v1/insights/droplet/memory_utilization_percent",
		"name",
		"[HOST/]OWNER/REPO",
		"login",
		"username",
		"expression",
		"field",
		"title",
		"query",
		"branch",
		"handle",
		"SHA",
		"directory",
		"repository",
		"key:value",
		"event",
		"user",
		"owner/number",
		"environment",
		"organization",
		"repositories",
		"doctl":
		return ox.StringT, "", "", "", u, nil
	case
		"URN",
		"DropletID",
		"Index",
		"EXTERNAL",
		"REGIONAL",
		"Development",
		"Digest",
		"Update",
		"text",
		"MonthToDateBalance",
		"Date",
		"ResourceID",
		"Tag",
		"Datetime",
		"Label",
		"A",
		"Day",
		"Domain",
		"Email",
		"Limit",
		"Name",
		"Size",
		"Slug",
		"slug",
		"URI",
		"User",
		"completed",
		"greater_than",
		"DropletCPUUtilizationPercent",
		"create",
		"custom",
		"db-s-1vcpu-1gb",
		"droplet",
		"frontend,backend",
		"issue",
		"key",
		"lb-small",
		"mysql",
		"none",
		"nyc",
		"nyc1",
		"pg",
		"protocol:tcp,ports:22,address:0.0.0.0/0",
		"protocol:tcp,ports:22,address:192.0.2.0/24",
		"s-1vcpu-1gb",
		"s-2vcpu-2gb",
		"session",
		"sfo2",
		"strict",
		"ubuntu-20-04-x64",
		"{{.ID}}",
		"nyc3",
		"caching_sha2_password",
		"k8s",
		"Volumes",
		"any",
		"GreaterThan",
		"production-alerts",
		"latency",
		"us_east",
		"ping",
		"PreemptLowerPriority",
		"https://index.docker.io/v1/",
		"merge",
		"strategic",
		"helm",
		"LoadRestrictionsRootOnly",
		"bridge",
		"background",
		"Always",
		"plaintext",
		"json",
		"yaml",
		"legacy",
		"*":
		return ox.StringT, "", "", u, "", nil
	}
	switch {
	case isDuration(u):
		return ox.DurationT, "", "", u, "", nil
	case isNumber(u):
		return ox.IntT, "", "", u, "", nil
	case strings.HasSuffix(u, "s"):
		if typ, _, _, _, _, err := cmd.parseFlagType(sect, name, strings.TrimSuffix(u, "s"), desc); err == nil {
			return ox.SliceT, "", typ, "", "", nil
		}
	case strings.HasSuffix(u, "Array"):
		if typ, _, _, _, _, err := cmd.parseFlagType(sect, name, strings.TrimSuffix(u, "Array"), desc); err == nil {
			return ox.SliceT, "", typ, "", "", nil
		}
	case strings.HasSuffix(u, "Slice"):
		if typ, _, _, _, _, err := cmd.parseFlagType(sect, name, strings.TrimSuffix(u, "Slice"), desc); err == nil {
			return ox.SliceT, "", typ, "", "", nil
		}
	case strings.HasPrefix(u, "--") && strings.HasSuffix(u, "=false"):
		return ox.BoolT, "", "", "", "", nil
	case strings.HasPrefix(u, "--"):
		return ox.StringT, "", "", "", u, nil
	case strings.Contains(typstr, "="):
		return ox.MapT, ox.StringT, ox.StringT, "", u, nil
	case strings.HasPrefix(u, "kubectl-"):
		return ox.StringT, "", "", u, "", nil
	case strings.HasPrefix(u, "^"):
		return ox.StringT, "", "", u, "regexp", nil
	case strings.HasPrefix(u, "["):
		return ox.StringT, "", "", u, "glob", nil
	case strings.HasPrefix(u, "/"):
		return ox.StringT, "", "", u, "path", nil
	}
	cmd.logUnknownFlagType(sect, name, typstr, desc)
	return ox.StringT, "", "", "", u, nil
}

func (cmd *command) logUnknownFlagType(sect, name, typstr, desc string) {
	f, err := os.OpenFile("unknown-flag-types.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(
		f,
		"----------\ncmnd: %s --help\nsect: %q\nflag: %s\ntype: %q\ndesc:\n\n  %s\n\n",
		cmd.String(),
		sect,
		name,
		unquote(typstr),
		ox.DefaultWrap(desc, 2),
	)
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
		return sectExample, nil
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
	case "docker endpoint config", "experimental":
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
	if banner := cmd.banner; banner != "" {
		if len(banner) > 30 {
			banner = banner[:30] + "..."
		}
		logger("  banner: %q", banner)
	}
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

type flag struct {
	name  string
	short string
	desc  string
	dflt  string
	spec  string
	typ   ox.Type
	key   ox.Type
	elem  ox.Type
}

type sectType int

const (
	sectNone sectType = iota
	sectUsage
	sectAliases
	sectExample
	sectFooter
	sectFlags
	sectCommands
)

func (typ sectType) String() string {
	switch typ {
	case sectNone:
		return "none"
	case sectUsage:
		return "usage"
	case sectAliases:
		return "aliases"
	case sectExample:
		return "example"
	case sectFooter:
		return "footer"
	case sectFlags:
		return "flags"
	case sectCommands:
		return "commands"
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

func isDuration(s string) bool {
	if s != "" {
		return durationRE.MatchString(s)
	}
	return false
}

var durationRE = regexp.MustCompile(`^([0-9]+h)?([0-9]+m)?([0-9]+s)?$`)

func isNumber(s string) bool {
	if s != "" {
		return numRE.MatchString(s)
	}
	return false
}

var numRE = regexp.MustCompile(`^-?[0-9]+$`)

func unquote(s string) string {
	switch {
	case strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'"),
		strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`):
		return s[1 : len(s)-1]
	}
	return s
}
