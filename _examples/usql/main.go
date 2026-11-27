// Command usql is a sample of the command line of usql, built with ox as
// usql builds it. It has every option of usql, and it does not connect to a
// database. It prints the options that it parsed, in the order of the command
// line.
//
// TestUsqlSample in the root of the repository builds it, and compares its
// help with testdata/usql/help.txt. See B42 and "A sample usql app" in
// docs/DESIGN.md.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xo/ox"
	"github.com/xo/ox/text"
)

// The sections of the help, in the order of psql's help.
const (
	sectionGeneral = iota
	sectionInputOutput
	sectionFormat
	sectionConnection
)

// errWrongNumberOfArguments is the error of more than one DSN.
var errWrongNumberOfArguments = errors.New("wrong number of arguments")

func main() {
	// psql writes an option and its value as --command=COMMAND
	text.FlagSpecSpacer = "="
	args := new(Args)
	err := ox.RunError(context.Background(), append(
		options(args),
		ox.Exec(args.run),
	)...)
	if err != nil {
		// usql writes its own message and its own hint
		fmt.Fprintf(os.Stderr, "usql: %v\n", err)
		if isUsageError(err) {
			fmt.Fprintln(os.Stderr, `Try "usql --help" for more information.`)
		}
		os.Exit(1)
	}
}

// Args are the options of the command line.
type Args struct {
	DSN               string
	CommandOrFiles    []string
	Out               string
	ForcePassword     bool
	NoPassword        bool
	NoInit            bool
	SingleTransaction bool
	// Options are the settings of the command line, as SECTION.NAME=VALUE,
	// in the order of the command line.
	Options    []string
	ConfigFile string
	Dump       bool
	// DumpOptions are the options of a dump, as NAME=VALUE.
	DumpOptions []string
}

// run prints the options.
func (args *Args) run(_ context.Context, v []string) error {
	switch {
	case len(v) != 0 && args.DSN != "":
		return errWrongNumberOfArguments
	case len(v) != 0:
		args.DSN = v[0]
	}
	fmt.Println("dsn:", args.DSN)
	fmt.Println("commands and files:", args.CommandOrFiles)
	fmt.Println("options:", args.Options)
	fmt.Println("out:", args.Out)
	fmt.Println("config:", args.ConfigFile)
	fmt.Println("no init:", args.NoInit, "single transaction:", args.SingleTransaction)
	fmt.Println("no password:", args.NoPassword, "force password:", args.ForcePassword)
	fmt.Println("dump:", args.Dump, args.DumpOptions)
	return nil
}

