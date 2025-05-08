// Command gen parses and generates xo/ox style run entries for well-known,
// commmon commands `docker`, `doctl`, `gh`, `helm`, `hugo`, `kubectl`,
// `podman`, `psql`, `du`, `cp`, `tar`.
//
// Generates the xo/ox API calls based on the command's `<command> help ...`
// output.
package main

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"github.com/xo/ox"
)

var defaultCommands = []string{
	`docker`,
	`doctl`,
	`gh`,
	`helm`,
	`hugo`,
	`kubectl`,
	`podman`,
	`psql`,
}

const banner = "" +
	"Parses and generates xo/ox style run entries for well-known, commmon commands %s." + "\n\n" +
	"Generates the xo/ox API calls based on the command's `<command> help ...` output."

func main() {
	args := &Args{}
	var commands string
	for i, s := range defaultCommands {
		if i != 0 {
			commands += "`, "
		}
		commands += "`" + s
	}
	commands += "`"
	ox.RunContext(
		context.Background(),
		ox.Defaults(ox.Banner(ox.DefaultWrap(fmt.Sprintf(banner, commands), 80, 0))),
		ox.From(args),
		ox.Exec(run(args)),
	)
}

type Args struct {
	Verbose  bool     `ox:"verbose,short:v"`
	Dump     bool     `ox:"dump,short:d"`
	Out      string   `ox:"out directory,short:o,default:gen"`
	Commands []string `ox:"command(s) to generate,short:c"`
}

func run(args *Args) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// set verbose logger
		if args.Verbose {
			logger = func(s string, v ...any) {
				fmt.Fprintf(os.Stderr, s+"\n", v...)
			}
		}
		apps := args.Commands
		if len(args.Commands) == 0 {
			apps = defaultCommands
		}
		slices.Sort(apps)
		logger("apps: %v", apps)
		for i, app := range apps {
			if i != 0 {
				logger("\n\n\n")
			}
			root, err := args.gen(ctx, app)
			if err != nil {
				return err
			}
			if err := args.write(root); err != nil {
				return err
			}
		}
		return nil
	}
}

func (args *Args) gen(ctx context.Context, app string) (*command, error) {
	exec, err := exec.LookPath(app)
	if err != nil {
		return nil, fmt.Errorf("unable to find command %q: %w", app, err)
	}
	logger("generating %s (%s)", app, exec)
	root := &command{
		app:     app,
		exec:    exec,
		name:    app,
		section: -1,
	}
	if err := root.parse(ctx); err != nil {
		return nil, err
	}
	return root, nil
}

func (args *Args) write(root *command) error {
	out := filepath.Join(args.Out, root.app, "main.go")
	logger("  writing %s -> %s", root.app, out)
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return fmt.Errorf("cannot create dir for %s: %w", root.app, err)
	}
	cmd := new(bytes.Buffer)
	if err := args.writeCommand(cmd, root, 2); err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	var extra string
	if root.app == "psql" {
		extra = initTempl
	}
	_, _ = fmt.Fprintf(buf, templ, root.name, cmd.String(), extra)
	if args.Dump {
		_, _ = os.Stdout.Write(buf.Bytes())
		return nil
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return os.WriteFile(out, src, 0o644)
}

