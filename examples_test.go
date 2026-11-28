package ox_test

// The examples in this file are the examples of docs/EXAMPLES.md. Each one
// is short, so that a coding agent can copy it. Change an example here and in
// docs/EXAMPLES.md at the same time.

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

// ExampleRun builds a command by hand, with [ox.Flags], and reads the values
// of the flags in the exec func.
func ExampleRun() {
	ox.Run(
		ox.Usage("greet", "say hello"),
		ox.Defaults(),
		ox.Flags().
			String("name", "the name to greet", ox.Short("n"), ox.Default("world")).
			Int("times", "the number of greetings", ox.Short("t"), ox.Default("1")),
		ox.Exec(func(ctx context.Context) error {
			for range otx.Int(ctx, "times") {
				fmt.Println("hello", otx.String(ctx, "name"))
			}
			return nil
		}),
		// a program leaves out Args, and ox reads os.Args
		ox.Args("--name", "ken", "-t", "2"),
	)
	// Output:
	// hello ken
	// hello ken
}

// ExampleBind binds each flag to a variable.
func ExampleBind() {
	var (
		name   string
		port   int
		tags   []string
		u      *url.URL
		urlSet bool
	)
	ox.Run(
		ox.Usage("serve", "serve files"),
		ox.Defaults(),
		ox.Flags().
			String("name", "the server name", ox.Bind(&name)).
			Int("port", "the port", ox.Default("8080"), ox.Bind(&port)).
			Slice("tag", "a tag", ox.Bind(&tags)).
			URL("upstream", "the upstream URL", ox.BindSet(&u, &urlSet)),
		ox.Exec(func() {
			fmt.Println("name:", name)
			fmt.Println("port:", port)
			fmt.Println("tags:", tags)
			fmt.Println("upstream set:", urlSet)
		}),
		ox.Args("--name", "files", "--tag", "a", "--tag", "b"),
	)
	// Output:
	// name: files
	// port: 8080
	// tags: [a b]
	// upstream set: false
}

// ExampleFrom builds the flags from the fields of a struct, with reflection.
func ExampleFrom() {
	args := struct {
		Name  string   `ox:"the name to greet,short:n,default:world"`
		Times int      `ox:"the number of greetings,short:t,default:1"`
		Tags  []string `ox:"a tag,short:T"`
	}{}
	ox.Run(
		ox.Usage("greet", "say hello"),
		ox.Defaults(),
		ox.From(&args),
		ox.Exec(func() {
			fmt.Println("name:", args.Name)
			fmt.Println("times:", args.Times)
			fmt.Println("tags:", args.Tags)
		}),
		ox.Args("-n", "ken", "-T", "a,b", "-T", "c"),
	)
	// Output:
	// name: ken
	// times: 1
	// tags: [a b c]
}

// Example_context reads the values of the flags from the context of the exec
// func, with the generic funcs of the otx package. A sub command reads the
// flags of its parents too.
func Example_context() {
	ox.Run(
		ox.Usage("tool", "a tool"),
		ox.Defaults(),
		ox.Flags().
			Count("verbose", "verbosity", ox.Short("v")).
			Slice("port", "a port", ox.IntT, ox.Short("p")),
		ox.Sub(
			ox.Usage("run", "run a job"),
			ox.Flags().
				Map("env", "an environment variable", ox.Short("e")),
			ox.Exec(func(ctx context.Context, args []string) error {
				fmt.Println("verbose:", otx.Get[int](ctx, "verbose"))
				fmt.Println("ports:", otx.Slice[int](ctx, "port"))
				fmt.Println("ports as strings:", otx.Slice[string](ctx, "port"))
				fmt.Println("env:", otx.Map[string, string](ctx, "env"))
				fmt.Println("args:", args)
				return nil
			}),
		),
		ox.Args("-v", "-p", "80,443", "run", "-e", "A=1", "-e", "B=2", "job1"),
	)
	// Output:
	// verbose: 1
	// ports: [80 443]
	// ports as strings: [80 443]
	// env: map[A:1 B:2]
	// args: [job1]
}

// ExampleFrom_tags shows each struct tag option that sets the name, the help
// or the section of a flag.
func ExampleFrom_tags() {
	args := struct {
		Host    string `ox:"the host,short:H,spec:HOST,section:0"`
		Port    int    `ox:"the port,short:p,default:5432,section:0"`
		User    string `ox:"the user name,name:username,section:1"`
		Debug   bool   `ox:"debug output,hidden"`
		Old     string `ox:"an old flag,deprecated"`
		Help    bool   `ox:"show this help\\, then exit,short:?,hook:help,section:1"`
		Version bool   `ox:"show the version\\, then exit,short:V,hook:version,section:1"`
	}{}
	ox.Run(
		ox.Usage("db", "connect to a database"),
		ox.Version(),
		ox.Help(ox.Sections("Connection options", "Other options")),
		ox.From(&args),
		ox.Args("--help"),
	)
	// Output:
	// db connect to a database
	//
	// Usage:
	//   db [flags] [args]
	//
	// Connection options:
	//   -H, --host HOST        the host
	//   -p, --port int         the port (default: 5432)
	//
	// Other options:
	//       --username string  the user name
	//   -?, --help             show this help, then exit
	//   -V, --version          show the version, then exit
}

// ExampleFrom_values shows each struct tag option that sets the type, the
// value or the aliases of a flag.
func ExampleFrom_values() {
	args := struct {
		Dir     string            `ox:"a directory,type:path"`
		Weights map[string]string `ox:"weights,mapkey:int,elem:float64"`
		Hosts   []string          `ox:"host addresses,elem:addr"`
		Path    []string          `ox:"a search path,split::"`
		Color   string            `ox:"when to use color,noarg:always,default:auto"`
		Level   string            `ox:"a log level,set:LevelSet"`
		User    string            `ox:"the user name,alias:login,aliases:usr|u"`

		LevelSet bool
	}{}
	ox.Run(
		ox.Usage("app", "an app"),
		ox.From(&args),
		ox.Exec(func() {
			fmt.Println("dir:", args.Dir)
			fmt.Println("weights:", args.Weights)
			fmt.Println("hosts:", args.Hosts)
			fmt.Println("path:", args.Path)
			fmt.Println("color:", args.Color)
			fmt.Println("level set:", args.LevelSet)
			fmt.Println("user:", args.User)
		}),
		ox.Args(
			"--dir", "/tmp",
			"--weights", "1=0.5", "--weights", "2=1.5",
			"--hosts", "10.0.0.1,10.0.0.2",
			"--path", "/bin:/usr/bin",
			"--color",
			"--level", "debug",
			"-u", "ken",
		),
	)
	// Output:
	// dir: /tmp
	// weights: map[1:0.5 2:1.5]
	// hosts: [10.0.0.1 10.0.0.2]
	// path: [/bin /usr/bin]
	// color: always
	// level set: true
	// user: ken
}