// options returns the options of the command line, which fill args.
func options(args *Args) []ox.Option {
	flags := ox.Flags().
		// general options
		Var("command", "run only single command (SQL or internal) and exit",
			ox.HookT, ox.Default(args.addCommandOrFile("command")), ox.Short("c"), ox.Spec("COMMAND"), ox.Section(sectionGeneral)).
		String("dbname", "database url or connection name to connect to",
			ox.Short("d"), ox.Spec("DSN"), ox.Bind(&args.DSN), ox.Section(sectionGeneral)).
		Var("file", "execute commands from file, then exit",
			ox.HookT, ox.Default(args.addCommandOrFile("file")), ox.Short("f"), ox.Spec("FILE"), ox.Section(sectionGeneral)).
		// psql's -l runs \l, then exits
		Var("list", "list available databases, then exit",
			ox.HookT, ox.NoArg(true, ""), ox.Default(func(string) error { return args.addCommandOrFile("command")(`\l`) }),
			ox.Short("l"), ox.Section(sectionGeneral)).
		Var("set", `set script variable NAME to VALUE (see \set command)`,
			ox.HookT, ox.Default(args.addOption("vars")), ox.Short("v"), ox.Spec("NAME=VALUE"),
			ox.Aliases("variable"), ox.HiddenAliases("var"), ox.Section(sectionGeneral)).
		Var("option", `set configuration option NAME (SECTION.NAME) to VALUE (see \config command)`,
			ox.HookT, ox.Default(args.addOption("")), ox.Short("O"), ox.Spec("NAME=VALUE"), ox.Section(sectionGeneral)).
		Var("uset", `set setting NAME of usql to VALUE (see \uset command)`,
			ox.HookT, ox.Default(args.addOption("usql")), ox.Short("U"), ox.Spec("NAME=VALUE"), ox.Section(sectionGeneral)).
		Var("cset", `set named connection NAME to DSN (see \cset command)`,
			ox.HookT, ox.Default(args.addOption("connections")), ox.Short("N"), ox.Spec("NAME=DSN"), ox.Section(sectionGeneral)).
		Bool("dump", "write the database as a usql script to standard output or to the file of -o, then exit",
			ox.Short("D"), ox.Bind(&args.Dump), ox.Section(sectionGeneral)).
		String("config", "config file",
			ox.Spec("FILE"), ox.Bind(&args.ConfigFile), ox.Section(sectionGeneral)).
		Var("version", "output version information, then exit",
			ox.Short("V"), ox.Special("hook:version"), ox.Section(sectionGeneral)).
		Bool("no-init", "do not execute initialization scripts",
			ox.Short("X"), ox.Bind(&args.NoInit), ox.HiddenAliases("no-psqlrc", "no-usqlrc"), ox.DeprecatedAliases("no-rc"),
			ox.Section(sectionGeneral)).
		Bool("single-transaction", "execute as a single transaction (if non-interactive)",
			ox.Short("1"), ox.Bind(&args.SingleTransaction), ox.Section(sectionGeneral)).
		Var("help", "show this help, then exit",
			ox.Short("?"), ox.Special("hook:help"), ox.Section(sectionGeneral)).
		// input and output options
		Var("echo-hidden", "display queries that internal commands generate",
			ox.HookT, ox.NoArg(true, ""), ox.Default(args.addFixed("usql.echo_hidden=on")), ox.Short("E"), ox.Section(sectionInputOutput)).
		String("output", "send query results to file",
			ox.Short("o"), ox.Spec("FILE"), ox.Bind(&args.Out), ox.HiddenAliases("out"), ox.Section(sectionInputOutput)).
		Var("quiet", "run quietly (no messages, only query output)",
			ox.HookT, ox.NoArg(true, ""), ox.Default(args.addFixed("usql.quiet=on")), ox.Short("q"), ox.Section(sectionInputOutput))
	// output format options, in psql's order. Each sets the settings of
	// entries, and an entry with a verb takes the value of the option.
	for _, o := range []struct {
		name, short, usage, spec string
		entries                  []string
	}{
		{"no-align", "A", "unaligned table output mode", "", []string{"format=unaligned"}},
		{"csv", "C", "CSV output mode", "", []string{"format=csv"}},
		{"field-separator", "F", `field separator for unaligned and CSV output (default "|" and ",")`, "STRING", []string{"fieldsep=%q", "csv_fieldsep=%q"}},
		{"vertical", "G", "vertical output mode", "", []string{"format=vertical"}},
		{"html", "H", "HTML table output mode", "", []string{"format=html"}},
		{"json", "J", "JSON output mode", "", []string{"format=json"}},
		{"markdown", "M", "markdown output mode", "", []string{"format=markdown"}},
		{"pset", "P", `set printing option VAR to ARG (see \pset command)`, "VAR=ARG", nil},
		{"record-separator", "R", `record separator for unaligned and CSV output (default \n)`, "STRING", []string{"recordsep=%q"}},
		{"tuples-only", "t", "print rows only", "", []string{"tuples_only=on"}},
		{"table-attr", "T", "set HTML table tag attributes (e.g., width, border)", "TEXT", []string{"tableattr=%q"}},
		{"expanded", "x", "turn on expanded table output", "", []string{"expanded=on"}},
		{"field-separator-zero", "z", "set field separator for unaligned and CSV output to zero byte", "", []string{"fieldsep_zero=on"}},
		{"record-separator-zero", "0", "set record separator for unaligned and CSV output to zero byte", "", []string{"recordsep_zero=on"}},
	} {
		hook := args.addOption("print")
		if o.entries != nil {
			entries := make([]string, len(o.entries))
			for i, e := range o.entries {
				entries[i] = "print." + e
			}
			hook = args.addFixed(entries...)
		}
		opts := []ox.Option{ox.HookT, ox.Default(hook), ox.Short(o.short), ox.Section(sectionFormat)}
		if o.spec == "" {
			opts = append(opts, ox.NoArg(true, ""))
		} else {
			opts = append(opts, ox.Spec(o.spec))
		}
		flags = flags.Var(o.name, o.usage, opts...)
	}
	// connection options
	flags = flags.
		Bool("no-password", "never prompt for password",
			ox.Short("w"), ox.Bind(&args.NoPassword), ox.Section(sectionConnection)).
		Bool("password", "force password prompt (should happen automatically)",
			ox.Short("W"), ox.Bind(&args.ForcePassword), ox.Section(sectionConnection))
	// the options of --dump are long options only, and hidden, because the
	// short options of psql's output use -t, -T and -N
	for _, o := range []struct {
		name, spec, usage string
	}{
		{"target", "DRIVER", "write the dump for another database, such as duckdb"},
		{"schema-only", "", "dump the schema and no rows"},
		{"data-only", "", "dump the rows and no schema"},
		{"schema", "PATTERN", "dump the schemas that match PATTERN"},
		{"exclude-schema", "PATTERN", "do not dump the schemas that match PATTERN"},
		{"table", "PATTERN", "dump the tables that match PATTERN"},
		{"exclude-table", "PATTERN", "do not dump the tables that match PATTERN"},
		{"allow-lossy", "", "allow a type of the target that holds less than the source type"},
		{"allow-text", "", "allow the string type of the target for a type that it does not have"},
		{"restrict-key", "KEY", `the key of the \restrict line of the dump`},
	} {
		opts := []ox.Option{ox.HookT, ox.Default(args.addDump(o.name)), ox.Hidden(true)}
		if o.spec == "" {
			opts = append(opts, ox.NoArg(true, "true"))
		} else {
			opts = append(opts, ox.Spec(o.spec))
		}
		flags = flags.Var(o.name, o.usage, opts...)
	}
	// psql's -h is the host, which usql takes in the DSN, so -h fails at once
	flags = flags.Var("bad-help", "",
		ox.HookT, ox.NoArg(true, ""), ox.Short("h"), ox.Hidden(true),
		ox.Default(func(string) error { return ox.ErrUnknownFlag }))
	return []ox.Option{
		ox.Usage("usql", "the universal command-line interface for SQL databases"),
		ox.Banner("usql, the universal command-line interface for SQL databases"),
		ox.Spec("[OPTION]... [DSN]"),
		ox.VersionString("0.0.0-dev"),
		ox.Version(),
		ox.Comp(),
		ox.Help(
			// psql starts each description at column 27
			ox.HelpColumn(27),
			ox.Sections(
				"General options",
				"Input and output options",
				"Output format options",
				"Connection options",
			),
			ox.Footer(`For more information, type "\?" (for internal commands) or "\help" (for SQL
commands) from within usql, or see https://github.com/xo/usql.`),
		),
		ox.ArgsFunc(func(v []string) error {
			if len(v) > 1 {
				return errWrongNumberOfArguments
			}
			return nil
		}),
		flags,
	}
}

