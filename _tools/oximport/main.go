// Command oximport reads the command-line interface of a program and writes
// an ox program that has the same commands and flags. Run it from the
// repository root:
//
//	go run ./_tools/oximport -o _examples/gen psql gh
//
// It reads a program built with cobra through `<prog> __complete` and its
// help, and every other program through its help. It reads the help of GNU
// getopt_long and argp (du, cp, tar), of psql, of cobra (gh, helm, kubectl)
// and of the flag package of Go. It writes three styles of each program, in
// <out>/<prog>/hand, <out>/<prog>/reflect and <out>/<prog>/app. The app style
// uses ox.AppFrom, and it builds after `go generate` writes its doc comments.
//
// With --record, it keeps what the program wrote in <dir>/<prog>.txtar. With
// --replay, it reads that file in place of the program, so that the tests run
// without the program. docs/DESIGN.md has the design, in "The importer".
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/xo/ox"
)

// args are the flags of oximport.
type args struct {
	Out     string   `ox:"directory to write the programs to,short:o,default:gen,spec:DIR"`
	Format  string   `ox:"format of the help: cobra\\, gnu\\, psql or goflag (default: find it),short:f,spec:FORMAT"`
	Style   []string `ox:"style to write: hand\\, reflect or app (default: all),short:s,spec:STYLE"`
	Path    string   `ox:"path of commands to read\\, such as \"pr list\",short:p,spec:PATH"`
	Depth   int      `ox:"levels of commands to read below the path (-1: all),short:d,default:-1"`
	NoValue bool     `ox:"do not ask __complete for the values of each flag,name:no-values"`
	Record  string   `ox:"directory to keep what each program wrote in,short:r,spec:DIR"`
	Replay  string   `ox:"directory to read what each program wrote from,spec:DIR"`
	Timeout string   `ox:"longest time for one run of a program,type:duration,default:10s,spec:DURATION"`
	Verbose bool     `ox:"write what the importer reads,short:v"`
}

func main() {
	args := new(args)
	ox.Run(
		ox.Defaults(),
		ox.Usage("oximport", "write ox programs from the interface of other programs"),
		ox.Spec("[flags] <program>..."),
		ox.From(args),
		ox.Exec(func(ctx context.Context, names []string) error {
			if len(names) == 0 {
				return fmt.Errorf("%w: no program", ox.ErrMissingArgument)
			}
			for _, name := range names {
				if err := args.run(ctx, name); err != nil {
					return err
				}
			}
			return nil
		}),
	)
}

// run imports a program, and writes each style.
func (args *args) run(ctx context.Context, prog string) error {
	name := filepath.Base(prog)
	runner, rec, err := args.runner(ctx, prog, name)
	if err != nil {
		return err
	}
	imp := &Importer{
		Runner: runner,
		Format: Format(args.Format),
		Values: !args.NoValue,
		Path:   strings.Fields(args.Path),
		Depth:  args.Depth,
	}
	if args.Verbose {
		imp.Logf = func(s string, v ...any) {
			_, _ = fmt.Fprintf(os.Stderr, s+"\n", v...)
		}
	}
	root, err := imp.Import(ctx, name)
	if err != nil {
		return err
	}
	if rec != nil {
		if err := os.MkdirAll(args.Record, 0o755); err != nil {
			return err
		}
		if err := rec.write(filepath.Join(args.Record, name+".txtar")); err != nil {
			return err
		}
	}
	styles := Styles
	if len(args.Style) != 0 {
		styles = nil
		for _, s := range args.Style {
			styles = append(styles, Style(s))
		}
	}
	for _, style := range styles {
		src, err := Write(root, style)
		if err != nil {
			return err
		}
		dir := filepath.Join(args.Out, name, string(style))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, "main.go"), src, 0o644); err != nil {
			return err
		}
	}
	return nil
}

// runner returns the runner for a program, and the archive that it records
// to.
func (args *args) runner(ctx context.Context, prog, name string) (Runner, *archive, error) {
	if args.Replay != "" {
		a, err := readArchive(filepath.Join(args.Replay, name+".txtar"))
		if err != nil {
			return nil, nil, err
		}
		return &replayRunner{archive: a}, nil, nil
	}
	path, err := exec.LookPath(prog)
	if err != nil {
		return nil, nil, err
	}
	// a time.Duration field cannot take a default from a tag (B19)
	timeout, err := time.ParseDuration(args.Timeout)
	if err != nil {
		return nil, nil, err
	}
	var runner Runner = &execRunner{path: path, timeout: timeout}
	if args.Record == "" {
		return runner, nil, nil
	}
	comment := fmt.Sprintf("oximport recorded what %s wrote on %s.", name, time.Now().Format(time.DateOnly))
	if out, err := runner.Run(ctx, "--version"); err == nil {
		if line, _, _ := strings.Cut(string(scrub(out, runner.Home())), "\n"); strings.ContainsAny(line, "0123456789") {
			comment += "\nVersion: " + strings.TrimSpace(line)
		}
	}
	comment += "\nThe home directory is " + archiveHome + "."
	a := newArchive(comment)
	return &recordRunner{runner: runner, archive: a}, a, nil
}
