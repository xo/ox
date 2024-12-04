package ox_test

import (
	"context"
	"net/url"
	"time"

	"github.com/xo/ox"
)

// Example is a quick example demonstrating the xo/ox package.
func Example() {
	type Verbosity int
	args := struct {
		MyString string    `ox:"a string,short:s"`
		MyBool   bool      `ox:"a bool,short:b"`
		Ints     []int     `ox:"a slice of ints,short:i"`
		Date     time.Time `ox:"formatted date,type:date,short:d"`
		Verbose  Verbosity `ox:"enable verbose,short:v,type:count"`
	}{}
	ox.Run(
		// ox.Exec(myFunc),
		ox.Usage("myApp", "my app"),
		ox.Defaults(
			ox.Banner(`an example ox app.`),
			ox.Footer(`See: https://github.com/xo/ox for more information.`),
		),
		ox.From(&args),
	)
	// Output:
	// an example ox app.
	//
	// Usage:
	//   myApp [flags] [args]
	//
	// Flags:
	//   -s, --my-string string  a string
	//   -b, --my-bool           a bool
	//   -i, --ints int          a slice of ints
	//   -d, --date date         formatted date
	//   -v, --verbose           enable verbose
	//       --version           show version, then exit
	//   -h, --help              show help, then exit
	//
	// See: https://github.com/xo/ox for more information.
}

// Example_argsTest provides an example for testing arbitrary command-line
// invocations by setting the arguments to [Parse] with [ox.Args].
func Example_argsTest() {
	args := struct {
		Number float64 `ox:"a number"`
	}{}
	subArgs := struct {
		URL *url.URL `ox:"a url,short:u"`
	}{}
	ox.RunContext(
		context.Background(),
		// ox.Exec(myFunc),
		ox.Usage("extest", "test example"),
		ox.Defaults(),
		ox.From(&args),
		ox.Sub(
			// ox.Exec(mySubFunc),
			ox.Usage("sub", "a sub command to test"),
			ox.From(&subArgs),
		),
		// the command line args to test
		ox.Args("sub", "--help"),
	)
	// Output:
	// sub a sub command to test
	//
	// Usage:
	//   extest sub [flags] [args]
	//
	// Flags:
	//   -u, --url url  a url
	//   -h, --help     show help, then exit
}