// ExampleFrom_args makes fields positional arguments with the tag option
// `arg`, in the order of the fields.
func ExampleFrom_args() {
	args := struct {
		Verbose bool     `ox:"write more,short:v"`
		Src     string   `ox:"the file to copy,arg,type:file"`
		Dsts    []string `ox:"the directories to copy to,arg,optional,elem:dir"`
	}{}
	ox.Run(
		ox.Usage("cp", "copy a file"),
		ox.From(&args),
		ox.Exec(func() {
			fmt.Println("verbose:", args.Verbose)
			fmt.Println("src:", args.Src)
			fmt.Println("dsts:", args.Dsts)
		}),
		ox.Args("-v", "go.mod", "docs", "testdata"),
	)
	// Output:
	// verbose: true
	// src: go.mod
	// dsts: [docs testdata]
}

// ExampleFrom_escaping shows how to escape a comma, a double quote and a
// backslash in a struct tag.
func ExampleFrom_escaping() {
	args := struct {
		File   string `ox:"read a file\\, then exit"`
		Listen string `ox:"host:port to listen on"`
		Say    string `ox:"say \"hi\"\\, loudly"`
		Slash  string `ox:"a back\\\\slash"`
		URL    string `ox:"a url,default:http://x:80/a\\,b"`
	}{}
	ox.Run(
		ox.Usage("esc", "escapes"),
		ox.Help(),
		ox.From(&args),
		ox.Args("--help"),
	)
	// Output:
	// esc escapes
	//
	// Usage:
	//   esc [flags] [args]
	//
	// Flags:
	//       --file string    read a file, then exit
	//       --listen string  host:port to listen on
	//       --say string     say "hi", loudly
	//       --slash string   a back\slash
	//       --url string     a url (default: http://x:80/a,b)
	//   -h, --help           show help, then exit
}

// ExampleFrom_unescapedComma shows the error that a comma gives when it is
// not escaped.
func ExampleFrom_unescapedComma() {
	args := struct {
		File string `ox:"read a file, then exit"`
	}{}
	_, err := ox.NewContext(
		ox.Usage("esc", "escapes"),
		ox.From(&args),
	)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ox.ErrUnknownTagOption))
	// Output:
	// From (*ox.Command): cannot build flags for *struct: field File: unknown tag option: " then exit"
	// true
}

// ExampleSub builds a tree of sub commands. A sub command has aliases, and
// it reads the flags of its parents.
func ExampleSub() {
	ox.Run(
		ox.Usage("git", "a version control tool"),
		ox.Defaults(),
		ox.Flags().
			String("dir", "the work directory", ox.Short("C"), ox.Default(".")),
		ox.Sub(
			ox.Usage("remote", "manage remotes"),
			ox.Sub(
				ox.Usage("remove", "remove a remote"),
				ox.Aliases("rm"),
				ox.ValidArgs(1, 1),
				ox.Exec(func(ctx context.Context, args []string) error {
					fmt.Println("dir:", otx.String(ctx, "dir"))
					fmt.Println("remove:", args[0])
					return nil
				}),
			),
		),
		ox.Args("-C", "/src", "remote", "rm", "origin"),
	)
	// Output:
	// dir: /src
	// remove: origin
}

// ExampleSub_help writes the help of a sub command. The help command and the
// --help flag write the same text.
func ExampleSub_help() {
	ox.Run(
		ox.Usage("git", "a version control tool"),
		ox.Defaults(),
		ox.Sub(
			ox.Usage("remote", "manage remotes"),
			ox.Sub(
				ox.Usage("remove", "remove a remote"),
				ox.Aliases("rm"),
			),
			ox.Sub(
				ox.Usage("add", "add a remote"),
			),
		),
		ox.Args("help", "remote"),
	)
	// Output:
	// remote manage remotes
	//
	// Usage:
	//   git remote [flags] [command] [args]
	//
	// Available Commands:
	//   add     add a remote
	//   remove  remove a remote
	//
	// Flags:
	//   -h, --help      show help, then exit
	//
	// Use "git remote [command] --help" for more information about a command.
}

// ExampleBefore runs a func before the exec func of a command and of each sub
// command, and a func after it. The after func runs also when the exec func
// fails.
func ExampleBefore() {
	err := ox.RunError(
		context.Background(),
		ox.Usage("db", "a database tool"),
		ox.Flags().
			String("dsn", "the database to open", ox.Default("app.db")),
		ox.Before(func(ctx context.Context) error {
			fmt.Println("open", otx.String(ctx, "dsn"))
			return nil
		}),
		ox.After(func(ctx context.Context) error {
			fmt.Println("close", otx.String(ctx, "dsn"))
			return nil
		}),
		ox.Sub(
			ox.Usage("list", "list the rows"),
			ox.Before(func() {
				fmt.Println("lock the table")
			}),
			ox.Exec(func() error {
				fmt.Println("list the rows")
				return errors.New("no rows")
			}),
		),
		ox.Args("--dsn", "test.db", "list"),
	)
	fmt.Println("error:", err)
	// Output:
	// open test.db
	// lock the table
	// list the rows
	// close test.db
	// error: no rows
}

// ExampleHidden hides a flag and a command from the help. A hidden or a
// deprecated flag still works on the command line.
func ExampleHidden() {
	run := func(args ...string) {
		ox.Run(
			ox.Usage("app", "an app"),
			ox.Defaults(),
			ox.Flags().
				Bool("debug", "debug output", ox.Hidden(true)).
				String("old-name", "use --name", ox.Deprecated(true)).
				String("name", "the name"),
			ox.Sub(
				ox.Usage("serve", "serve files"),
			),
			ox.Sub(
				ox.Usage("internal", "internal tools"),
				ox.Hidden(true),
			),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("debug:", otx.Bool(ctx, "debug"))
				fmt.Println("old-name:", otx.String(ctx, "old-name"))
			}),
			ox.Args(args...),
			// the warning of a deprecated flag goes to standard error
			ox.Pipe(os.Stdin, os.Stdout, os.Stdout),
		)
	}
	run("--help")
	run("--debug", "--old-name", "x")
	// Output:
	// app an app
	//
	// Usage:
	//   app [flags] [command] [args]
	//
	// Available Commands:
	//   completion  generate completion script for a specified shell
	//   help        show help for any command
	//   serve       serve files
	//   version     show app version information
	//
	// Flags:
	//       --name string  the name
	//   -h, --help         show help, then exit
	//
	// Use "app [command] --help" for more information about a command.
	// warning: --old-name is deprecated
	// debug: true
	// old-name: x
}