func (args *Args) writeCommand(w io.Writer, cmd *command, indent int) error {
	padding := strings.Repeat("\t", indent)
	if cmd.banner != "" {
		fmt.Fprintf(w, "%sox.Banner(%q),\n", padding, cmd.banner)
	}
	fmt.Fprintf(w, "%sox.Usage(%q, %q),\n", padding, cmd.name, cmd.usage)
	if cmd.spec != "" {
		fmt.Fprintf(w, "%sox.Spec(%q),\n", padding, cmd.spec)
	}
	if len(cmd.aliases) != 0 {
		fmt.Fprintf(w, "%sox.Aliases(%s),\n", padding, qlist(cmd.aliases))
	}
	if cmd.example != "" {
		fmt.Fprintf(w, "%sox.Example(%q),\n", padding, cmd.example)
	}
	if len(cmd.commandSections) != 0 {
		fmt.Fprintf(w, "%sox.Sections(%s),\n", padding, qlist(cmd.commandSections))
	}
	if cmd.section != -1 {
		fmt.Fprintf(w, "%sox.Section(%d),\n", padding, cmd.section)
	}
	if len(cmd.sections) != 0 {
		fmt.Fprintf(w, "%sox.Help(ox.Sections(\n", padding)
		for _, sect := range cmd.sections {
			fmt.Fprintf(w, "%s\t%q,\n", padding, sect)
		}
		fmt.Fprintf(w, "%s)),\n", padding)
	}
	if cmd.footer != "" {
		fmt.Fprintf(w, "%sox.Footer(%q),\n", padding, cmd.footer)
	}
	for _, c := range cmd.commands {
		fmt.Fprintf(w, "%sox.Sub(\n", padding)
		if err := args.writeCommand(w, c, indent+1); err != nil {
			return err
		}
		fmt.Fprintf(w, "%s),\n", padding)
	}
	if len(cmd.flags) != 0 {
		fmt.Fprintf(w, "%sox.Flags()", padding)
		for _, g := range cmd.flags {
			if err := args.writeFlag(w, g, indent+1); err != nil {
				return err
			}
		}
		fmt.Fprintf(w, ",\n")
	}
	return nil
}

func (args *Args) writeFlag(w io.Writer, g flag, indent int) error {
	var v []string
	if g.spec != "" {
		v = append(v, fmt.Sprintf("ox.Spec(%q)", g.spec))
	}
	if g.key != "" {
		v = append(v, fmt.Sprintf("ox.MapKey(ox.%sT)", oxType(g.key)))
	}
	if g.elem != "" {
		v = append(v, fmt.Sprintf("ox.Elem(ox.%sT)", oxType(g.elem)))
	}
	if g.dflt != "" {
		v = append(v, fmt.Sprintf("ox.Default(%q)", g.dflt))
	}
	if g.short != "" {
		v = append(v, fmt.Sprintf("ox.Short(%q)", g.short))
	}
	if g.section != -1 {
		v = append(v, fmt.Sprintf("ox.Section(%d)", g.section))
	}
	var opts string
	if len(v) != 0 {
		opts = ", " + strings.Join(v, ", ")
	}
	fmt.Fprintf(w, ".\n%s%s(%q, %q%s)", strings.Repeat("\t", indent), oxType(g.typ), g.name, g.desc, opts)
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
	commandSections []string
	sections        []string
	commands        []*command
	flags           []flag
}

