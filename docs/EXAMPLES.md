# Examples for coding agents

This document holds short examples of `ox`, by task. Find the task that you
are doing, and copy its example. D32 decides this document.

Each example is a Go `Example` func in a test file, so that `go test` runs
it and compares what it prints with its `// Output:` comment:

- `examples_test.go` in the root holds the examples that need only `ox`.
- `test/example_test.go` holds the examples that need the `yaml`, the
  `toml` or the `hcl` module.

`TestEachExampleIsInTheDocument` in `repo_test.go` fails when an example in
this document differs from its func in a test file. If you change an
example, change it in both places.

An example passes `ox.Args(...)`, so that the test can give it a command
line. A program leaves out `ox.Args`, and then `ox` reads `os.Args`. A
program imports `github.com/xo/ox`, and `github.com/xo/ox/otx` when it reads
a value from the context.

Read [Known faults](#known-faults) before you copy an example into a
program.

## Contents

1. [Choose a style](#choose-a-style)
2. [Write struct tags](#write-struct-tags)
3. [Build commands and help](#build-commands-and-help)
4. [Read a value of each type](#read-a-value-of-each-type)
5. [Read config files and the environment](#read-config-files-and-the-environment)
6. [Test a command](#test-a-command)
7. [Known faults](#known-faults)

## Choose a style

`ox` has four styles. Each style builds the same command tree, so you can
use more than one style in a program.

### Build the flags by hand

`ox.Flags()` returns a flag set, and each method adds a flag of one type.
The exec func reads the values with the `otx` package.

```go
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
```

### Bind a flag to a variable

`ox.Bind` sets a variable each time that the flag gets a value. `ox.BindSet`
also sets a bool to true when the command line gives the flag.

```go
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
```

### Build the flags from a struct

`ox.From` adds a flag for each field that has an `ox` tag. The tag starts
with the description, and options follow, separated by commas. The name of
the flag is the name of the field in kebab case.

```go
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
```

### Build a program from a struct and its doc comments

`ox.App` builds a command from a struct. The doc comment of the struct type
gives the help, in sections such as `Usage:`, `Description:`, `Banner:`,
`Examples:` and `Footer:`. The first line of `Usage:` holds the names of the
parent commands, the name of the command and its spec. Each field with an
`ox` tag is a flag, as with `ox.From`. When the tag has no description, the
doc comment of the field is the usage, so a comma needs no escape. A method
`Commands() []any` gives the sub commands, and a method `Exec` gives the exec
func. `ox.AppFrom(&app)` runs the program, with `ox.Defaults()`.

Go keeps no doc comment in a program. Add this line to the program, and run
`go generate`:

```text
//go:generate go run github.com/xo/ox/cmd/oxdoc
```

`oxdoc` writes `oxdoc_gen.go`, which calls `ox.RegisterDoc` for each struct
type that has a doc comment. This example calls `ox.RegisterDoc` itself, and
declares the struct type inside the func, which `oxdoc` does not read.

```go
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
```
### Read the values from the context

`otx.Get[T]` reads a value as the type `T`, and it converts the value when it
can. `otx.Slice[E]` and `otx.Map[K, V]` read a slice and a map. A sub command
reads the flags of its parents, and the args that are left.

```go
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
```

## Write struct tags

A tag of `From` has the form `ox:"description,option,key:value"`. Each
option after the description is one of these:

| Option | What it does |
| --- | --- |
| `name:NAME` | sets the name of the flag |
| `short:C` | sets the short name, one character |
| `alias:NAME` | adds one alias |
| `aliases:A\|B` | adds more than one alias, separated by `\|` |
| `spec:TEXT` | sets the text after the flag in the help, such as `FILE` |
| `default:VALUE` | sets the default, which can hold variables |
| `noarg:VALUE` | makes the value optional, and gives the value without one |
| `type:TYPE` | sets the type, such as `count`, `date` or `path` |
| `mapkey:TYPE` | sets the type of the keys of a map |
| `elem:TYPE` | sets the type of the elements of a slice or a map |
| `split:SEP` | sets the separator of a slice or a map, `,` by default |
| `set:FIELD` | sets the bool field `FIELD` to true when the command line gives the flag |
| `section:N` | puts the flag in section `N` of the help |
| `hidden` | hides the flag in the help |
| `deprecated` | hides the flag in the help |
| `hook:help` or `hook:version` | makes the flag the help or the version flag |
| `key:KEY` or `key:TYPE\|KEY` | sets the config key, for every format or for one format |
| `config` | makes the flag name the config file |
| `env:NAME` | fills the flag from the environment variable `NAME` |
| `noenv` | keeps the flag out of the environment |
| `noconfig` | keeps the flag out of the config files |

An option that is not in this table is an error.

### Set the name, the help and the section of a flag

The `hook:help` and `hook:version` options replace the flags that `Help` and
`Version` add. A hidden or a deprecated flag is not in the help.

```go
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
```

### Set the type and the value of a flag

`type`, `mapkey` and `elem` set the type that `ox` reads, and the Go field
gets the text of the value. `noarg` gives the value of a flag without a
value. `set` names a bool field that becomes true when the command line
gives the flag. A one-character alias is a short alias, such as `-u`.

```go
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
```

### Escape a comma, a quote and a backslash

A tag is a Go string inside backquotes. Go removes one level of
backslashes, and then `ox` removes one more. A comma ends the description or
an option, so a comma in the text needs `\\,`:

| To get this | Write this in the tag | Why |
| --- | --- | --- |
| `read a file, then exit` | `ox:"read a file\\, then exit"` | `\\,` becomes `\,` for `ox`, which keeps the comma |
| `host:port to listen on` | `ox:"host:port to listen on"` | A colon in the description needs no escape |
| `say "hi", loudly` | `ox:"say \"hi\"\\, loudly"` | A double quote ends the tag value, so Go needs `\"` |
| `a back\slash` | `ox:"a back\\\\slash"` | Each level removes one backslash |
| the default `http://x:80/a,b` | `ox:"a url,default:http://x:80/a\\,b"` | The first colon of an option ends the key, and the rest is the value |

```go
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
```

### The error of a comma that is not escaped

A comma that is not escaped starts an option. The text after it is not an
option, so `From` gives `ox.ErrUnknownTagOption`, with the name of the field.

```go
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
```

## Build commands and help

### Add sub commands and aliases

`ox.Sub` adds a sub command, with the same options as the root. `ox.Aliases`
gives a command more names. `ox.ValidArgs(1, 1)` requires exactly one arg.

```go
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
```

### Write the help of a sub command

`ox.Defaults()` adds a `help` command when the root has sub commands.
`app help remote` and `app remote --help` write the same help.

```go
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
```

### Hide a flag or a command

`ox.Hidden(true)` and `ox.Deprecated(true)` keep a flag or a command out of
the help. The command line still accepts the flag. A deprecated flag writes
a warning to standard error when the command line uses it.

```go
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
```

### Put the flags in sections

Inside `ox.Defaults` or `ox.Help`, `ox.Sections` names the sections of the
flags. `ox.Section(n)` puts a flag in section `n`. A flag without a section
goes in "Flags", before the named sections.

Outside `ox.Help`, `ox.Sections` names the sections of the sub commands, and
`ox.Section(n)` on a sub command puts it in section `n`.

```go
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
```

### Set the banner, the spec, the examples and the footer

These are options of `ox.Defaults` or `ox.Help`. `ox.Spec` replaces the text
after the command name in "Usage".

```go
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
```

### Describe the positional arguments

`ox.ArgUsage(name, usage)` writes a positional argument in an `Arguments:`
section of the help. `Command.UsageString` returns the help as a string. It
expands no variable in a default, because it has no context.

```go
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
```

### Put the help and version flags in a section of their own

A flag with `ox.Special("hook:help")` becomes the help flag, and a flag with
`ox.Special("hook:version")` becomes the version flag. Give them a short
name and a section, and pass `ox.Help()` and `ox.Version()`. `ox.Defaults()`
is `ox.Version()`, `ox.Comp()` and `ox.Help()`, in that order.

```go
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
```

### Write the version

`ox.VersionString` sets the version. Without it, `ox` reads the version of
the Go module from the build information.

```go
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
```

### Write a help like psql

This example builds the sections of psql by hand. The help flag is `-?`,
because psql uses `-h` for the host. `ox` adds `-h` to the help flag only
when no other flag has `-h`.

```go
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
```

## Read a value of each type

### A slice

A slice flag takes a value each time, and it splits each value at a comma.
`ox.Split` sets another separator. A type option, such as `ox.IntT`, sets the
type of the elements.

```go
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
```

### A map

A map flag takes `key=value` pairs. `ox.MapKey` sets the type of the keys,
and a type option, such as `ox.IntT`, sets the type of the values.

```go
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
```

### A count

A count adds one for each use of the flag.

```go
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
```

### A size and a rate

`ox.Size` is a number of bytes, such as `1.5GiB` or `10MB`. `ox.Rate` is a
size for each unit of time, such as `10MiB/s`.

```go
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
```

### A duration

A duration has the syntax of `time.ParseDuration`.

```go
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
```

### A file or a directory

`File` makes a flag for a file that exists, and `Dir` makes a flag for a
directory that exists. `Path` makes a flag for any path. The value of each
is a `string`. `ox` checks only a value from the command line, never a
default. Each one expands a leading `~` to the home directory. These options
change a path flag:

1. `ox.Abs()` makes the path absolute.
2. `ox.Real()` resolves each symbolic link.
3. `ox.Home()` shows a default under the home directory as `~` in help.
4. `ox.Ext(".yaml", ".yml")` completes only files with these extensions.
5. `ox.NotExist()` makes a path that must not exist yet.

In a struct tag, write `type:file` or `type:dir`, and `abs`, `real`, `home`,
`notexist` or `ext:.yaml|.yml`.

```go
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
```

### A URL

```go
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
```

### A date and a time

`Date` reads `2006-01-02`, `DateTime` reads `2006-01-02 15:04:05`, and
`Timestamp` reads RFC 3339.

```go
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
```

### A hook that runs for each value

A flag of type `ox.HookT` calls its default func for each use, in the order
of the command line. The func can take the value as a string. psql's
`-v NAME=VALUE` needs this, because the order of the values counts.

```go
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
```

### An optional value

`ox.NoArg(true, value)` makes the value optional. `--color` takes the value
of `NoArg`, and `--color=never` takes `never`. `--color never` does not take
`never`, which stays an arg.

```go
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
```

### Show an optional value in help

`ox.OptionalValue(spec)` makes the value optional, as `ox.NoArg` does, and
the help writes the spec in brackets, such as `--color[=WHEN]`. Without a
value, the flag takes the empty string. `ox.HelpValue(value, usage)` writes a
value of the flag on a line of its own, as psql writes `--help=commands`.

```go
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
```

### Allowed values

`ox.Valid` allows only the values that it names. Another value is an error
that wraps `ox.ErrInvalidValue`.

```go
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
```

## Read config files and the environment

D34 proposes these rules. Each flag takes its value from the first source
that has one:

1. The command line.
2. The environment, when the root has `ox.EnvPrefix` or the flag has
   `ox.EnvVar`.
3. The config files. The last file that has the key wins.
4. The default of the flag.

The root package reads JSON. A blank import of `github.com/xo/ox/yaml` or
`github.com/xo/ox/toml` adds YAML or TOML. The extension of a file gives its
format.

### Read config files

`ox.ConfigFiles` names the files, in order. The key of a flag is its name.
`ox.Key("", "server.port")` reads a key inside a map. A list in a file sets
a slice.

```go
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
```

### Name the config file with a flag

`ox.ConfigFlag()` makes a flag name the config file. Its file replaces the
files of `ox.ConfigFiles`. If its file does not exist, `Parse` gives
`ox.ErrConfigNotFound`.

```go
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
```

### Read the environment

`ox.EnvPrefix("APP")` fills the flag `--log-level` from `APP_LOG_LEVEL`.
`ox.EnvVar` names the variable of one flag. `ox.NoEnv()` keeps a flag out of
the environment.

```go
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
```

### Give a flag a config key

`ox.Key("", key)` sets the key for every format. `ox.Key("json", key)` sets
the key for one format. `ox.NoConfig()` keeps a flag out of the config files,
for a value such as a password.

```go
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
```

### The order of the sources

```go
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
```

### Use variables in a default

A default can hold variables, such as `$HOME`, `$APPNAME`, `$APPCONFIG` and
`${env::NAME}`. A filter after `|` changes the value, such as `|lower`. `\$`
is a `$` that does not start a variable.

```go
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
```

### Read a YAML file

A key that starts with `$` is a YAML path. `ox.Key` takes the path as it is.
A default reads a key with `${yaml::key}`, and it escapes the `$` of a YAML
path as `\$`. A dotted key needs no escape.

```go
func Example_yaml() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.yaml")
	data := "server:\n  port: 8080\nstore:\n  book:\n    - author: jon\n    - author: ken\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		// a program uses a path such as "$APPCONFIG/config.yaml"
		ox.ConfigFiles(path),
		ox.Flags().
			Int("port", "the port", ox.Key("yaml", "server.port")).
			String("first", "the first author", ox.Key("yaml", "$.store.book[0].author")).
			String("authors", "the authors", ox.Default(`${yaml::\$.store.book[*].author}`)).
			String("listen", "the address", ox.Default("localhost:${yaml::server.port}")),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("first:", otx.String(ctx, "first"))
			fmt.Println("authors:", otx.String(ctx, "authors"))
			fmt.Println("listen:", otx.String(ctx, "listen"))
		}),
		ox.Args(),
	)
	// Output:
	// port: 8080
	// first: jon
	// authors: jon,ken
	// listen: localhost:8080
}
```

### Read a YAML path in a struct tag

In a struct tag, the `\$` of a default needs `\\\\$`. Go removes one
backslash, and `ox` removes one more. `key:yaml|server.port` sets the key for
YAML only.

```go
func Example_yamlTag() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.yaml")
	data := "server:\n  port: 8080\nstore:\n  book:\n    - author: jon\n    - author: ken\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	args := struct {
		Port    int    `ox:"the port,key:yaml|server.port"`
		Authors string `ox:"the authors,default:${yaml::\\\\$.store.book[*].author}"`
	}{}
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.From(&args),
		ox.Exec(func() {
			fmt.Println("port:", args.Port)
			fmt.Println("authors:", args.Authors)
		}),
		ox.Args(),
	)
	// Output:
	// port: 8080
	// authors: jon,ken
}
```

### Read a TOML file

```go
func Example_toml() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.toml")
	data := "region = \"us\"\ntags = [\"a\", \"b\"]\n\n[server]\nport = 8080\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.Flags().
			Int("port", "the port", ox.Key("", "server.port")).
			Slice("tags", "the tags").
			String("region", "the region", ox.Default("${toml::region|upper}"), ox.NoConfig()),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("tags:", otx.Slice[string](ctx, "tags"))
			fmt.Println("region:", otx.String(ctx, "region"))
		}),
		ox.Args(),
	)
	// Output:
	// port: 8080
	// tags: [a b]
	// region: US
}
```

### Read an HCL file

Import `github.com/xo/ox/hcl` for its side effect. A block nests its body
under its type, then under each of its labels, so the key of `port` in
`server "db" { port = 5432 }` is `server.db.port`. An expression in the file
can use no variable and no function.

```go
func Example_hcl() {
	dir, err := os.MkdirTemp("", "app")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "config.hcl")
	data := "region = \"us\"\ntags = [\"a\", \"b\"]\n\nserver \"db\" {\n  port = 5432\n}\n"
	_ = os.WriteFile(path, []byte(data), 0o600)
	ox.Run(
		ox.Usage("app", "an app"),
		ox.ConfigFiles(path),
		ox.Flags().
			Int("port", "the port", ox.Key("", "server.db.port")).
			Slice("tags", "the tags").
			String("region", "the region", ox.Default("${hcl::region|upper}"), ox.NoConfig()),
		ox.Exec(func(ctx context.Context) {
			fmt.Println("port:", otx.Int(ctx, "port"))
			fmt.Println("tags:", otx.Slice[string](ctx, "tags"))
			fmt.Println("region:", otx.String(ctx, "region"))
		}),
		ox.Args(),
	)
	// Output:
	// port: 5432
	// tags: [a b]
	// region: US
}
```

## Test a command

### Return the error of a run

`ox.RunError` runs a command as `ox.Run` does, and returns the error. It
does not write the error and does not exit, so the program chooses the
message and the exit status. After `--help` or `--version`, it returns nil.

```go
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
```

### Run a command in a test

`ox.NewContext` builds the context without a run. `ox.Pipe` sets the
standard input, output and error. The exec func writes to the `Stdout` of
the context, which it gets with `ox.Ctx`. Then the test calls `Parse`,
`Validate` and `Run`, the same steps as `ox.Run`.

Do not call `ox.Run` in a test for a command line that fails. `ox.Run`
writes the error and exits the process. Call `ox.RunError` in its place.

```go
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
```

### Check the errors

Each error wraps a constant of `ox`, so check it with `errors.Is`. `Parse`
gives the errors of the flags, and `Validate` gives the errors of the args.
After `--help` writes the help, `Parse` gives `ox.ErrExit`.

```go
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
```

### Test a command with a table

A func that returns the options of a command lets `main` and the test build
the same command. Call the func for each test, because a flag set must not
be used by two commands.

```go
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
```

```go
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
```

## Known faults

The examples found these faults on 2026-11-26. The backlog agent fixed the
others on 2026-11-27 (D36). Until a fault is fixed, do what the last column
says.

| Fault | What you see | Do this |
| --- | --- | --- |
| `ox.Hook` takes a `func(context.Context) error`, which a hook cannot call | `invalid hook func func(context.Context) error` | Use a flag of type `ox.HookT`, or `FlagSet.Hook` |
| The default layout of the help does not write the aliases of a flag | `--username` and not `--username, --login` | Pass `ox.HelpColumn(n)`, which writes them |
