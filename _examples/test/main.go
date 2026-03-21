package main

import (
	"fmt"
	"os"

	"github.com/xo/ox"
)

func main() {
	args := new(Args)
	ox.Run(
		ox.Defaults(),
		ox.Usage("test", "the test app"),
		ox.Exec(run(args)),
		ox.From(args),
	)
}

func run(args *Args) func() error {
	return func() error {
		fmt.Fprintf(os.Stderr, "Arg0: %v\n", args.Arg0)
		fmt.Fprintf(os.Stderr, "SubArgs: %#v\n", args.SubArgs)
		fmt.Fprintf(os.Stderr, "SubB: %#v\n", args.SubB)
		return nil
	}
}

type Args struct {
	Arg0    string   `ox:"arg zero!,short:0"`
	SubArgs SubArgs  `ox:"blah"`
	SubB    *SubArgs `ox:"yes"`
}

type SubArgs struct {
	Arg1  int            `ox:"a sub a,short:a"`
	MyMap map[string]int `ox:"a map"`
	Arg2  []int64        `ox:"the sub 2 arg,short:b"`
}