func (cmd *command) parse(ctx context.Context) error {
	args := append([]string{"help"}, cmd.names()[1:]...)
	if cmd.app == "psql" {
		args = []string{"--help"}
	}
	logger("run: %s %s", cmd.exec, strings.Join(args, " "))
	buf, err := runCommand(ctx, cmd.exec, args...)
	if err != nil {
		return err
	}
	buf = cmd.preprocess(buf)
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
			spec, extra, _ := strings.Cut(strings.TrimSpace(s), "\n")
			spec = strings.TrimSpace(strings.TrimPrefix(spec, cmd.String()))
			logger("    spec: %q", spec)
			cmd.spec = spec
			if extra = strings.TrimSpace(extra); extra != "" {
				switch cmd.app {
				case "docker":
					cmd.banner = extra
				}
			}
		case sectAliases:
			m, aliases := make(map[string]bool), strings.Split(strings.TrimSpace(s), ",")
			for _, alias := range aliases {
				v := strings.Split(strings.TrimSpace(alias), " ")
				i := slices.IndexFunc(v, func(s string) bool {
					s = strings.TrimSpace(s)
					return s != cmd.app && s != cmd.name
				})
				if i != -1 {
					v = v[i:]
				}
				str := strings.TrimSpace(strings.Join(v, " "))
				if len(v) == 0 || m[str] || str == cmd.app || str == cmd.name || str == cmd.app+" "+cmd.name {
					continue
				}
				cmd.aliases, m[str] = append(cmd.aliases, str), true
			}
			logger("    aliases: %q", cmd.aliases)
		case sectExample:
			if strings.HasPrefix(s, "\n\n") {
				s = strings.TrimPrefix(s, "\n")
			}
			cmd.example = trimRight(s)
		case sectFooter:
			cmd.footer += strings.TrimSpace(s)
		case sectFlags:
			cmd.parseFlags(section, s)
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

func (cmd *command) parseFlags(sect, s string) {
	m := flagRE.FindAllStringSubmatch(s, -1)
	if m == nil {
		logger("  section %q has no flags!", sect)
		return
	}
	for _, v := range m {
		short, name, typstr, desc := v[1], v[2], strings.TrimSuffix(v[3], ":"), strings.TrimSpace(v[4])
		shortstr := ""
		if short != "" {
			shortstr = " (-" + short + ")"
		}
		logger("    flag --%s%s (%s): %q", name, shortstr, typstr, desc)
		if name == "help" || name == "version" {
			logger("  skipping flag %q", name)
			continue
		}
		cmd.addFlag(sect, name, short, typstr, desc)
	}
}

func (cmd *command) addFlag(sect, name, short, typstr, desc string) {
	const defaultprefix = "(default "
	typ, key, elem, dflt, spec := cmd.parseFlagType(sect, name, typstr, desc)
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
			"$APPCACHE",
			"$APPCONFIG",
			"$CACHE",
			"$CONFIG",
			"$HOME",
			"$USER",
			"$APPNAME",
		} {
			val, err := c.Expand(s)
			if err != nil {
				panic(err)
			}
			dflt = strings.ReplaceAll(dflt, val, s)
		}
	}
	if sect != "Flags" && !slices.Contains(cmd.sections, sect) {
		cmd.sections = append(cmd.sections, sect)
	}
	desc = strings.TrimSpace(strings.ReplaceAll(desc, "\n", " "))
	cmd.flags = append(cmd.flags, flag{
		name:    name,
		short:   short,
		typ:     typ,
		key:     key,
		elem:    elem,
		desc:    desc,
		dflt:    dflt,
		spec:    spec,
		section: len(cmd.sections) - 1,
	})
}

var flagRE = regexp.MustCompile(`(?im)^\s{1,}(?:-(.),\s+)?--([^\s=]+)[ =]([^\s]+)?\s+(.*)$`)

