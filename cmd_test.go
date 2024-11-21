package ox

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	root := testCommand(t)
	for i, test := range parseTests() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := WithRoot(context.Background(), root)
			vars := make(Vars)
			cmd, args, err := Parse(ctx, root, test.v[1:], vars)
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case cmd == nil:
				t.Fatal("cmd is nil")
			}
			t.Logf("cmd: %s", cmd.Name())
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
			s := strings.TrimSpace(stdout.String())
			if exp := strings.Join(test.exp, "\n"); s != exp {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", exp, s)
			}
		})
	}
}

type parseTest struct {
	v   []string
	exp []string
}

func parseTests() []parseTest {
	return []parseTest{
		{
			ss(""),
			[]string{
				"exec: cmd",
				"root: cmd",
				"name: cmd",
				"tree: [cmd]",
				"args: []",
				"vars: []",
			},
		},
		{
			ss("a//"),
			[]string{
				"exec: cmd",
				"root: cmd",
				"name: cmd",
				"tree: [cmd]",
				"args: []",
				"vars: []",
			},
		},
		{
			ss("a//--foo=a//one////--foo//b//two//--foo//c//three//four//blah//yay"),
			[]string{
				"exec: one",
				"root: cmd",
				"name: one",
				"tree: [cmd one]",
				"args: [ two three four blah yay]",
				"vars: [foo:[a b c]]",
			},
		},
		{
			ss("a//--m=A=100//--m//FOO=200//one//two"),
			[]string{
				"exec: two",
				"root: cmd",
				"name: two",
				"tree: [cmd one two]",
				"args: []",
				"vars: [map:[A:100 FOO:200]]",
			},
		},
		{
			ss("a//-b//one//-b=false//two//-b=t//three//--b=1"),
			[]string{
				"exec: three",
				"root: cmd",
				"name: three",
				"tree: [cmd one two three]",
				"args: []",
				"vars: [bvar:true]",
			},
		},
		{
			ss("a//--inc//one//--inc//--inc//two//--inc//three"),
			[]string{
				"exec: three",
				"root: cmd",
				"name: three",
				"tree: [cmd one two three]",
				"args: []",
				"vars: [inc:4]",
			},
		},
		{
			ss("a//-fa=b//-iiib//one//two//-bbb//three"),
			[]string{
				"exec: three",
				"root: cmd",
				"name: three",
				"tree: [cmd one two three]",
				"args: []",
				"vars: [bvar:true foo:[a=b] inc:3]",
			},
		},
		{
			ss("a//one//two//four//-b"),
			[]string{
				"exec: two",
				"root: cmd",
				"name: two",
				"tree: [cmd one two]",
				"args: [four]",
				"vars: [bvar:true]",
			},
		},
		{
			ss("a//one//four//-iii//--inc//fun"),
			[]string{
				"exec: four",
				"root: cmd",
				"name: four",
				"tree: [cmd one four]",
				"args: [fun]",
				"vars: [inc:4]",
			},
		},
		{
			ss("a//five//-u//file:a//-ufile:b//-c=1.2.3.4/24//--cidr//2.4.6.8/0//foo//bar"),
			[]string{
				"exec: five",
				"root: cmd",
				"name: five",
				"tree: [cmd five]",
				"args: [foo bar]",
				"vars: [cidr:[1.2.3.4/24 2.4.6.8/0] url:[file:a file:b] val:125]",
			},
		},
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
			Flags().
				String("val", "", Short("l"), Default(125)).
				Slice("cidr", "", Short("c"), CIDRT).
				Slice("url", "", Short("u"), URLT),
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
		_, _ = fmt.Fprintln(Stdout(ctx), "name:", cmd.Name())
		_, _ = fmt.Fprintln(Stdout(ctx), "tree:", cmd.Tree())
		_, _ = fmt.Fprintln(Stdout(ctx), "args:", args)
		vars, _ := VarsOK(ctx)
		_, _ = fmt.Fprintln(Stdout(ctx), "vars:", vars)
		return nil
	}
}
