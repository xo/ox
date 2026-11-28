package metadoc_test

import (
	"os"
	"time"

	"github.com/xo/ox"
	"github.com/xo/ox/metadoc"
)

// Example adds the hidden command `__docs` to a program. Then
// `app __docs man` writes the man page of the program, and
// `app __docs markdown` writes its Markdown docs. With `--dir DIR`, each
// writes a file for each command in the folder. The date of the page is
// fixed, so that the output does not change.
func Example() {
	date := time.Date(2026, 11, 28, 0, 0, 0, 0, time.UTC)
	ox.Run(
		ox.Usage("app", "copy files"),
		ox.VersionString("v1.0.0"),
		ox.Flags().
			String("config", "the config file", ox.Short("c"), ox.Spec("FILE")).
			Int("retries", "the number of retries", ox.Default(3)),
		ox.Arg("path", "the file to copy"),
		ox.Exec(func() {}),
		metadoc.Command(metadoc.ManPage(metadoc.Date(date)), metadoc.Markdown()),
		// a program leaves out ox.Args
		ox.Args("__docs", "man"),
	)
	// Output:
	// .TH "APP" "1" "2026-11-28" "app v1.0.0"
	// .SH "NAME"
	// app \- copy files
	// .SH "SYNOPSIS"
	// .PP
	// \fBapp\fR [flags] path
	// .SH "DESCRIPTION"
	// .PP
	// copy files
	// .SH "OPTIONS"
	// .TP
	// \fB\-c\fR, \fB\-\-config\fR FILE
	// the config file
	// .TP
	// \fB\-\-retries\fR int
	// the number of retries (default: 3)
	// .SH "ARGUMENTS"
	// .TP
	// \fIpath\fR
	// the file to copy
}

// ExampleWrite writes the Markdown docs of a command tree with the library
// API, as a test or `go generate` does, so that the program does not link
// the writer.
func ExampleWrite() {
	root, err := ox.NewCommand(
		ox.Usage("app", "copy files"),
		ox.Flags().
			Bool("verbose", "write more", ox.Short("v")),
		ox.Sub(
			ox.Usage("get", "get a file"),
			ox.Arg("path", "the file to get"),
			ox.Exec(func() {}),
		),
	)
	if err != nil {
		panic(err)
	}
	if err := metadoc.Write(os.Stdout, root, metadoc.Markdown()); err != nil {
		panic(err)
	}
	// Output:
	// # app
	//
	// copy files
	//
	// ## Synopsis
	//
	// ```
	// app [flags] [command] [args]
	// ```
	//
	// ## Options
	//
	// - `-v, --verbose`: write more
	//
	// ## Commands
	//
	// - [app get](#app-get): get a file
	//
	// ### app get
	//
	// get a file
	//
	// #### Synopsis
	//
	// ```
	// app get path
	// ```
	//
	// #### Arguments
	//
	// - `path`: the file to get
}