// ExampleSections puts the flags in sections of the help.
func ExampleSections() {
	ox.Run(
		ox.Usage("db", "connect to a database"),
		ox.Defaults(
			ox.Sections("Connection options", "Output options"),
		),
		ox.Flags().
			String("host", "the host", ox.Section(0)).
			Int("port", "the port", ox.Default("5432"), ox.Section(0)).
			Bool("csv", "write CSV", ox.Section(1)).
			Bool("json", "write JSON", ox.Section(1)).
			String("log", "the log file"),
		ox.Args("--help"),
	)
	// Output:
	// db connect to a database
	//
	// Usage:
	//   db [flags] [args]
	//
	// Flags:
	//       --log string   the log file
	//   -v, --version      show version, then exit
	//   -h, --help         show help, then exit
	//
	// Connection options:
	//       --host string  the host
	//       --port int     the port (default: 5432)
	//
	// Output options:
	//       --csv          write CSV
	//       --json         write JSON
}

// ExampleHelp sets the banner, the usage spec, the examples and the footer of
// the help.
func ExampleHelp() {
	ox.Run(
		ox.Usage("fetch", "fetch a URL"),
		ox.Defaults(
			ox.Banner("fetch gets a URL, and writes the body."),
			ox.Spec("[OPTION]... URL"),
			ox.Example("  fetch https://example.com\n  fetch -o out.html https://example.com"),
			ox.Footer("Report bugs to <bugs@example.com>."),
		),
		ox.Flags().
			String("output", "the output file", ox.Short("o"), ox.Spec("FILE")),
		ox.Args("--help"),
	)
	// Output:
	// fetch gets a URL, and writes the body.
	//
	// Usage:
	//   fetch [OPTION]... URL
	//
	// Examples:
	//   fetch https://example.com
	//   fetch -o out.html https://example.com
	//
	// Flags:
	//   -o, --output FILE  the output file
	//   -v, --version      show version, then exit
	//   -h, --help         show help, then exit
	//
	// Report bugs to <bugs@example.com>.
}

// Example_psqlHelp builds a help that looks like the help of psql, by hand.
// The help and version flags are in the "General options" section, with
// psql's short names.
func Example_psqlHelp() {
	ox.Run(
		ox.Usage("psql", "the PostgreSQL interactive terminal"),
		ox.VersionString("18.1"),
		ox.Version(),
		ox.Help(
			ox.Banner("psql is the PostgreSQL interactive terminal."),
			ox.Spec("[OPTION]... [DBNAME [USERNAME]]"),
			ox.Sections(
				"General options",
				"Connection options",
			),
			ox.Footer("PostgreSQL home page: <https://www.postgresql.org/>"),
		),
		ox.Flags().
			String("command", "run only single command (SQL or internal) and exit", ox.Short("c"), ox.Section(0)).
			String("file", "execute commands from file, then exit", ox.Short("f"), ox.Spec("FILENAME"), ox.Section(0)).
			Bool("version", "output version information, then exit", ox.Short("V"), ox.Special("hook:version"), ox.Section(0)).
			Bool("help", "show this help, then exit", ox.Short("?"), ox.Special("hook:help"), ox.Section(0)).
			String("host", "database server host or socket directory", ox.Short("h"), ox.Spec("HOSTNAME"), ox.Section(1)).
			Uint("port", "database server port", ox.Short("p"), ox.Spec("PORT"), ox.Default("5432"), ox.Section(1)),
		ox.Args("-?"),
	)
	// Output:
	// psql is the PostgreSQL interactive terminal.
	//
	// Usage:
	//   psql [OPTION]... [DBNAME [USERNAME]]
	//
	// General options:
	//   -c, --command string  run only single command (SQL or internal) and exit
	//   -f, --file FILENAME   execute commands from file, then exit
	//   -V, --version         output version information, then exit
	//   -?, --help            show this help, then exit
	//
	// Connection options:
	//   -h, --host HOSTNAME   database server host or socket directory
	//   -p, --port PORT       database server port (default: 5432)
	//
	// PostgreSQL home page: <https://www.postgresql.org/>
}

// ExampleSpecial puts the help and version flags in a section of their own.
// A flag with Special("hook:help") or Special("hook:version") replaces the
// flag that Help or Version adds.
func ExampleSpecial() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.VersionString("1.2.3"),
		ox.Version(),
		ox.Help(ox.Sections("Options", "Help options")),
		ox.Flags().
			String("name", "the name", ox.Section(0)).
			Bool("help", "show this help, then exit", ox.Short("?"), ox.Special("hook:help"), ox.Section(1)).
			Bool("version", "show the version, then exit", ox.Special("hook:version"), ox.Section(1)),
		ox.Args("-?"),
	)
	// Output:
	// app an app
	//
	// Usage:
	//   app [flags] [args]
	//
	// Options:
	//       --name string  the name
	//
	// Help options:
	//   -?, --help         show this help, then exit
	//       --version      show the version, then exit
}

// ExampleVersion writes the version of the program. Version adds --version
// to a root without sub commands. It adds a version command to a root that
// has sub commands.
func ExampleVersion() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.VersionString("1.2.3"),
		ox.Version(),
		ox.Help(),
		ox.Args("--version"),
	)
	// Output:
	// app 1.2.3
}

// ExampleLocal keeps a flag to its command, so that a sub command refuses
// it. A sub command inherits every other flag of its parents.
func ExampleLocal() {
	run := func(args ...string) error {
		return ox.RunError(
			context.Background(),
			ox.Usage("app", "an app"),
			ox.Flags().
				String("config", "the config file").
				Bool("init", "make a new config file", ox.Local()),
			ox.Sub(
				ox.Usage("show", "show the config"),
				ox.Exec(func(ctx context.Context) {
					fmt.Println("config:", otx.String(ctx, "config"))
				}),
			),
			ox.Args(args...),
		)
	}
	fmt.Println(run("show", "--config", "a.yaml"))
	err := run("show", "--init")
	fmt.Println(errors.Is(err, ox.ErrUnknownFlag), err)
	// Output:
	// config: a.yaml
	// <nil>
	// true --init: unknown flag
}

// ExampleInheritedFlags writes the flags that a sub command inherits in its
// help, in a section of their own. A local flag is not inherited.
func ExampleInheritedFlags() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Defaults(ox.InheritedFlags(true)),
		ox.Flags().
			String("config", "the config file", ox.Short("c")).
			Bool("init", "make a new config file", ox.Local()),
		ox.Sub(
			ox.Usage("show", "show the config"),
			ox.Flags().
				Bool("all", "show every key"),
		),
		ox.Args("show", "--help"),
	)
	// Output:
	// show show the config
	//
	// Usage:
	//   app show [flags] [args]
	//
	// Flags:
	//       --all       show every key
	//   -h, --help      show help, then exit
	//
	// Global Flags:
	//   -c, --config string  the config file
}

// ExampleDefaultWrapWidth sets the width at which help wraps. A program that
// wants the width of the terminal reads it, and sets DefaultWrapWidth.
func ExampleDefaultWrapWidth() {
	ox.DefaultWrapWidth = 50
	defer func() {
		ox.DefaultWrapWidth = 95
	}()
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Help(),
		ox.Flags().
			String("name", "the name to greet, which help wraps over more than one line", ox.Short("n")),
		ox.Args("--help"),
	)
	// Output:
	// app an app
	//
	// Usage:
	//   app [flags] [args]
	//
	// Flags:
	//   -n, --name string  the name to greet, which help
	//                      wraps over more than one line
	//   -h, --help         show help, then exit
}

