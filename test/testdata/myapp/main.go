// Command my-app is the program that TestShellCompletion completes in each
// shell. Its name has a hyphen, because a hyphen is where the name of the
// program and the name of its shell variables differ.
package main

import "github.com/xo/ox"

func main() {
	ox.Run(
		ox.Usage("my-app", "an app"),
		ox.Defaults(),
		ox.Sub(
			ox.Usage("shell", "run a shell"),
			ox.ValidArgs(0, 1, "bash", "zsh", "fish"),
		),
	)
}