// Example_psql demonstrates building complex help output, based on original
// output of `psql --help`. The output formatting has been slightly changed, as
// the generated help output alters the column formatting, and the output
// cannot be duplicated perfectly -- that said, a faithful attempt has been
// made to stick to the original help output wherever possible.
func Example_psql() {
	args := struct {
		Command           string            `ox:"run only single command (SQL or internal) and exit,short:c,section:0"`
		Dbname            string            `ox:"database name to connect to,short:d,default:$USER,section:0"`
		File              string            `ox:"execute commands from file\\, then exit,short:f,spec:FILENAME,section:0"`
		List              bool              `ox:"list databases\\, then exit,short:l,section:0"`
		Variable          map[string]string `ox:"set psql variable NAME to VALUE,short:v,alias:set,spec:NAME=VALUE,section:0"`
		Version           bool              `ox:"output version information\\, then exit,hook:version,short:V,section:0"`
		NoPsqlrc          bool              `ox:"do not read startup file (~/.psqlrc),short:X,section:0"`
		SingleTransaction bool              `ox:"execute as a single transaction (if non-interactive),short:1,section:0"`
		Help              bool              `ox:"show this help\\, then exit,short:?,hook:help,section:0"`

		EchoAll     bool   `ox:"echo all input from script,short:a,section:1"`
		EchoErrors  bool   `ox:"echo failed commands,short:b,section:1"`
		EchoQueries bool   `ox:"echo commands sent to server,short:e,section:1"`
		EchoHidden  bool   `ox:"disply queries that internal commands generate,short:E,section:1"`
		LogFile     string `ox:"send session log to file,spec:FILENAME,short:L,section:1"`
		NoReadline  bool   `ox:"display enhanced command line editing (readline),short:L,section:1"`
		Output      string `ox:"send query results to file (or |pipe),short:o,section:1"`
		Quiet       bool   `ox:"run quietly (no messages\\, only query output),short:q,section:1"`
		SingleStep  bool   `ox:"single-step mode (confirm each query),short:s,section:1"`
		SingleLine  bool   `ox:"single-line mode (end of line terminates SQL command),short:S,section:1"`

		NoAlign             bool              `ox:"unaligned table output mode,short:A,section:2"`
		CSV                 bool              `ox:"CSV (Comma-Separated Values) table output mode,section:2"`
		FieldSeparator      string            `ox:"field separator for unaligned output,default:|,short:F,section:2"`
		HTML                bool              `ox:"HTML table output mode,short:H,section:2"`
		Pset                map[string]string `ox:"set printing option VAR to ARG (see \\pset command),short:P,spec:VAR[=ARG],section:2"`
		RecordSeparator     string            `ox:"record separator for unaligned output,default:newline,short:R,section:2"`
		TuplesOnly          bool              `ox:"print rows only,short:t,section:2"`
		TableAttr           string            `ox:"set HTML table tag attributes (e.g.\\, width\\, border),short:T,section:2"`
		Expanded            bool              `ox:"turn on expanded table output,short:x,section:2"`
		FieldSeparatorZero  string            `ox:"set field separator for unaligned output to zero byte,short:z,section:2"`
		RecordSeparatorZero string            `ox:"set record separator for unaligned output to zero byte,short:0,section:2"`

		Host       string `ox:"database server host or socket directory,default:local socket,spec:HOSTNAME,short:h,section:3"`
		Port       uint   `ox:"database server port,default:5432,spec:PORT,short:p,section:3"`
		Username   string `ox:"database user name,default:$USER,spec:USERNAME,short:U,section:3"`
		NoPassword bool   `ox:"never prompt for password,short:w,section:3"`
		Password   bool   `ox:"force password prompt (should happen automatically),short:W,section:3"`
	}{}
	ox.Run(
		// ox.Exec(myFunc),
		ox.Usage("psql", "the PostgreSQL interactive terminal"),
		ox.Help(
			ox.Banner("psql is the PostgreSQL interactive terminal."),
			ox.Spec("[OPTION]... [DBNAME [USERNAME]]"),
			ox.Sections(
				"General options",
				"Input and output options",
				"Output format options",
				"Connection options",
			),
			ox.Footer(`For more information, type "\?" (for internal commands) or "\help" (for SQL
commands) from within psql, or consult the psql section in the PostgreSQL
documentation.

Report bugs to <pgsql-bugs@lists.postgresql.org>.
PostgreSQL home page: <https://www.postgresql.org/>`),
		),
		ox.From(&args),
		// this is used to override the expansion parameters
		ox.Override(map[string]string{
			"$USER": "fuser",
		}),
	)
	// Output:
	// psql is the PostgreSQL interactive terminal.
	//
	// Usage:
	//   psql [OPTION]... [DBNAME [USERNAME]]
	//
	// General options:
	//   -c, --command string                run only single command (SQL or internal) and exit
	//   -d, --dbname string                 database name to connect to (default: fuser)
	//   -f, --file FILENAME                 execute commands from file, then exit
	//   -l, --list                          list databases, then exit
	//   -v, --variable NAME=VALUE           set psql variable NAME to VALUE
	//   -V, --version                       output version information, then exit
	//   -X, --no-psqlrc                     do not read startup file (~/.psqlrc)
	//   -1, --single-transaction            execute as a single transaction (if non-interactive)
	//   -?, --help                          show this help, then exit
	//
	// Input and output options:
	//   -a, --echo-all                      echo all input from script
	//   -b, --echo-errors                   echo failed commands
	//   -e, --echo-queries                  echo commands sent to server
	//   -E, --echo-hidden                   disply queries that internal commands generate
	//   -L, --log-file FILENAME             send session log to file
	//   -L, --no-readline                   display enhanced command line editing (readline)
	//   -o, --output string                 send query results to file (or |pipe)
	//   -q, --quiet                         run quietly (no messages, only query output)
	//   -s, --single-step                   single-step mode (confirm each query)
	//   -S, --single-line                   single-line mode (end of line terminates SQL command)
	//
	// Output format options:
	//   -A, --no-align                      unaligned table output mode
	//       --csv                           CSV (Comma-Separated Values) table output mode
	//   -F, --field-separator string        field separator for unaligned output (default: |)
	//   -H, --html                          HTML table output mode
	//   -P, --pset VAR[=ARG]                set printing option VAR to ARG (see pset command)
	//   -R, --record-separator string       record separator for unaligned output (default: newline)
	//   -t, --tuples-only                   print rows only
	//   -T, --table-attr string             set HTML table tag attributes (e.g., width, border)
	//   -x, --expanded                      turn on expanded table output
	//   -z, --field-separator-zero string   set field separator for unaligned output to zero byte
	//   -0, --record-separator-zero string  set record separator for unaligned output to zero byte
	//
	// Connection options:
	//   -h, --host HOSTNAME                 database server host or socket directory (default: local
	//                                       socket)
	//   -p, --port PORT                     database server port (default: 5432)
	//   -U, --username USERNAME             database user name (default: fuser)
	//   -w, --no-password                   never prompt for password
	//   -W, --password                      force password prompt (should happen automatically)
	//
	// For more information, type "\?" (for internal commands) or "\help" (for SQL
	// commands) from within psql, or consult the psql section in the PostgreSQL
	// documentation.
	//
	// Report bugs to <pgsql-bugs@lists.postgresql.org>.
	// PostgreSQL home page: <https://www.postgresql.org/>
}

// Example_sections demonstrates setting the help section for commands and
// flags, including default `--help` flag and `help` command.
func Example_sections() {
	args := struct {
		Config string           `ox:"config file,spec:FILE,section:1"`
		MyInts []int            `ox:"a integer slice,short:i"`
		URLMap map[int]*url.URL `ox:"urls,short:U"`
	}{}
	ox.Run(
		ox.Usage("tree", "a command tree"),
		ox.Defaults(ox.Sections(
			"Normal flags",
			"More flags",
			"Other flags",
		)),
		ox.Sub(
			ox.Usage("sub1", "the sub1 command"),
			ox.Section(0),
		),
		ox.Sub(
			ox.Usage("sub2.a", "the sub2.a command"),
			ox.Section(1),
		),
		ox.Sub(
			ox.Usage("sub2.b", "the sub2.b command"),
		),
		ox.Sections(
			"Primary commands",
			"Secondary commands",
		),
		ox.SectionMap{
			"help":         0,
			"sub2.b":       1,
			"flag:help":    0,
			"flag:my-ints": 2,
		},
		ox.From(&args),
	)
	// Output:
	// tree a command tree
	//
	// Usage:
	//   tree [flags] [command] [args]
	//
	// Primary commands:
	//   sub1    the sub1 command
	//   help    show help for any command
	//
	// Secondary commands:
	//   sub2.a  the sub2.a command
	//   sub2.b  the sub2.b command
	//
	// Flags:
	//   -U, --url-map int=url  urls
	//   -v, --version          show version, then exit
	//
	// Normal flags:
	//   -h, --help             show help, then exit
	//
	// More flags:
	//       --config FILE      config file
	//
	// Other flags:
	//   -i, --my-ints int      a integer slice
	//
	// Use "tree [command] --help" for more information about a command.
}