// ExampleStopAtArg ends the flags of a command at its first positional
// argument, so that a wrapper command passes the rest to the command that it
// runs.
func ExampleStopAtArg() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Bool("verbose", "write more", ox.Short("v")),
		ox.Sub(
			ox.Usage("exec", "run a command"),
			ox.StopAtArg(),
			ox.Flags().
				String("dir", "the work directory", ox.Short("C")),
			ox.Exec(func(ctx context.Context, args []string) {
				fmt.Println("verbose:", otx.Bool(ctx, "verbose"))
				fmt.Println("dir:", otx.String(ctx, "dir"))
				fmt.Printf("args: %q\n", args)
			}),
		),
		ox.Args("exec", "-v", "-C", "/src", "ls", "-l", "--all"),
	)
	// Output:
	// verbose: true
	// dir: /src
	// args: ["ls" "-l" "--all"]
}

// ExampleAbbrev takes a unique prefix of a long flag, as GNU getopt_long
// does. A prefix that matches more than one flag is an error.
func ExampleAbbrev() {
	run := func(args ...string) error {
		return ox.RunError(
			context.Background(),
			ox.Usage("app", "an app"),
			ox.Abbrev(),
			ox.Flags().
				Bool("verbose", "write more").
				Bool("version", "write the version").
				String("output", "the output file"),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("output:", otx.String(ctx, "output"))
				fmt.Println("verbose:", otx.Bool(ctx, "verbose"))
			}),
			ox.Args(args...),
		)
	}
	fmt.Println(run("--out", "a.txt", "--verb"))
	err := run("--ver")
	fmt.Println(errors.Is(err, ox.ErrAmbiguousFlag), err)
	// Output:
	// output: a.txt
	// verbose: true
	// <nil>
	// true --ver: ambiguous flag: --verbose, --version
}

// ExampleNegateBools gives each bool flag a negated form, such as
// --no-color, and help writes the flag as --[no-]color.
func ExampleNegateBools() {
	run := func(args ...string) {
		ox.Run(
			ox.Usage("app", "an app"),
			ox.Help(),
			ox.NegateBools("no", "skip"),
			ox.Flags().
				Bool("color", "use color", ox.Default(true)).
				Bool("cache", "use the cache", ox.Default(true)).
				Bool("no-rc", "do not read the startup file").
				Bool("force", "overwrite each file", ox.Negatable(false)),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("color:", otx.Bool(ctx, "color"))
				fmt.Println("cache:", otx.Bool(ctx, "cache"))
			}),
			ox.Args(args...),
		)
	}
	run("--no-color", "--skip-cache")
	run("--help")
	// Output:
	// color: false
	// cache: false
	// app an app
	//
	// Usage:
	//   app [flags] [args]
	//
	// Flags:
	//       --[no-]color  use color
	//       --[no-]cache  use the cache
	//       --no-rc       do not read the startup file
	//       --force       overwrite each file
	//   -h, --help        show help, then exit
}

// ExampleFlagSet_Slice reads a flag more than one time, or a list split by a
// comma.
func ExampleFlagSet_Slice() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Slice("tag", "a tag", ox.Short("t")).
			Slice("port", "a port", ox.IntT).
			Slice("path", "a search path", ox.Split(":")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("tags:", otx.Slice[string](ctx, "tag"))
			fmt.Println("ports:", otx.Slice[int](ctx, "port"))
			fmt.Println("path:", otx.Slice[string](ctx, "path"))
		}),
		ox.Args("-t", "a", "-t", "b,c", "--port", "80,443", "--path", "/bin:/usr/bin"),
	)
	// Output:
	// tags: [a b c]
	// ports: [80 443]
	// path: [/bin /usr/bin]
}

// ExampleFlagSet_Map reads key=value pairs.
func ExampleFlagSet_Map() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Map("label", "a label", ox.Short("l")).
			Map("limit", "a limit", ox.IntT),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("labels:", otx.Map[string, string](ctx, "label"))
			fmt.Println("limits:", otx.Map[string, int](ctx, "limit"))
		}),
		ox.Args("-l", "a=1", "-l", "b=2,c=3", "--limit", "cpu=2"),
	)
	// Output:
	// labels: map[a:1 b:2 c:3]
	// limits: map[cpu:2]
}

// ExampleFlagSet_Count counts each use of a flag, such as -vvv.
func ExampleFlagSet_Count() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Count("verbose", "verbosity", ox.Short("v")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("verbose:", otx.Count(ctx, "verbose"))
		}),
		ox.Args("-vvv", "--verbose"),
	)
	// Output:
	// verbose: 4
}

// ExampleFlagSet_Size reads a size in bytes, and a rate in bytes per unit of
// time.
func ExampleFlagSet_Size() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Size("max", "the maximum size", ox.Default("1MiB")).
			Rate("limit", "the rate limit"),
		ox.Exec(func(ctx context.Context) {
			size := otx.Get[ox.Size](ctx, "max")
			rate := otx.Get[ox.Rate](ctx, "limit")
			fmt.Println("max:", size, int64(size))
			fmt.Println("limit:", rate)
		}),
		ox.Args("--max", "1.5GiB", "--limit", "10MiB/s"),
	)
	// Output:
	// max: 1.5 GiB 1610612736
	// limit: 10 MiB/s
}

// ExampleSize_Format writes a size in the exact form of a Kubernetes quantity.
func ExampleSize_Format() {
	for _, s := range []string{"6Gi", "1.5Gi", "1500Mi", "6e9", "100"} {
		size, err := ox.ParseSize(s)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %q % q %Q %e\n", s, size, size, size, size)
	}
	rate, err := ox.ParseRate("1.5Gi/s")
	if err != nil {
		panic(err)
	}
	fmt.Printf("rate: %q\n", rate)
	// Output:
	// 6Gi: 6Gi 6 Gi 6442450944 6Gi
	// 1.5Gi: 1536Mi 1536 Mi 1610612736 1.5Gi
	// 1500Mi: 1500Mi 1500 Mi 1572864k 1.46Gi
	// 6e9: 5859375Ki 5859375 Ki 6G 5.59Gi
	// 100: 100 100 100 100
	// rate: 1536Mi/s
}

// ExampleFlagSet_Duration reads a duration, such as 1m30s.
func ExampleFlagSet_Duration() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Duration("timeout", "the timeout", ox.Default("30s")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("timeout:", otx.Duration(ctx, "timeout"))
		}),
		ox.Args("--timeout", "1m30s"),
	)
	// Output:
	// timeout: 1m30s
}

