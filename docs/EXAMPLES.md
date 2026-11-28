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
6. [Require flags](#require-flags)
7. [Test a command](#test-a-command)
8. [Known faults](#known-faults)

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
| `local` | keeps the flag to its command, so that a sub command refuses it |
| `negatable` or `negatable:false` | gives a bool flag a negated form such as `--no-debug`, or keeps it out of `NegateBools` |
| `hook:help` or `hook:version` | makes the flag the help or the version flag |
| `key:KEY` or `key:TYPE\|KEY` | sets the config key, for every format or for one format |
| `config` | makes the flag name the config file |
| `env:NAME` | fills the flag from the environment variable `NAME` |
| `noenv` | keeps the flag out of the environment |
| `noconfig` | keeps the flag out of the config files |
| `required` | makes the flag required |
| `onlyoneof:GROUP` | adds the flag to the group `GROUP`, which allows one flag at most |
| `allornoneof:GROUP` | adds the flag to the group `GROUP`, which needs each flag or none |
| `atleastoneof:GROUP` | adds the flag to the group `GROUP`, which needs one flag at least |

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

### Make a field a positional argument

The tag option `arg` makes a field a positional argument and not a flag.
The arguments take the order of the fields. `optional` makes an argument
optional. A slice field is variadic, so it takes each argument that is
left, and `elem` sets the type of each value. `From` and `ox.App` read these
options. `ox.FlagsFrom` leaves the arguments out.

```go
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

### Run a func before and after a command

`ox.Before(f)` runs `f` before the exec func of the command and of each sub
command below it. `ox.After(f)` runs `f` after it. Each takes a func with a
signature that `ox.Exec` takes. The before funcs run from the root down, and
the after funcs run from the sub command up. These rules apply:

1. If a before func returns an error, the run stops. No other func runs.
2. The after funcs run also when the exec func returns an error. Each after
   func runs, whatever an earlier one returns.
3. The run returns the errors of the exec func and of the after funcs,
   joined with `errors.Join`.
4. `--help`, `--version`, the `help` command, the `version` command and
   completion run no before func and no after func.

With `ox.App`, a method `Before` or `After` of the struct sets the func.

```go
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

### Declare typed positional arguments

`ox.Arg(name, usage, opts...)` adds a positional argument to a command. It
takes the options of a flag, such as a type, `ox.Default`, `ox.Valid`,
`ox.Bind`, `ox.Spec` and `ox.CompFunc`. These rules apply:

1. The run converts each argument to its type and checks it, as for a flag.
   An `ox.FileT` must be a file that exists.
2. The exec func reads the value by the name of the argument, with `otx`.
3. An argument is required, unless it has a default or `ox.Optional()`. A
   required argument cannot follow an optional one.
4. A missing argument gives `ox.ErrMissingArgument`, and one too many gives
   `ox.ErrInvalidArgCount`.
5. Help writes each argument in the usage line and in the `Arguments:`
   section. Completion completes a file, a directory, or the values of
   `ox.Valid`, and else calls the `ox.CompFunc` of the command.

```go
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
```

### A variadic argument

The last argument can be variadic. Give it the type `ox.SliceT`, and the
type of each value with `ox.Elem`. It takes each argument that is left, and
`otx.Slice` reads it. `ox.Spec` sets the name that help writes.

```go
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

### Keep a flag to its command

A sub command inherits each flag of its parents, so `app show --config
a.yaml` sets the `--config` flag of `app`. `ox.Local()` keeps a flag to its
command. The command line of a sub command refuses a local flag with
`ox.ErrUnknownFlag`, and completion does not offer it there. The tag option
is `local`.

```go
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
```

### Show the inherited flags in the help

The help of a sub command shows only its own flags. `ox.InheritedFlags(true)`
adds a `Global Flags:` section with the flags that it inherits, as cobra
does. Give it to `ox.Defaults` or to `ox.Help` on the root, and each sub
command gets it. The section leaves out a local flag, a hidden flag, and a
flag of a parent that a flag of the sub command hides.

```go
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
```

### Set the width of the help

Help wraps each description at `ox.DefaultWrapWidth`, 95 columns by
default. Each line ends inside the width, unless a word is wider. `ox` does
not read the width of the terminal. If a program wants that width, it reads
the width, for example with `golang.org/x/term`, and sets
`ox.DefaultWrapWidth` before it calls `ox.Run`.

```go
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
```

### Stop at the first argument

A wrapper command passes its args to another command. `ox.StopAtArg()` ends
the flags of the command at its first positional argument. Each later
argument is a positional argument, even one that starts with `-`. A sub
command name comes before the first positional argument, so `ox` still
finds it. A sub command does not inherit `ox.StopAtArg()`.

```go
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
```

### Take a prefix of a long flag

`ox.Abbrev()` lets the command line give a long flag by a unique prefix of
its name, as GNU `getopt_long` does. An exact name always wins. A prefix
that matches more than one flag fails with `ox.ErrAmbiguousFlag`, and the
error names the flags. A prefix does not match a hidden flag, a deprecated
flag or a deprecated alias. A sub command inherits `ox.Abbrev()`.

```go
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
```

### Give a bool flag a negated form

`ox.NegateBools()` gives each bool flag of the command tree a negated form,
such as `--no-color` for `--color`. The negated form sets the flag to false,
and takes no value. Each prefix adds a form, and without a prefix the prefix
is `no`. Help writes the flag as `--[no-]color`, with the first prefix. A
flag whose name starts with a prefix and a hyphen, such as `--no-rc`, gets
no second form. `ox.Negatable(false)` keeps one flag out, and
`ox.Negatable(true)` gives one flag the form without `ox.NegateBools()`. A
count flag, a hook flag, an alias and a short name get no negated form. If a
negated form has the name of another flag, the build of the command fails.
`ox.NegateBools()` is not one of the `ox.Defaults()`.

```go
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
```

### Take a plus form of a short flag

`ox.AllowPlus()` lets the short name of a flag take a `+` as well as a `-`,
as bash takes `+o name` to turn an option off and `+ex` to turn two letters
off. The plus form of a plain bool flag sets it to false. The plus form of
any other flag takes its value as the `-` form does, and its hook reads the
sign with `ctx.IsPlus()`. The hooks run in the order of the command line, so
a program that needs the order, as a shell does, gives each option a hook.
A word that starts with `+` is read as flags only when its first letter names
a flag with `ox.AllowPlus()`, so a word such as `+%Y` stays an argument. A
lone `+`, and each word after `--` or after the first argument with
`ox.StopAtArg()`, is an argument. The struct tag option is `plus` (D52).

```go
func ExampleAllowPlus() {
	options := map[string]bool{}
	var order []string
	set := func(ctx *ox.Context, name string) {
		options[name] = !ctx.IsPlus()
		order = append(order, fmt.Sprintf("%s=%t", name, options[name]))
	}
	letter := func(name string) func(*ox.Context) {
		return func(ctx *ox.Context) {
			set(ctx, name)
		}
	}
	run := func(args ...string) {
		order = nil
		ox.Run(
			ox.Usage("sh", "a shell"),
			ox.Help(),
			ox.StopAtArg(),
			ox.Flags().
				Var("set-option", "set an option", ox.HookT, ox.Default(set), ox.Short("o"), ox.AllowPlus(), ox.Spec("NAME")).
				Hook("errexit", "exit on an error", letter("errexit"), ox.Short("e"), ox.AllowPlus()).
				Hook("xtrace", "trace each command", letter("xtrace"), ox.Short("x"), ox.AllowPlus()),
			ox.Exec(func(_ context.Context, args []string) {
				fmt.Println(order, args)
			}),
			ox.Args(args...),
		)
	}
	run("-o", "errexit", "-x", "+o", "errexit", "script.sh", "+x")
	run("-ex", "+ex", "+%Y")
	run("--help")
	// Output:
	// [errexit=true xtrace=true errexit=false] [script.sh +x]
	// [errexit=true xtrace=true errexit=false xtrace=false] [+%Y]
	// sh a shell
	//
	// Usage:
	//   sh [flags] [args]
	//
	// Flags:
	//   -o, +o, --set-option NAME  set an option
	//   -e, +e, --errexit          exit on an error
	//   -x, +x, --xtrace           trace each command
	//   -h,     --help             show help, then exit
}
```

### Write a man page or Markdown docs

Import `github.com/xo/ox/metadoc`, and add the root option
`metadoc.Command(metadoc.ManPage(), metadoc.Markdown())`. It adds the hidden
command `__docs`, which help and completion do not show. D48 decides it.
Then the program writes its docs:

```sh
app __docs man > app.1                # the whole tree as one man page
app __docs markdown > app.md          # the whole tree as one Markdown document
app __docs man --dir man/man1         # app.1, app-sub.1, and so on
app __docs markdown --dir docs/cli    # app.md, app-sub.md, and so on
```

The date of a man page is the option `metadoc.Date`, else
`$SOURCE_DATE_EPOCH`, else today. A program that does not want the writer in
its binary calls `metadoc.Write`, `metadoc.WriteCommand` or
`metadoc.WriteDir` from a test or from `go generate`.
[`metadoc/example_test.go`](../metadoc/example_test.go) holds the examples.

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

### A size as Kubernetes writes it

The verbs `%q` and `%Q` write a size or a rate in the exact form of a
Kubernetes quantity. Each writes an integer with the largest suffix that
divides the size, so no byte is lost. `%q` uses a binary suffix, such as
`Mi` or `Gi`, and `%Q` uses a decimal suffix, such as `M` or `G`. A space
flag, as in `% q`, puts a space before the suffix. `%e` rounds to the best
fitting suffix instead. `ox.ParseSize` reads each form back.

```go
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

### A port range

`ox.PortRange` is a range of ports, from `Lo` to `Hi`. A flag reads one
port, such as `80`, or a range, such as `100-200` or `100:200`. Each port is
1 to 65535, and the first port is not greater than the last. Help and
`String` write the range with a `-`, or one port when `Lo` and `Hi` are the
same. `Contains` tells whether the range holds a port, and `Ports` walks each
port of the range. A slice of port ranges splits its value at each comma. A
field of type `ox.PortRange` gets the type with `ox.From` and `ox.App`.

```go
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
```

### A type that nothing registers

A type that reads its own text needs no registration. Its value, or a
pointer to it, has `UnmarshalText`, `UnmarshalBinary` or `Set(string)
error`, as a `flag.Value` has. `ox.TypeFor[T]()` gives a flag the type `T`,
and `ox.From` and `ox.App` do the same for a field. Help writes the name of
the type in lower case, such as `level` for `slog.Level` and `color` for
`colors.Color`. `ox` has no type for a semantic version, such as `1.2.3`.
A program uses the version type of its semver package, because that type
reads its own text.

```go
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

### A time in a strftime layout

`ox.Strftime` reads and writes a time in a strftime layout, such as
`%d/%m/%Y`. It converts the layout to a Go layout. A flag of type `Date`,
`DateTime`, `Time` or `Timestamp` keeps its type, and any other flag gets the
type `timestamp`. `Strftime` takes only the directives that Go writes and
reads in the same way, such as `%Y`, `%m`, `%d`, `%H`, `%M`, `%S`, `%b`,
`%p`, `%z`, `%T` and `%F`. The doc comment of `Strftime` lists each one.
`Strftime` refuses any other directive. It also refuses text that Go reads
as a part of a time, such as `1` or `Jan`, because a Go layout cannot escape
text. A slice, an array or a map flag refuses it.

```go
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

### Write a help topic for a value of the help flag

`ox.HelpValue(value, usage, funcs...)` takes funcs that run when the command
line gives the flag that value. Here `--help=commands` writes a list of
commands, as psql does, and `--help` writes the help of the command. After
the funcs, the run stops as help does. If a func returns an error, the run
stops with that error. A value without funcs, and a value that no
`HelpValue` names, write the help of the command. To refuse a topic that the
flag does not know, as psql does, give the flag `ox.Valid` with each topic
and `""` for a bare `--help`, such as `ox.Valid("", "commands")`. Then
`--help=bogus` fails with `ox.ErrInvalidValue` (D51).

```go
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

The root package reads JSON. A blank import of `github.com/xo/ox/yaml`,
`github.com/xo/ox/toml`, `github.com/xo/ox/hcl` or `github.com/xo/ox/ini`
adds YAML, TOML, HCL or INI. The extension of a file gives its format. Each
of these packages has two examples of its own, `Example_build` and
`Example_reflect`, which read the same settings.

The key of a flag is the path of its command below the root, then the name
of the flag, joined with dots (D43). So `--steps` of `app migrate` reads
`migrate.steps`. `ox.Key` gives a flag a whole key of its own.

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

### Read the key of a sub command

A flag of a sub command reads its key under the path of the command. A field
of a nested struct reads the path of the field. So the field `Source.Dir` of
`app db migrate` reads `db.migrate.source.dir`, and its flag is
`--source-dir`. A flag that a sub command inherits, such as `--verbose`
here, keeps the key of the command that owns it.

```go
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

### Read an INI file

Import `github.com/xo/ox/ini` for its side effect. It uses only the standard
library, so it is in the module of `ox`. A section nests its keys under its
name, and a dot in the name of a section nests it again. So the key of `dir`
in `[migrate.source]` is `migrate.source.dir`. A key that is written more
than once is a list. [`ini/example_test.go`](../ini/example_test.go) holds
the examples.

## Require flags

D37 proposes these rules. A flag counts as set when it gets a value from the
command line, the environment or a config file. A default does not count.
`Context.Validate` checks the flags after the parse, and before the exec
func runs. `ox.RunError` and `ox.Run` call it.

The flags are not checked for a command without an exec func, because `ox`
writes its help. They are not checked for the help, version and completion
commands either. `--help` and `--version` work without a required flag.

### Make a flag required

`ox.Required` makes a flag required. A flag without a value is an error that
wraps `ox.ErrRequiredFlag`. Help writes `(required)` in place of the
default.

```go
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
```

### Add rules between flags

`ox.OnlyOneOf`, `ox.AllOrNoneOf` and `ox.AtLeastOneOf` take the names of
flags of the command or of its parents. Each rule has its own error:
`ox.ErrOnlyOneOf`, `ox.ErrAllOrNoneOf` and `ox.ErrAtLeastOneOf`. The error
names the flags. A rule of a parent applies to each of its sub commands.

```go
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
```

### Require flags with struct tags

The tag option `required` makes a flag required. The tag options
`onlyoneof:GROUP`, `allornoneof:GROUP` and `atleastoneof:GROUP` add the flag
to a group. The flags with the same option and the same group name are one
group. A flag can be in more than one group.

```go
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
