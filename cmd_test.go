package kobra

import (
	"bytes"
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	root := testCommand(t)
	for i, test := range parseTests() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := WithRoot(context.Background(), root)
			cmd, args, vars, err := Parse(ctx, root, test.v[1:])
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case cmd == nil:
				t.Fatal("cmd is nil")
			}
			t.Logf("cmd: %s", cmd.Descs[0].Name)
			t.Logf("args: %q", args)
			t.Logf("vars: %v", args)
			var stdout, stderr bytes.Buffer
			ctx = WithStdout(ctx, &stdout)
			ctx = WithStderr(ctx, &stderr)
			ctx = WithCmd(ctx, cmd)
			ctx = WithVars(ctx, vars)
			if err := cmd.Exec(ctx, args); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := stdout.String(); s != test.cmds {
				t.Errorf("\nexpected output:\n%s\ngot:\n%s", test.cmds, s)
			}
		})
	}
}

type parseTest struct {
	v    []string
	cmds string
	args string
	vars string
}

func parseTests() []parseTest {
	return []parseTest{
		{ss(""), "cmd", "", ""},
		{ss("a//"), "cmd", "", ""},
		{ss("a//--foo=a//one////--foo//b//two//--foo//c//three//four//blah//yay"), "cmd//one//two//three//four", "", ""},
		{ss("a//--m=A=100//--m//FOO=200//one//two"), "cmd//one//two", "", ""},
		{ss("a//--m=A=100//--m//FOO=200//one//two"), "cmd//one//two", "", ""},
		{ss("a//--b//one//--b=false//two//--b=t//three//--b=1"), "cmd//one/two/three", "", ""},
		{ss("a//--inc//one//--inc//--inc//two//--inc//three"), "cmd//one//two//three", "", ""},
		{ss("a//-fa=b//-iiib//one//two//-bbb//three"), "cmd//one//two//three", "", ""},
	}
}

func ss(s string) []string {
	return strings.Split(s, "//")
}

func testCommand(t *testing.T) *Command {
	t.Helper()
	cmd, err := NewCommand(
		testDump(t, "cmd"),
		Usage("cmd", ""),
		Flags().
			String("avar", "", Short("a")).
			Bool("bvar", "", Short("b")).
			Count("inc", "", Short("i")).
			Map("map", "", RunesT, Short("m")).
			Slice("foo", "", Short("f")),
		Sub(
			testDump(t, "one"),
			Usage("one", ""),
			Sub(
				testDump(t, "two"),
				Usage("two", ""),
				Sub(
					testDump(t, "three"),
					Usage("three", ""),
				),
			),
			Sub(
				testDump(t, "four"),
				Usage("four", ""),
			),
		),
		Sub(
			testDump(t, "five"),
			Usage("five", ""),
		),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return cmd
}

func testDump(t *testing.T, name string) func(ctx context.Context, args []string) error {
	t.Helper()
	return func(ctx context.Context, args []string) error {
		_, _ = fmt.Fprintln(Stdout(ctx), "exec:", name)
		_, _ = fmt.Fprintln(Stdout(ctx), "root:", RootName(ctx))
		cmd := Cmd(ctx)
		_, _ = fmt.Fprintln(Stdout(ctx), "name:", cmd.Descs[0].Name)
		var v []string
		for c := cmd; c != nil; {
			v, c = append(v, c.Descs[0].Name), c.Parent
		}
		slices.Reverse(v)
		_, _ = fmt.Fprintln(Stdout(ctx), "tree:", strings.Join(v, "//"))
		_, _ = fmt.Fprintln(Stdout(ctx), "args:", strings.Join(args, "//"))
		vars, _ := VarsOK(ctx)
		_, _ = fmt.Fprint(Stdout(ctx), "[")
		var i int
		for k, val := range vars {
			if i != 0 {
				_, _ = fmt.Fprint(Stdout(ctx), " ")
			}
			var s string
			if v, ok := val.Var.(interface{ String() string }); ok {
				s = v.String()
			} else {
				s, _ = val.Var.Get()
			}
			_, _ = fmt.Fprint(Stdout(ctx), k+":"+s)
			i++
		}
		_, _ = fmt.Fprintln(Stdout(ctx), "]")
		return nil
	}
}