// ExampleFlagSet_File reads a file and a directory that must exist.
func ExampleFlagSet_File() {
	run := func(args ...string) {
		err := ox.RunError(
			context.Background(),
			ox.Usage("app", "an app"),
			ox.Flags().
				File("config", "the config file", ox.Ext(".yaml", ".yml")).
				Dir("out", "the output directory"),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("config:", otx.String(ctx, "config"), "out:", otx.String(ctx, "out"))
			}),
			ox.Args(args...),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
	run("--config", "go.mod", "--out", "docs")
	run("--config", "missing.yaml")
	run("--config", "docs")
	run("--out", "go.mod")
	// Output:
	// config: go.mod out: docs
	// --config: invalid value: missing.yaml: does not exist
	// --config: invalid value: docs: is a directory
	// --out: invalid value: go.mod: not a directory
}

// ExampleFlagSet_URL reads a URL.
func ExampleFlagSet_URL() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			URL("endpoint", "the endpoint"),
		ox.Exec(func(ctx context.Context) {
			u := otx.URL(ctx, "endpoint")
			fmt.Println("host:", u.Host)
			fmt.Println("path:", u.Path)
		}),
		ox.Args("--endpoint", "https://example.com:8443/api"),
	)
	// Output:
	// host: example.com:8443
	// path: /api
}

// ExampleFlagSet_PortRange reads a port range, and a slice of port ranges.
func ExampleFlagSet_PortRange() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			PortRange("ports", "the ports to listen on", ox.Default("8000-8099")).
			Slice("allow", "the ports to allow", ox.PortRangeT),
		ox.Exec(func(ctx context.Context) {
			ports := otx.Get[ox.PortRange](ctx, "ports")
			fmt.Println("ports:", ports, ports.Lo, ports.Hi, ports.Contains(8080))
			for _, r := range otx.Slice[ox.PortRange](ctx, "allow") {
				fmt.Println("allow:", r, slices.Collect(r.Ports()))
			}
		}),
		ox.Args("--ports", "8080:8081", "--allow", "22,443-445"),
	)
	// Output:
	// ports: 8080-8081 8080 8081 true
	// allow: 22 [22]
	// allow: 443-445 [443 444 445]
}

// ExampleTypeFor gives a flag the type slog.Level, which nothing registers.
// The flag reads its value with the UnmarshalText of slog.Level, and help
// writes its type in lower case.
func ExampleTypeFor() {
	run := func(args ...string) {
		ox.Run(
			ox.Usage("app", "an app"),
			ox.Help(),
			ox.Flags().
				Var("log", "the log level", ox.TypeFor[slog.Level](), ox.Default("info")),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("log:", otx.Get[slog.Level](ctx, "log"))
			}),
			ox.Args(args...),
		)
	}
	run("--log", "debug")
	run("--help")
	// Output:
	// log: DEBUG
	// app an app
	//
	// Usage:
	//   app [flags] [args]
	//
	// Flags:
	//       --log level  the log level (default: info)
	//   -h, --help       show help, then exit
}

// ExampleFlagSet_Date reads a date, a date and time, and a timestamp.
func ExampleFlagSet_Date() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Date("since", "the first day").
			DateTime("at", "a date and time").
			Timestamp("until", "an RFC 3339 timestamp"),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("since:", otx.TimeTime(ctx, "since").Format(time.DateOnly))
			fmt.Println("at:", otx.TimeTime(ctx, "at").Format(time.DateTime))
			fmt.Println("until:", otx.TimeTime(ctx, "until").Format(time.RFC3339))
		}),
		ox.Args("--since", "2026-11-01", "--at", "2026-11-26 09:30:00", "--until", "2026-12-01T00:00:00Z"),
	)
	// Output:
	// since: 2026-11-01
	// at: 2026-11-26 09:30:00
	// until: 2026-12-01T00:00:00Z
}

// ExampleStrftime reads a time in a strftime layout.
func ExampleStrftime() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Var("since", "the first day", ox.Strftime("%d/%m/%Y"), ox.Default("01/11/2026")).
			Var("at", "a date and time", ox.Strftime("%d %b %Y %I:%M %p")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("since:", otx.Get[time.Time](ctx, "since").Format(time.DateOnly))
			fmt.Println("at:", otx.Time(ctx, "at"), "is", otx.TimeTime(ctx, "at").Format(time.DateTime))
		}),
		ox.Args("--at", "26 Nov 2026 09:30 PM"),
	)
	// Output:
	// since: 2026-11-01
	// at: 26 Nov 2026 09:30 PM is 2026-11-26 21:30:00
}

// ExampleFlagSet_Hook calls a func for each use of a flag, in the order of
// the command line. The flag takes a value, because it has no NoArg.
func ExampleFlagSet_Hook() {
	var vars []string
	set := func(s string) {
		vars = append(vars, s)
	}
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			Var("set", "set variable NAME to VALUE", ox.HookT, ox.Default(set), ox.Short("v")).
			String("file", "a file", ox.Short("f")),
		ox.Exec(func() {
			fmt.Println("vars:", vars)
		}),
		ox.Args("--set", "a=1", "-f", "x.sql", "-v", "b=2", "--set=c=3"),
	)
	// Output:
	// vars: [a=1 b=2 c=3]
}

// ExampleNoArg gives a flag an optional value. Without a value, the flag
// takes the value of NoArg. With --flag=value, it takes that value.
func ExampleNoArg() {
	run := func(args ...string) {
		ox.Run(
			ox.Usage("app", "an app"),
			ox.Flags().
				String("color", "when to use color", ox.NoArg(true, "always"), ox.Default("auto")),
			ox.Exec(func(ctx context.Context, args []string) {
				fmt.Println("color:", otx.String(ctx, "color"), "args:", args)
			}),
			ox.Args(args...),
		)
	}
	run()
	run("--color")
	run("--color=never")
	run("--color", "never")
	// Output:
	// color: auto args: []
	// color: always args: []
	// color: never args: []
	// color: always args: [never]
}

// ExampleOptionalValue shows an optional value in the help, and writes some
// of its values on lines of their own.
func ExampleOptionalValue() {
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			String("color", "when to use color", ox.OptionalValue("WHEN"),
				ox.HelpValue("always", "always use color"),
				ox.HelpValue("never", "never use color"),
			).
			Bool("quiet", "run quietly", ox.Short("q")),
		ox.Help(),
		ox.Args("--help"),
	)
	// Output:
	// app an app
	//
	// Usage:
	//   app [flags] [args]
	//
	// Flags:
	//       --color[=WHEN]  when to use color
	//       --color=always  always use color
	//       --color=never   never use color
	//   -q, --quiet         run quietly
	//   -h, --help          show help, then exit
}

