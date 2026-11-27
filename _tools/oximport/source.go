package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Error is an error of the importer.
type Error string

// Error satisfies the [error] interface.
func (err Error) Error() string {
	return string(err)
}

// Errors.
const (
	// ErrNotRecorded is the error when an archive does not hold the output of
	// a command line.
	ErrNotRecorded Error = "not recorded"
	// ErrUnknownFormat is the error when the importer cannot tell the format
	// of the help.
	ErrUnknownFormat Error = "unknown help format"
	// ErrInvalidArchive is the error when an archive cannot be read.
	ErrInvalidArchive Error = "invalid archive"
)

// Runner runs the program with arguments, and returns what it wrote.
type Runner interface {
	// Run runs the program.
	Run(ctx context.Context, args ...string) ([]byte, error)
	// Home returns the home directory that the program sees.
	Home() string
}

// execRunner runs a program.
type execRunner struct {
	// path is the path of the program.
	path string
	// timeout is the longest time that one run can take.
	timeout time.Duration
}

// Run satisfies the [Runner] interface. It returns what the program wrote to
// standard output, or to standard error when it wrote nothing to standard
// output, as the flag package of Go does. A program that exits with an error
// after it wrote help is not an error. The program gets no standard input,
// and runs with LC_ALL=C, so that it writes its help in English. The rest of
// the environment stays, because some defaults come from it, as NO_COLOR
// changes the default of --color in helm.
func (r *execRunner) Run(ctx context.Context, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, r.path, args...)
	// GNU programs write argv[0] in their usage
	cmd.Args[0] = filepath.Base(r.path)
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	cmd.Env = append(os.Environ(), "LC_ALL=C")
	err := cmd.Run()
	out := stdout.Bytes()
	if len(out) == 0 {
		out = stderr.Bytes()
	}
	var exitErr *exec.ExitError
	switch {
	case err == nil:
	case errors.As(err, &exitErr) && len(out) != 0 && ctx.Err() == nil:
	default:
		return nil, fmt.Errorf("running %s %s: %w", r.path, argsKey(args), err)
	}
	return out, nil
}

// Home satisfies the [Runner] interface.
func (r *execRunner) Home() string {
	home, _ := os.UserHomeDir()
	return home
}

// recordRunner runs a program, and keeps what it wrote in an archive.
type recordRunner struct {
	// runner runs the program.
	runner Runner
	// archive holds the output.
	archive *archive
}

// Run satisfies the [Runner] interface. It keeps the output of __complete
// only when it is from cobra, so that the archive does not keep the error that
// another program writes for it.
func (r *recordRunner) Run(ctx context.Context, args ...string) ([]byte, error) {
	out, err := r.runner.Run(ctx, args...)
	if err != nil {
		return nil, err
	}
	keep := out
	if len(args) != 0 && args[0] == "__complete" {
		if _, ok := parseComplete(out); !ok {
			keep = nil
		}
	}
	r.archive.set(argsKey(args), scrub(keep, r.runner.Home()))
	return out, nil
}

// Home satisfies the [Runner] interface.
func (r *recordRunner) Home() string {
	return r.runner.Home()
}

// scrub replaces the home directory and the host name in out, so that an
// archive does not hold the name of the user or of the host.
func scrub(out []byte, home string) []byte {
	if home != "" && home != "/" {
		out = bytes.ReplaceAll(out, []byte(home), []byte(archiveHome))
	}
	if host, err := os.Hostname(); err == nil && len(host) > 3 {
		out = bytes.ReplaceAll(out, []byte(host), []byte("localhost"))
	}
	return out
}

// archiveHome is the home directory in an archive.
const archiveHome = "/home/user"

// archive holds the output of each command line of a program. Its file is a
// txtar file: a comment, then each output after a line `-- <args> --`.
type archive struct {
	// comment is the text before the first output.
	comment string
	// keys are the command lines, in the order of the file.
	keys []string
	// outputs are the outputs, by command line.
	outputs map[string][]byte
}

// newArchive creates an archive.
func newArchive(comment string) *archive {
	return &archive{
		comment: comment,
		outputs: make(map[string][]byte),
	}
}

// readArchive reads an archive file.
func readArchive(name string) (*archive, error) {
	buf, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	a := newArchive("")
	key, inFile := "", false
	var comment, out []string
	flush := func() {
		if inFile {
			s := strings.Join(out, "\n")
			if len(out) != 0 {
				s += "\n"
			}
			a.set(key, []byte(s))
		}
	}
	for line := range strings.SplitSeq(strings.TrimSuffix(string(buf), "\n"), "\n") {
		if k, ok := archiveName(line); ok {
			flush()
			key, inFile, out = k, true, nil
			continue
		}
		if inFile {
			out = append(out, line)
		} else {
			comment = append(comment, line)
		}
	}
	flush()
	if len(a.keys) == 0 {
		return nil, fmt.Errorf("%w: %s: no output", ErrInvalidArchive, name)
	}
	a.comment = strings.Join(comment, "\n")
	return a, nil
}

// archiveName returns the command line of a line `-- <args> --`.
func archiveName(line string) (string, bool) {
	if !strings.HasPrefix(line, "-- ") || !strings.HasSuffix(line, " --") || len(line) < 6 {
		return "", false
	}
	return strings.TrimSpace(line[3 : len(line)-3]), true
}

// set sets the output of a command line.
func (a *archive) set(key string, out []byte) {
	if _, ok := a.outputs[key]; !ok {
		a.keys = append(a.keys, key)
	}
	a.outputs[key] = out
}

// write writes the archive to a file.
func (a *archive) write(name string) error {
	var buf bytes.Buffer
	if a.comment != "" {
		_, _ = buf.WriteString(strings.TrimRight(a.comment, "\n") + "\n")
	}
	for _, key := range a.keys {
		_, _ = fmt.Fprintf(&buf, "-- %s --\n", key)
		out := a.outputs[key]
		_, _ = buf.Write(out)
		if len(out) != 0 && out[len(out)-1] != '\n' {
			_ = buf.WriteByte('\n')
		}
	}
	return os.WriteFile(name, buf.Bytes(), 0o644)
}

// replayRunner returns the output that an archive holds.
type replayRunner struct {
	// archive holds the output.
	archive *archive
}

// Run satisfies the [Runner] interface.
func (r *replayRunner) Run(_ context.Context, args ...string) ([]byte, error) {
	key := argsKey(args)
	out, ok := r.archive.outputs[key]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrNotRecorded, key)
	}
	return out, nil
}

// Home satisfies the [Runner] interface.
func (r *replayRunner) Home() string {
	return archiveHome
}

// argsKey returns the command line of args, with each empty argument, and each
// argument that has a space or a quote, in double quotes.
func argsKey(args []string) string {
	v := slices.Clone(args)
	for i, s := range v {
		if s == "" || strings.ContainsAny(s, " \t\"'") {
			v[i] = strconv.Quote(s)
		}
	}
	return strings.Join(v, " ")
}
