package ox_test

import (
	"context"

	"github.com/xo/ox"
)

func Example() {
	args := struct {
		Count uint64 `ox:"type:count"`
	}{}
	run := func() {
	}
	ox.Run(
		ox.Exec(run),
		ox.Defaults(),
		ox.From(&args),
	)
	// Output:
}

func Example_test() {
	args := struct {
		Number float64 `ox:"a number"`
	}{}
	root := func(context.Context, []string) {
	}
	sub := func(context.Context, []string) error {
		return nil
	}
	ox.RunContext(
		context.Background(),
		ox.Exec(root),
		ox.Usage("exhelp", "help example"),
		ox.Help(),
		ox.Sub(
			ox.Exec(sub),
		),
		ox.From(&args),
		ox.Args("--help"),
	)
	// Output:
}

// Example_psql demonstrates building complex help output, based on original
// output of `psql --help`. The output formatting has been slightly changed, as
// the generated help output alters the column formatting, and the output
// cannot be duplicated perfectly -- that said, a faithful attempt has been
// made to stick to the original help output wherever possible.
func Example_psql() {
	args := struct {
		Command string `ox:"run only single command (SQL or internal) and exit,short:c,section:0"`
		// default:ken here forces it to match output, but `default:$USER` will
		// use current user's name as default value
		Dbname   string            `ox:"database name to connect to,short:d,default:ken,section:0"`
		File     string            `ox:"execute commands from file\\, then exit,short:f,spec:FILENAME,section:0"`
		List     bool              `ox:"list databases\\, then exit,short:l,section:0"`
		Variable map[string]string `ox:"set psql variable NAME to VALUE,short:v,alias:set,spec:NAME=VALUE,section:0"`
		Version  bool              `ox:"output version information\\, then exit,section:0"`
	}{
		Dbname: "ken",
	}
	run := func() {
		// empty func
	}
	ox.Run(
		ox.Exec(run),
		ox.Usage("psql", ""),
		ox.Help(
			ox.Banner("psql is the PostgreSQL interactive terminal."),
			ox.Spec("[OPTION]... [DBNAME [USERNAME]]"),
			ox.Sections(
				"General options",
				"Input and output options",
				"Output format options",
				"Connection options",
			),
			ox.Footer(``),
		),
		ox.From(&args),
		ox.Args("--help"),
	)
	// Output:
	// psql is the PostgreSQL interactive terminal.
	//
	// Usage:
	//   psql [OPTION]... [DBNAME [USERNAME]]
	//
	// General options:
	//   -c, --command=COMMAND    run only single command (SQL or internal) and exit
	//   -d, --dbname=DBNAME      database name to connect to (default: "ken")
	//   -f, --file=FILENAME      execute commands from file, then exit
	//   -l, --list               list available databases, then exit
	//   -v, --set=, --variable=NAME=VALUE
	//                            set psql variable NAME to VALUE
	//                            (e.g., -v ON_ERROR_STOP=1)
	//   -V, --version            output version information, then exit
	//   -X, --no-psqlrc          do not read startup file (~/.psqlrc)
	//   -1 ("one"), --single-transaction
	//                            execute as a single transaction (if non-interactive)
	//   -?, --help[=options]     show this help, then exit
	//       --help=commands      list backslash commands, then exit
	//       --help=variables     list special variables, then exit
	//
	// Input and output options:
	//   -a, --echo-all           echo all input from script
	//   -b, --echo-errors        echo failed commands
	//   -e, --echo-queries       echo commands sent to server
	//   -E, --echo-hidden        display queries that internal commands generate
	//   -L, --log-file=FILENAME  send session log to file
	//   -n, --no-readline        disable enhanced command line editing (readline)
	//   -o, --output=FILENAME    send query results to file (or |pipe)
	//   -q, --quiet              run quietly (no messages, only query output)
	//   -s, --single-step        single-step mode (confirm each query)
	//   -S, --single-line        single-line mode (end of line terminates SQL command)
	//
	// Output format options:
	//   -A, --no-align           unaligned table output mode
	//       --csv                CSV (Comma-Separated Values) table output mode
	//   -F, --field-separator=STRING
	//                            field separator for unaligned output (default: "|")
	//   -H, --html               HTML table output mode
	//   -P, --pset=VAR[=ARG]     set printing option VAR to ARG (see \pset command)
	//   -R, --record-separator=STRING
	//                            record separator for unaligned output (default: newline)
	//   -t, --tuples-only        print rows only
	//   -T, --table-attr=TEXT    set HTML table tag attributes (e.g., width, border)
	//   -x, --expanded           turn on expanded table output
	//   -z, --field-separator-zero
	//                            set field separator for unaligned output to zero byte
	//   -0, --record-separator-zero
	//                            set record separator for unaligned output to zero byte
	//
	// Connection options:
	//   -h, --host=HOSTNAME      database server host or socket directory (default: "local socket")
	//   -p, --port=PORT          database server port (default: "5432")
	//   -U, --username=USERNAME  database user name (default: "ken")
	//   -w, --no-password        never prompt for password
	//   -W, --password           force password prompt (should happen automatically)
	//
	// For more information, type "\?" (for internal commands) or "\help" (for SQL
	// commands) from within psql, or consult the psql section in the PostgreSQL
	// documentation.
	//
	// Report bugs to <pgsql-bugs@lists.postgresql.org>.
	// PostgreSQL home page: <https://www.postgresql.org/>
}