// addCommandOrFile returns the hook of -c or -f, which adds the value to the
// commands and files in the order of the command line.
func (args *Args) addCommandOrFile(kind string) func(string) error {
	return func(value string) error {
		args.CommandOrFiles = append(args.CommandOrFiles, kind+":"+value)
		return nil
	}
}

// addOption returns the hook of an option that sets SECTION.NAME=VALUE. With
// no section, the value names its section.
func (args *Args) addOption(section string) func(string) error {
	return func(value string) error {
		if section != "" {
			value = section + "." + value
		}
		args.Options = append(args.Options, value)
		return nil
	}
}

// addFixed returns the hook of an option that sets the settings of entries.
// An entry with a verb, such as print.fieldsep=%q, takes the value of the
// option.
func (args *Args) addFixed(entries ...string) func(string) error {
	return func(value string) error {
		for _, e := range entries {
			if strings.Contains(e, "%") {
				e = fmt.Sprintf(e, value)
			}
			args.Options = append(args.Options, e)
		}
		return nil
	}
}

// addDump returns the hook of an option of a dump.
func (args *Args) addDump(name string) func(string) error {
	return func(value string) error {
		args.DumpOptions = append(args.DumpOptions, name+"="+value)
		return nil
	}
}

// isUsageError returns true when err is an error of the command line, after
// which usql suggests --help.
func isUsageError(err error) bool {
	for _, target := range []error{
		errWrongNumberOfArguments,
		ox.ErrUnknownFlag,
		ox.ErrMissingArgument,
		ox.ErrInvalidValue,
		ox.ErrInvalidArgCount,
	} {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