// ExampleHelpValue writes a help topic of its own for a value of the help
// flag. Without a value, the help flag writes the help of the command.
func ExampleHelpValue() {
	commands := func(ctx *ox.Context) error {
		fmt.Fprintln(ctx.Stdout, `\q  quit`)
		fmt.Fprintln(ctx.Stdout, `\?  show the commands`)
		return nil
	}
	run := func(args ...string) {
		ox.Run(
			ox.Usage("app", "an app"),
			ox.Flags().
				Bool("help", "show this help, then exit", ox.Special("hook:help"), ox.OptionalValue("TOPIC"),
					ox.HelpValue("commands", "list the commands, then exit", commands),
				),
			ox.Help(),
			ox.Args(args...),
		)
	}
	run("--help=commands")
	run("--help")
	// Output:
	// \q  quit
	// \?  show the commands
	// app an app
	//
	// Usage:
	//   app [flags] [args]
	//
	// Flags:
	//   --help[=TOPIC]   show this help, then exit
	//   --help=commands  list the commands, then exit
}

// ExampleValid allows only some values for a flag.
func ExampleValid() {
	ctx, err := ox.NewContext(
		ox.Usage("app", "an app"),
		ox.Flags().
			String("format", "the output format", ox.Valid("json", "yaml")),
		ox.Args("--format", "xml"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ctx.Parse()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ox.ErrInvalidValue))
	// Output:
	// --format: invalid value: "xml": allowed values: "json", "yaml"
	// true
}

// ExampleConfigFiles reads the flags from config files. A later file wins
// over an earlier one, and a file that does not exist is skipped.
func ExampleConfigFiles() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	system := filepath.Join(dir, "system.json")
	user := filepath.Join(dir, "user.json")
	_ = os.WriteFile(system, []byte(`{"log-level": "info", "server": {"port": 8080}, "tags": ["a", "b"]}`), 0o600)
	_ = os.WriteFile(user, []byte(`{"log-level": "debug"}`), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		// a program uses a path such as "$APPCONFIG/config.json"
		ox.ConfigFiles(system, filepath.Join(dir, "missing.json"), user),
		ox.Flags().
			String("log-level", "the log level", ox.Default("warn")).
			Int("port", "the port", ox.Key("", "server.port")).
			Slice("tags", "the tags"),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("log-level:", otx.String(ctx, "log-level"))
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("tags:", otx.Slice[string](ctx, "tags"))
		}),
		ox.Args(),
	)
	// Output:
	// log-level: debug
	// port: 8080
	// tags: [a b]
}

// ExampleConfigFlag names the config file with a flag. The file of the flag
// replaces the files of ConfigFiles, and it must exist.
func ExampleConfigFlag() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "other.json")
	_ = os.WriteFile(path, []byte(`{"name": "from other.json"}`), 0o600)
	newContext := func(args ...string) (*ox.Context, error) {
		return ox.NewContext(
			ox.Usage("app", "an app"),
			ox.ConfigFiles("$APPCONFIG/config.json"),
			ox.Flags().
				String("config", "the config file", ox.Short("c"), ox.ConfigFlag()).
				String("name", "the name"),
			ox.Args(args...),
		)
	}
	// the flag names a file
	ctx, err := newContext("-c", path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ctx.Parse(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("name:", ctx.Vars["name"])
	// the flag names a file that does not exist
	ctx, err = newContext("-c", filepath.Join(dir, "missing.json"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ctx.Parse()
	fmt.Println(errors.Is(err, ox.ErrConfigNotFound))
	// Output:
	// name: from other.json
	// true
}

// ExampleEnvPrefix fills each flag from an environment variable. EnvPrefix
// names the variable of each flag. EnvVar gives one flag a variable of its
// own, and NoEnv keeps a flag out of the environment.
func ExampleEnvPrefix() {
	_ = os.Setenv("APP_LOG_LEVEL", "debug")
	_ = os.Setenv("API_TOKEN", "secret")
	_ = os.Setenv("APP_HOME", "/from/env")
	defer func() { _ = os.Unsetenv("APP_LOG_LEVEL") }()
	defer func() { _ = os.Unsetenv("API_TOKEN") }()
	defer func() { _ = os.Unsetenv("APP_HOME") }()
	ox.Run(
		ox.Usage("app", "an app"),
		ox.EnvPrefix("APP"),
		ox.Flags().
			String("log-level", "the log level", ox.Default("info")).
			String("token", "the API token", ox.EnvVar("API_TOKEN")).
			String("home", "the home directory", ox.NoEnv(), ox.Default("/default")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("log-level:", otx.String(ctx, "log-level"))
			fmt.Println("token:", otx.String(ctx, "token"))
			fmt.Println("home:", otx.String(ctx, "home"))
		}),
		ox.Args(),
	)
	// Output:
	// log-level: debug
	// token: secret
	// home: /default
}

// ExampleKey gives a flag a config key that is not its name, for every
// format or for one format. NoConfig keeps a flag out of the config files.
func ExampleKey() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.json")
	_ = os.WriteFile(path, []byte(`{"db": {"host": "db1", "port": 5433}, "password": "in-file"}`), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.Flags().
			String("host", "the host", ox.Key("", "db.host")).
			Int("port", "the port", ox.Key("json", "db.port"), ox.Key("yaml", "database.port")).
			String("password", "the password", ox.NoConfig()),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("host:", otx.String(ctx, "host"))
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Printf("password: %q\n", otx.String(ctx, "password"))
		}),
		ox.Args(),
	)
	// Output:
	// host: db1
	// port: 5433
	// password: ""
}

// ExampleConfigFiles_subCommand reads the key of a flag of a sub command,
// under the path of the command. A field of a nested struct reads the path
// of the field, and an inherited flag keeps the key of its command.
func ExampleConfigFiles_subCommand() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.json")
	data := `{"verbose": true, "db": {"migrate": {"steps": 3, "source": {"dir": "./up"}}}}`
	_ = os.WriteFile(path, []byte(data), 0o600)
	var args struct {
		Steps  int `ox:"the number of steps"`
		Source struct {
			Dir string `ox:"the folder of the migrations"`
		} `ox:""`
	}
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.Flags().
			Bool("verbose", "write more"),
		ox.Sub(
			ox.Usage("db", "the database"),
			ox.Sub(
				ox.Usage("migrate", "migrate the database"),
				// --steps reads db.migrate.steps, and --source-dir reads
				// db.migrate.source.dir
				ox.From(&args),
				ox.Exec(func(ctx context.Context) {
					fmt.Println("verbose:", otx.Bool(ctx, "verbose"))
					fmt.Println("steps:", args.Steps)
					fmt.Println("dir:", args.Source.Dir)
				}),
			),
		),
		ox.Args("db", "migrate"),
	)
	// Output:
	// verbose: true
	// steps: 3
	// dir: ./up
}

// Example_sources shows the order of the sources of a value. The command
// line wins over the environment. The environment wins over the config
// files, and the config files win over the default.
func Example_sources() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.json")
	_ = os.WriteFile(path, []byte(`{"level": "from-file"}`), 0o600)
	run := func(args ...string) {
		ox.Run(
			ox.Usage("app", "an app"),
			ox.EnvPrefix("APP"),
			ox.ConfigFiles(path),
			ox.Flags().
				String("level", "the level", ox.Default("from-default")),
			ox.Exec(func(ctx context.Context) {
				fmt.Println(otx.String(ctx, "level"))
			}),
			ox.Args(args...),
		)
	}
	_ = os.Setenv("APP_LEVEL", "from-env")
	run("--level", "from-command-line")
	run()
	_ = os.Unsetenv("APP_LEVEL")
	run()
	_ = os.Remove(path)
	run()
	// Output:
	// from-command-line
	// from-env
	// from-file
	// from-default
}