func (cmd *command) parseFlagType(sect, name, typstr, desc string) (ox.Type, ox.Type, ox.Type, string, string) {
	if typstr == "''" || typstr == `""` {
		return ox.StringT, "", "", "", ""
	}
	u := unquote(typstr)
	switch u {
	case "":
		return ox.BoolT, "", "", "", ""
	case "false", "true":
		return ox.BoolT, "", "", "", u
	case "[]", "list":
		return ox.SliceT, "", ox.StringT, "", ""
	case "stringToString", "map":
		return ox.MapT, ox.StringT, ox.StringT, "", ""
	case "int", "uint", "duration", "string", "float32", "float64", "UUID", "uuid", "URL", "url", "date", "int32", "uint16", "uint32":
		return ox.Type(strings.ToLower(u)), "", "", "", ""
	case "byte", "port":
		return ox.UintT, "", "", "", u
	case "number":
		return ox.IntT, "", "", "", u
	case "float", "decimal":
		return ox.Float32T, "", "", "", u
	case "ip":
		return ox.AddrT, "", "", "", ""
	case "0.0.0.0", "127.0.0.1":
		return ox.AddrT, "", "", u, ""
	case "ipNet":
		return ox.CIDRT, "", "", "", ""
	case "node-addr":
		return ox.AddrPortT, "", "", "", ""
	case "10.116.0.0/20", "10.244.0.0/16", "10.245.0.0/16":
		return ox.CIDRT, "", "", u, ""
	case "Y", "N", "Yes", "No":
		return ox.BoolT, "", "", u, ""
	case "[localhost]":
		return ox.SliceT, "", ox.StringT, "localhost", ""
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
		"doctl",
		"COMMAND",
		"DBNAME",
		"FILENAME",
		"STRING",
		"TEXT",
		"HOSTNAME",
		"PORT",
		"USERNAME":
		return ox.StringT, "", "", "", u
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
		return ox.StringT, "", "", u, ""
	}
	switch {
	case isDuration(u):
		return ox.DurationT, "", "", u, ""
	case isNumber(u):
		return ox.IntT, "", "", u, ""
	case strings.HasSuffix(u, "s"):
		typ, _, _, _, _ := cmd.parseFlagType(sect, name, strings.TrimSuffix(u, "s"), desc)
		return ox.SliceT, "", typ, "", ""
	case strings.HasSuffix(u, "Slice"):
		typ, _, _, _, _ := cmd.parseFlagType(sect, name, strings.TrimSuffix(u, "Slice"), desc)
		return ox.SliceT, "", typ, "", ""
	case strings.HasSuffix(u, "Array"):
		typ, _, _, _, _ := cmd.parseFlagType(sect, name, strings.TrimSuffix(u, "Array"), desc)
		return ox.ArrayT, "", typ, "", ""
	case strings.HasPrefix(u, "--") && strings.HasSuffix(u, "=false"):
		return ox.BoolT, "", "", "", ""
	case strings.HasPrefix(u, "--"):
		return ox.StringT, "", "", "", u
	case strings.Contains(typstr, "="):
		return ox.MapT, ox.StringT, ox.StringT, "", u
	case strings.HasPrefix(u, "kubectl-"):
		return ox.StringT, "", "", u, ""
	case strings.HasPrefix(u, "^"):
		return ox.StringT, "", "", u, "regexp"
	case strings.HasPrefix(u, "["):
		return ox.StringT, "", "", u, "glob"
	case strings.HasPrefix(u, "/"):
		return ox.StringT, "", "", u, "path"
	}
	cmd.logUnknownFlagType(sect, name, typstr, desc)
	return ox.StringT, "", "", "", u
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
		ox.DefaultWrap(desc, ox.DefaultWrapWidth, 2),
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
			logger("  skipping command %q", name)
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
			if !slices.Contains(cmd.commandSections, sect) {
				cmd.commandSections = append(cmd.commandSections, sect)
			}
			c.section = len(cmd.commandSections) - 1
		}
		if err := c.parse(ctx); err != nil {
			return fmt.Errorf("%q: %w", name, err)
		}
		cmd.commands = append(cmd.commands, c)
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
	return sectNone, fmt.Errorf("unknown docker section %q", sect)
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
	return sectNone, fmt.Errorf("unknown gh section %q", sect)
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

func (cmd *command) preprocess(buf []byte) []byte {
	if cmd.app != "psql" {
		return buf
	}
	return bytes.ReplaceAll(buf, []byte("PostgreSQL home page:"), []byte("PostgreSQL home page -"))
}

func (cmd *command) firstSection(buf []byte) (string, int) {
	const expstr = "Experimental: "
	offset := 0
	if bytes.HasPrefix(buf, []byte(expstr)) {
		offset = len(expstr)
	}
	var re *regexp.Regexp
	switch cmd.app {
	case "helm", "doctl", "hugo", "podman", "psql":
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

type flag struct {
	name    string
	short   string
	desc    string
	dflt    string
	spec    string
	section int
	typ     ox.Type
	key     ox.Type
	elem    ox.Type
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

func oxType(typ ox.Type) string {
	switch typ {
	case "uuid":
		return "UUID"
	case "url":
		return "URL"
	case "cidr":
		return "CIDR"
	case "addrport":
		return "AddrPort"
	case "datetime":
		return "DateTime"
	case "bigint":
		return "BigInt"
	case "bigfloat":
		return "BigFloat"
	case "bigrat":
		return "BigRat"
	}
	return strings.ToUpper(string(typ[0])) + string(typ[1:])
}

func qlist(v []string) string {
	var str string
	for i, s := range v {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%q", s)
	}
	return str
}

const templ = `// Command %[1]s is a xo/ox version of ` + "`+%[1]s`" + `. 
//
// Generated from _examples/gen.go.
package main

import (
	"context"

	"github.com/xo/ox"
)%[3]s

func main() {
	ox.RunContext(
		context.Background(),
		ox.Defaults(),
%[2]s
	)
}
`

const initTempl = `
import (
	"github.com/xo/ox/text"
)

func init() {
	text.FlagSpecSpacer = "="
}`
