package test

import (
	"context"
	"flag"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/alecthomas/kong"
	goflags "github.com/jessevdk/go-flags"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/urfave/cli/v3"

	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

// BenchmarkCommandLine runs the same command line with ox and with other Go
// flag packages. Each run builds the command, parses the command line and
// calls the func of the command, as a program does each time it starts. The
// command line is:
//
//	app serve -v --host localhost --port 8080 --timeout 30s --tags a --tags b --debug file.txt
//
// Before the timer starts, each package must give the values of want, so that
// a package that parses less cannot look faster. ox runs with ox.Defaults, as
// cobra and kingpin add help and completion by default.
//
// Run it with:
//
//	go test -run '^$' -bench CommandLine -benchmem
func BenchmarkCommandLine(b *testing.B) {
	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			if got, err := bench.run(benchArgs); err != nil {
				b.Fatalf("expected no error, got: %v", err)
			} else if !reflect.DeepEqual(got, want) {
				b.Fatalf("expected %+v, got: %+v", want, got)
			}
			b.ReportAllocs()
			for b.Loop() {
				if _, err := bench.run(benchArgs); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// TestCommandLine holds that each package of BenchmarkCommandLine parses the
// command line, so that a change to a package fails a test and not only a
// benchmark.
func TestCommandLine(t *testing.T) {
	t.Parallel()
	for _, bench := range benchmarks {
		t.Run(bench.name, func(t *testing.T) {
			t.Parallel()
			got, err := bench.run(benchArgs)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("expected %+v, got: %+v", want, got)
			}
		})
	}
}

// result holds the values that a run parsed.
type result struct {
	Verbose bool
	Config  string
	Host    string
	Port    int
	Timeout time.Duration
	Tags    []string
	Debug   bool
	Args    []string
}

// benchArgs is the command line of each run, without the name of the
// program.
var benchArgs = strings.Fields("serve -v --host localhost --port 8080 --timeout 30s --tags a --tags b --debug file.txt")

// want is the result of each run. --config is not on the command line, so
// it has its default.
var want = result{
	Verbose: true,
	Config:  "app.yaml",
	Host:    "localhost",
	Port:    8080,
	Timeout: 30 * time.Second,
	Tags:    []string{"a", "b"},
	Debug:   true,
	Args:    []string{"file.txt"},
}

// benchmarks are the packages of BenchmarkCommandLine.
var benchmarks = []struct {
	name string
	run  func([]string) (result, error)
}{
	{"ox", runOx},
	{"ox-struct", runOxStruct},
	{"flag", runFlag},
	{"pflag", runPflag},
	{"cobra", runCobra},
	{"urfave-cli", runUrfave},
	{"kong", runKong},
	{"kingpin", runKingpin},
	{"go-flags", runGoFlags},
}

// runOx builds the flags by hand, and reads them with otx.
func runOx(args []string) (result, error) {
	var r result
	err := ox.RunError(
		context.Background(),
		ox.Usage("app", "an app"),
		ox.Defaults(),
		ox.Sub(
			ox.Usage("serve", "serve files"),
			ox.Flags().
				Bool("verbose", "write more", ox.Short("v")).
				String("config", "the config file", ox.Default("app.yaml")).
				String("host", "the host").
				Int("port", "the port").
				Duration("timeout", "the timeout").
				Slice("tags", "the tags").
				Bool("debug", "debug"),
			ox.ValidArgs(1, 1),
			ox.Exec(func(ctx context.Context, args []string) {
				r = result{
					Verbose: otx.Bool(ctx, "verbose"),
					Config:  otx.String(ctx, "config"),
					Host:    otx.String(ctx, "host"),
					Port:    otx.Int(ctx, "port"),
					Timeout: otx.Duration(ctx, "timeout"),
					Tags:    otx.Slice[string](ctx, "tags"),
					Debug:   otx.Bool(ctx, "debug"),
					Args:    args,
				}
			}),
		),
		ox.Args(args...),
	)
	return r, err
}

// runOxStruct reads the flags into a struct with ox.From.
func runOxStruct(args []string) (result, error) {
	var r result
	var v struct {
		Verbose bool          `ox:"write more,short:v"`
		Config  string        `ox:"the config file,default:app.yaml"`
		Host    string        `ox:"the host"`
		Port    int           `ox:"the port"`
		Timeout time.Duration `ox:"the timeout"`
		Tags    []string      `ox:"the tags"`
		Debug   bool          `ox:"debug"`
	}
	err := ox.RunError(
		context.Background(),
		ox.Usage("app", "an app"),
		ox.Defaults(),
		ox.Sub(
			ox.Usage("serve", "serve files"),
			ox.From(&v),
			ox.ValidArgs(1, 1),
			ox.Exec(func(args []string) {
				r = result{v.Verbose, v.Config, v.Host, v.Port, v.Timeout, v.Tags, v.Debug, args}
			}),
		),
		ox.Args(args...),
	)
	return r, err
}

// runFlag uses the standard flag package. It has no sub command, so the run
// takes the first arg as the command, as a program with flag does.
func runFlag(args []string) (result, error) {
	var r result
	fs := flag.NewFlagSet("serve", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.BoolVar(&r.Verbose, "verbose", false, "write more")
	fs.BoolVar(&r.Verbose, "v", false, "write more")
	fs.StringVar(&r.Config, "config", "app.yaml", "the config file")
	fs.StringVar(&r.Host, "host", "", "the host")
	fs.IntVar(&r.Port, "port", 0, "the port")
	fs.DurationVar(&r.Timeout, "timeout", 0, "the timeout")
	fs.Func("tags", "the tags", func(s string) error {
		r.Tags = append(r.Tags, s)
		return nil
	})
	fs.BoolVar(&r.Debug, "debug", false, "debug")
	if len(args) == 0 || args[0] != "serve" {
		return r, flag.ErrHelp
	}
	if err := fs.Parse(args[1:]); err != nil {
		return r, err
	}
	r.Args = fs.Args()
	return r, nil
}

// runPflag uses pflag, which has no sub command either.
func runPflag(args []string) (result, error) {
	var r result
	fs := pflag.NewFlagSet("serve", pflag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.BoolVarP(&r.Verbose, "verbose", "v", false, "write more")
	fs.StringVar(&r.Config, "config", "app.yaml", "the config file")
	fs.StringVar(&r.Host, "host", "", "the host")
	fs.IntVar(&r.Port, "port", 0, "the port")
	fs.DurationVar(&r.Timeout, "timeout", 0, "the timeout")
	fs.StringSliceVar(&r.Tags, "tags", nil, "the tags")
	fs.BoolVar(&r.Debug, "debug", false, "debug")
	if len(args) == 0 || args[0] != "serve" {
		return r, pflag.ErrHelp
	}
	if err := fs.Parse(args[1:]); err != nil {
		return r, err
	}
	r.Args = fs.Args()
	return r, nil
}

// runCobra uses cobra, with pflag.
func runCobra(args []string) (result, error) {
	var r result
	root := &cobra.Command{Use: "app", Short: "an app", SilenceUsage: true, SilenceErrors: true}
	serve := &cobra.Command{
		Use:   "serve",
		Short: "serve files",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			r.Args = args
		},
	}
	fs := serve.Flags()
	fs.BoolVarP(&r.Verbose, "verbose", "v", false, "write more")
	fs.StringVar(&r.Config, "config", "app.yaml", "the config file")
	fs.StringVar(&r.Host, "host", "", "the host")
	fs.IntVar(&r.Port, "port", 0, "the port")
	fs.DurationVar(&r.Timeout, "timeout", 0, "the timeout")
	fs.StringSliceVar(&r.Tags, "tags", nil, "the tags")
	fs.BoolVar(&r.Debug, "debug", false, "debug")
	root.AddCommand(serve)
	root.SetArgs(args)
	root.SetOut(io.Discard)
	root.SetErr(io.Discard)
	err := root.Execute()
	return r, err
}

// runUrfave uses urfave/cli v3.
func runUrfave(args []string) (result, error) {
	var r result
	cmd := &cli.Command{
		Name:      "app",
		Usage:     "an app",
		Writer:    io.Discard,
		ErrWriter: io.Discard,
		Commands: []*cli.Command{{
			Name:  "serve",
			Usage: "serve files",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "write more"},
				&cli.StringFlag{Name: "config", Value: "app.yaml", Usage: "the config file"},
				&cli.StringFlag{Name: "host", Usage: "the host"},
				&cli.IntFlag{Name: "port", Usage: "the port"},
				&cli.DurationFlag{Name: "timeout", Usage: "the timeout"},
				&cli.StringSliceFlag{Name: "tags", Usage: "the tags"},
				&cli.BoolFlag{Name: "debug", Usage: "debug"},
			},
			Action: func(_ context.Context, c *cli.Command) error {
				r = result{
					Verbose: c.Bool("verbose"),
					Config:  c.String("config"),
					Host:    c.String("host"),
					Port:    c.Int("port"),
					Timeout: c.Duration("timeout"),
					Tags:    c.StringSlice("tags"),
					Debug:   c.Bool("debug"),
					Args:    c.Args().Slice(),
				}
				return nil
			},
		}},
	}
	err := cmd.Run(context.Background(), append([]string{"app"}, args...))
	return r, err
}

// kongCLI is the command line of runKong.
type kongCLI struct {
	Serve struct {
		Verbose bool          `short:"v" help:"write more"`
		Config  string        `default:"app.yaml" help:"the config file"`
		Host    string        `help:"the host"`
		Port    int           `help:"the port"`
		Timeout time.Duration `help:"the timeout"`
		Tags    []string      `help:"the tags"`
		Debug   bool          `help:"debug"`
		File    string        `arg:"" help:"the file"`
	} `cmd:"" help:"serve files"`
}

// runKong uses kong.
func runKong(args []string) (result, error) {
	var c kongCLI
	parser, err := kong.New(&c, kong.Name("app"), kong.Writers(io.Discard, io.Discard))
	if err != nil {
		return result{}, err
	}
	if _, err := parser.Parse(args); err != nil {
		return result{}, err
	}
	s := c.Serve
	return result{s.Verbose, s.Config, s.Host, s.Port, s.Timeout, s.Tags, s.Debug, []string{s.File}}, nil
}

// runKingpin uses kingpin v2.
func runKingpin(args []string) (result, error) {
	app := kingpin.New("app", "an app").Terminate(nil)
	app.ErrorWriter(io.Discard).UsageWriter(io.Discard)
	serve := app.Command("serve", "serve files")
	verbose := serve.Flag("verbose", "write more").Short('v').Bool()
	config := serve.Flag("config", "the config file").Default("app.yaml").String()
	host := serve.Flag("host", "the host").String()
	port := serve.Flag("port", "the port").Int()
	timeout := serve.Flag("timeout", "the timeout").Duration()
	tags := serve.Flag("tags", "the tags").Strings()
	debug := serve.Flag("debug", "debug").Bool()
	file := serve.Arg("file", "the file").Required().String()
	if _, err := app.Parse(args); err != nil {
		return result{}, err
	}
	return result{*verbose, *config, *host, *port, *timeout, *tags, *debug, []string{*file}}, nil
}

// goFlagsServe is the serve command of runGoFlags.
type goFlagsServe struct {
	Verbose bool          `short:"v" long:"verbose" description:"write more"`
	Config  string        `long:"config" default:"app.yaml" description:"the config file"`
	Host    string        `long:"host" description:"the host"`
	Port    int           `long:"port" description:"the port"`
	Timeout time.Duration `long:"timeout" description:"the timeout"`
	Tags    []string      `long:"tags" description:"the tags"`
	Debug   bool          `long:"debug" description:"debug"`
	Args    struct {
		File string `positional-arg-name:"file" required:"yes"`
	} `positional-args:"yes"`
}

// runGoFlags uses go-flags.
func runGoFlags(args []string) (result, error) {
	var serve goFlagsServe
	parser := goflags.NewNamedParser("app", goflags.HelpFlag)
	if _, err := parser.AddCommand("serve", "serve files", "", &serve); err != nil {
		return result{}, err
	}
	if _, err := parser.ParseArgs(args); err != nil {
		return result{}, err
	}
	s := serve
	return result{s.Verbose, s.Config, s.Host, s.Port, s.Timeout, s.Tags, s.Debug, []string{s.Args.File}}, nil
}