// Example_defaultVariables expands variables in a default. A \$ in a
// default is a $ that does not start a variable.
func Example_defaultVariables() {
	_ = os.Setenv("APP_DIR", "/srv/app")
	defer func() { _ = os.Unsetenv("APP_DIR") }()
	ox.Run(
		ox.Usage("app", "an app"),
		ox.Flags().
			String("dir", "the directory", ox.Default("${env::APP_DIR}/data")).
			String("rc", "the startup file", ox.Default("/etc/.${APPNAME|lower}rc")).
			String("price", "the price", ox.Default(`\$5`)),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("dir:", otx.String(ctx, "dir"))
			fmt.Println("rc:", otx.String(ctx, "rc"))
			fmt.Println("price:", otx.String(ctx, "price"))
		}),
		ox.Args(),
	)
	// Output:
	// dir: /srv/app/data
	// rc: /etc/.apprc
	// price: $5
}

// ExampleRequired makes a flag required. A value from the command line, the
// environment or a config file counts, and a default does not. Help marks
// the flag, and --help works without it.
func ExampleRequired() {
	run := func(args ...string) error {
		return ox.RunError(
			context.Background(),
			ox.Usage("deploy", "deploy the app"),
			ox.Help(),
			ox.EnvPrefix("DEPLOY"),
			ox.Flags().
				String("target", "the target", ox.Required()).
				String("region", "the region", ox.Default("east")),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("deploy to", otx.String(ctx, "target"))
			}),
			ox.Args(args...),
		)
	}
	err := run()
	fmt.Println(errors.Is(err, ox.ErrRequiredFlag), err)
	fmt.Println(run("--target", "prod"))
	_ = os.Setenv("DEPLOY_TARGET", "staging")
	defer func() { _ = os.Unsetenv("DEPLOY_TARGET") }()
	fmt.Println(run())
	fmt.Println(run("--help"))
	// Output:
	// true command deploy: required flag: --target
	// deploy to prod
	// <nil>
	// deploy to staging
	// <nil>
	// deploy deploy the app
	//
	// Usage:
	//   deploy [flags] [args]
	//
	// Flags:
	//       --target string  the target (required)
	//       --region string  the region (default: east)
	//   -h, --help           show help, then exit
	// <nil>
}

// ExampleOnlyOneOf adds rules between the flags of a command. OnlyOneOf
// allows one flag of the group at most. AllOrNoneOf needs each flag of the
// group, or none of them. AtLeastOneOf needs one flag of the group at least.
func ExampleOnlyOneOf() {
	run := func(args ...string) error {
		return ox.RunError(
			context.Background(),
			ox.Usage("export", "export the data"),
			ox.Flags().
				Bool("json", "write JSON").
				Bool("yaml", "write YAML").
				String("user", "the user").
				String("pass", "the password").
				String("file", "the file to read").
				String("url", "the URL to read"),
			ox.OnlyOneOf("json", "yaml"),
			ox.AllOrNoneOf("user", "pass"),
			ox.AtLeastOneOf("file", "url"),
			ox.Exec(func() {
				fmt.Println("export")
			}),
			ox.Args(args...),
		)
	}
	fmt.Println(run("--file", "data.csv", "--json"))
	err := run("--file", "data.csv", "--json", "--yaml")
	fmt.Println(errors.Is(err, ox.ErrOnlyOneOf), err)
	err = run("--url", "https://example.com", "--user", "ken")
	fmt.Println(errors.Is(err, ox.ErrAllOrNoneOf), err)
	err = run("--json")
	fmt.Println(errors.Is(err, ox.ErrAtLeastOneOf), err)
	// Output:
	// export
	// <nil>
	// true command export: only one of the flags can be set: --json, --yaml
	// true command export: all or none of the flags must be set: --user, --pass (missing --pass)
	// true command export: at least one of the flags must be set: --file, --url
}

// ExampleFrom_rules makes a flag required, and adds rules between flags,
// with struct tags. The flags with the same rule and the same group name are
// one group.
func ExampleFrom_rules() {
	args := struct {
		Target string `ox:"the target,required"`
		JSON   bool   `ox:"write JSON,onlyoneof:format"`
		YAML   bool   `ox:"write YAML,onlyoneof:format"`
		User   string `ox:"the user,allornoneof:auth"`
		Pass   string `ox:"the password,allornoneof:auth"`
	}{}
	run := func(v ...string) error {
		return ox.RunError(
			context.Background(),
			ox.Usage("deploy", "deploy the app"),
			ox.From(&args),
			ox.Exec(func() {
				fmt.Println("deploy to", args.Target)
			}),
			ox.Args(v...),
		)
	}
	fmt.Println(run("--target", "prod", "--json"))
	fmt.Println(run("--json"))
	fmt.Println(run("--target", "prod", "--json", "--yaml"))
	fmt.Println(run("--target", "prod", "--pass", "secret"))
	// Output:
	// deploy to prod
	// <nil>
	// command deploy: required flag: --target
	// command deploy: only one of the flags can be set: --json, --yaml
	// command deploy: all or none of the flags must be set: --user, --pass (missing --user)
}

// ExampleNewContext tests a command. The test builds a context with the
// command line to test, and a buffer for the output. The exec func writes to
// the Stdout of the context, so that the test reads it.
func ExampleNewContext() {
	var stdout strings.Builder
	ctx, err := ox.NewContext(
		ox.Usage("greet", "say hello"),
		ox.Flags().
			String("name", "the name to greet", ox.Default("world")),
		ox.ValidArgs(0, 0),
		ox.Exec(func(ctx context.Context) error {
			c, _ := ox.Ctx(ctx)
			fmt.Fprintln(c.Stdout, "hello", otx.String(ctx, "name"))
			return nil
		}),
		ox.Pipe(strings.NewReader(""), &stdout, &stdout),
		ox.Args("--name", "ken"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ctx.Parse(); err != nil {
		fmt.Println(err)
		return
	}
	if err := ctx.Validate(); err != nil {
		fmt.Println(err)
		return
	}
	if err := ctx.Run(context.Background()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(stdout.String())
	// Output:
	// hello ken
}

// ExampleContext_Parse checks the errors of a command line with errors.Is.
// Parse gives the errors of the flags. Validate gives the errors of the
// args. The --help flag gives ErrExit after it writes the help.
func ExampleContext_Parse() {
	check := func(args ...string) error {
		ctx, err := ox.NewContext(
			ox.Usage("greet", "say hello"),
			ox.Help(),
			ox.Flags().
				String("name", "the name to greet"),
			ox.ValidArgs(0, 0),
			ox.Pipe(strings.NewReader(""), io.Discard, io.Discard),
			ox.Args(args...),
		)
		if err != nil {
			return err
		}
		if err := ctx.Parse(); err != nil {
			return err
		}
		return ctx.Validate()
	}
	err := check("--nam", "ken")
	fmt.Println(errors.Is(err, ox.ErrUnknownFlag), err)
	err = check("--name")
	fmt.Println(errors.Is(err, ox.ErrMissingArgument), err)
	err = check("extra")
	fmt.Println(errors.Is(err, ox.ErrInvalidArgCount), err)
	err = check("--help")
	fmt.Println(errors.Is(err, ox.ErrExit), err)
	// Output:
	// true --nam: unknown flag
	// true --name: missing argument
	// true command greet: invalid arg count: takes no args
	// true --help: exit
}

// greetOptions returns the options of the greet command. A func that
// returns new options lets a test build the command more than one time.
func greetOptions() []ox.Option {
	return []ox.Option{
		ox.Usage("greet", "say hello"),
		ox.Help(),
		ox.Flags().
			String("name", "the name to greet", ox.Default("world")),
		ox.ValidArgs(0, 0),
		ox.Exec(func(ctx context.Context) error {
			c, _ := ox.Ctx(ctx)
			fmt.Fprintln(c.Stdout, "hello", otx.String(ctx, "name"))
			return nil
		}),
	}
}

// TestGreet tests the greet command with a table of command lines.
func TestGreet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args []string
		exp  string
		err  error
	}{
		{"default", nil, "hello world\n", nil},
		{"name", []string{"--name", "ken"}, "hello ken\n", nil},
		{"unknown flag", []string{"--nam", "ken"}, "", ox.ErrUnknownFlag},
		{"extra arg", []string{"extra"}, "", ox.ErrInvalidArgCount},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var stdout strings.Builder
			ctx, err := ox.NewContext(append(
				greetOptions(),
				ox.Pipe(strings.NewReader(""), &stdout, &stdout),
				ox.Args(test.args...),
			)...)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err = ctx.Parse(); err == nil {
				if err = ctx.Validate(); err == nil {
					err = ctx.Run(context.Background())
				}
			}
			switch {
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case stdout.String() != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, stdout.String())
			}
		})
	}
}

// ExampleRunError runs a command and returns its error, so that the program
// chooses the message and the exit status.
func ExampleRunError() {
	run := func(args ...string) error {
		return ox.RunError(
			context.Background(),
			ox.Usage("greet", "say hello"),
			ox.Flags().
				String("name", "the name to greet", ox.Default("world")),
			ox.ValidArgs(0, 0),
			ox.Exec(func(ctx context.Context) error {
				fmt.Println("hello", otx.String(ctx, "name"))
				return nil
			}),
			ox.Args(args...),
		)
	}
	fmt.Println(run("--name", "ken"))
	err := run("--nam", "ken")
	fmt.Println(errors.Is(err, ox.ErrUnknownFlag), err)
	// Output:
	// hello ken
	// <nil>
	// true --nam: unknown flag
}

// ExampleArgUsage describes a positional argument in the help, and gets the
// help as a string.
func ExampleArgUsage() {
	ctx, err := ox.NewContext(
		ox.Usage("sql", "run SQL"),
		ox.Help(
			ox.Spec("[OPTION]... [DSN]"),
			ox.ArgUsage("DSN", "a URL, or the name of a connection"),
		),
		ox.Flags().
			String("command", "run one command, then exit", ox.Short("c")),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(ctx.Root.UsageString())
	// Output:
	// sql run SQL
	//
	// Usage:
	//   sql [OPTION]... [DSN]
	//
	// Arguments:
	//   DSN     a URL, or the name of a connection
	//
	// Flags:
	//   -c, --command string  run one command, then exit
	//   -h, --help            show help, then exit
}

// ExampleArg declares two positional arguments with a type. The second has a
// default, so it is optional.
func ExampleArg() {
	run := func(args ...string) {
		err := ox.RunError(
			context.Background(),
			ox.Usage("cp", "copy a file"),
			ox.Arg("src", "the file to copy", ox.FileT),
			ox.Arg("dst", "the directory to copy to", ox.DirT, ox.Default(".")),
			ox.Exec(func(ctx context.Context) {
				fmt.Println("copy", otx.String(ctx, "src"), "to", otx.String(ctx, "dst"))
			}),
			ox.Args(args...),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
	run("go.mod", "docs")
	run("go.mod")
	run("missing.txt")
	run()
	run("go.mod", "docs", "more")
	// Output:
	// copy go.mod to docs
	// copy go.mod to .
	// command cp: arg src: invalid value: missing.txt: does not exist
	// command cp: missing argument: src
	// command cp: invalid arg count: takes max 2 arg(s)
}

// ExampleArg_variadic declares a last argument that takes each argument that
// is left, and writes the help.
func ExampleArg_variadic() {
	run := func(args ...string) {
		err := ox.RunError(
			context.Background(),
			ox.Usage("cat", "print files"),
			ox.Arg("files", "the files to print", ox.SliceT, ox.Elem(ox.FileT), ox.Spec("FILE")),
			ox.Help(),
			ox.Exec(func(ctx context.Context) {
				fmt.Println(otx.Slice[string](ctx, "files"))
			}),
			ox.Args(args...),
		)
		if err != nil {
			fmt.Println(err)
		}
	}
	run("go.mod", "LICENSE")
	run("--help")
	// Output:
	// [go.mod LICENSE]
	// cat print files
	//
	// Usage:
	//   cat [flags] FILE...
	//
	// Arguments:
	//   FILE    the files to print
	//
	// Flags:
	//   -h, --help      show help, then exit
}

// ExampleApp builds a command from a struct and its doc comments.
func ExampleApp() {
	// Greet says hello.
	//
	// Usage:
	//
	//	greet [flags] [NAME]
	//
	// Footer:
	//
	//	See the docs.
	type Greet struct {
		// the greeting, such as hello, or hi
		Greeting string `ox:",short:g,default:hello"`
	}
	// oxdoc writes this call in oxdoc_gen.go, from go generate
	ox.RegisterDoc[Greet](
		"Greet says hello.\n\nUsage:\n\n\tgreet [flags] [NAME]\n\nFooter:\n\n\tSee the docs.\n",
		map[string]string{"Greeting": "the greeting, such as hello, or hi\n"},
	)
	var greet Greet
	err := ox.RunError(
		context.Background(),
		ox.App(&greet),
		ox.Help(),
		ox.Args("--help"),
	)
	fmt.Println(err)
	// Output:
	// greet
	//
	// Usage:
	//   greet [flags] [NAME]
	//
	// Flags:
	//   -g, --greeting string  the greeting, such as hello, or hi (default: hello)
	//   -h, --help             show help, then exit
	//
	// See the docs.
	// <nil>
}
