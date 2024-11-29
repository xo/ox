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
	for i, test := range parseTests() {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Logf("invoked: %v", test.v)
			root, vars := testCommand(t), make(Vars)
			cmd, args, err := Parse(root, test.v[1:], vars)
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case cmd == nil:
				t.Fatal("cmd is nil")
			}
			t.Logf("cmd: %s", cmd.Name)
			t.Logf("args: %q", args)
			t.Logf("vars: %s", vars)
			buf := new(bytes.Buffer)
			ctx := WithContext(context.Background(), &Context{
				Args:   test.v,
				Root:   root,
				Exec:   cmd,
				Stdout: buf,
				Vars:   vars,
			})
			if err := cmd.Exec(ctx, args); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			s := strings.TrimSpace(buf.String())
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
				"name: cmd",
				"path: []",
				"args: []",
				"vars: [int:15]",
			},
		},
		{
			ss("a//"),
			[]string{
				"exec: cmd",
				"name: cmd",
				"path: []",
				"args: []",
				"vars: [int:15]",
			},
		},
		{
			ss("a//--foo=a//one////--foo//b//two//--foo//c//three//four//blah//yay"),
			[]string{
				"exec: one",
				"name: one",
				"path: [one]",
				"args: [ two three four blah yay]",
				"vars: [foo:[a b c] int:15]",
			},
		},
		{
			ss("a//-m=A=100//-m//FOO=200//one//two"),
			[]string{
				"exec: two",
				"name: two",
				"path: [one two]",
				"args: []",
				"vars: [int:15 map:[A:100 FOO:200]]",
			},
		},
		{
			ss("a//-b//one//-b=false//two//-b=t//three//-b=1"),
			[]string{
				"exec: three",
				"name: three",
				"path: [one two three]",
				"args: []",
				"vars: [bvar:true int:15]",
			},
		},
		{
			ss("a//--inc//one//--inc//--inc//two//--inc//three"),
			[]string{
				"exec: three",
				"name: three",
				"path: [one two three]",
				"args: []",
				"vars: [inc:4 int:15]",
			},
		},
		{
			ss("a//-fa=b//-iiib//one//two//-bbb//three"),
			[]string{
				"exec: three",
				"name: three",
				"path: [one two three]",
				"args: []",
				"vars: [bvar:true foo:[a=b] inc:3 int:15]",
			},
		},
		{
			ss("a//one//two//four//-b"),
			[]string{
				"exec: two",
				"name: two",
				"path: [one two]",
				"args: [four]",
				"vars: [bvar:true int:15]",
			},
		},
		{
			ss("a//one//four//-iii//--inc//fun"),
			[]string{
				"exec: four",
				"name: four",
				"path: [one four]",
				"args: [fun]",
				"vars: [inc:4 int:15]",
			},
		},
		{
			ss("a//five//-u//file:a//-ufile:b//-c=1.2.3.4/24//--cidr//2.4.6.8/0//foo//bar"),
			[]string{
				"exec: five",
				"name: five",
				"path: [five]",
				"args: [foo bar]",
				"vars: [cidr:[1.2.3.4/24 2.4.6.8/0] int:15 url:[file:a file:b] val:125]",
			},
		},
		{
			ss("a//five//--time//A=07:15:13//a//-d2001-12-25//-d=2002-01-15//--time=B=12:15:32//b"),
			[]string{
				"exec: five",
				"name: five",
				"path: [five]",
				"args: [a b]",
				"vars: [date:[2001-12-25 2002-01-15] int:15 time:[A:07:15:13 B:12:15:32] val:125]",
			},
		},
		{
			ss("a//--//five//--a=b"),
			[]string{
				"exec: cmd",
				"name: cmd",
				"path: []",
				"args: [five --a=b]",
				"vars: [int:15]",
			},
		},
		{
			ss("a//five//-T//255=07:15:32//-T//128=12:15:20//-C=16.1=A//-iiiiC25=J//-C//16.100=C//-C//17.0=1//-C17.0000=//-C//17=//--//--//-a//-b=c"),
			[]string{
				"exec: five",
				"name: five",
				"path: [five]",
				"args: [-- -a -b=c]",
				"vars: [countmap:[16.1:2 17:3 25:1] inc:4 int:15 timemap:[128:12:15:20 255:07:15:32] val:125]",
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
		Exec(testDump(t, "cmd")),
		Usage("cmd", ""),
		Flags().
			String("avar", "", Short("a")).
			Bool("bvar", "", Short("b")).
			Count("inc", "", Short("i")).
			Map("map", "", RunesT, Short("m")).
			Slice("foo", "", Short("f")).
			Int("int", "", Short("n"), Default(float64(15.0))),
		Sub(
			Exec(testDump(t, "one")),
			Usage("one", ""),
			Sub(
				Exec(testDump(t, "two")),
				Usage("two", ""),
				Sub(
					Exec(testDump(t, "three")),
					Usage("three", ""),
				),
			),
			Sub(
				Exec(testDump(t, "four")),
				Usage("four", ""),
			),
		),
		Sub(
			Exec(testDump(t, "five")),
			Usage("five", ""),
			Flags().
				String("val", "", Short("l"), Default(125)).
				Slice("cidr", "", Short("c"), CIDRT).
				Slice("url", "", Short("u"), URLT).
				Slice("date", "", Short("d"), DateT).
				Map("time", "", Short("t"), TimeT).
				Map("timemap", "", Short("T"), MapKey(Uint64T), TimeT).
				Map("countmap", "", Short("C"), MapKey(Float64T), CountT),
		),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return cmd
}

func testDump(t *testing.T, name string) func(context.Context, []string) {
	t.Helper()
	return func(ctx context.Context, args []string) {
		c, ok := Ctx(ctx)
		if !ok {
			t.Fatalf("expected non-nil context")
		}
		_, _ = fmt.Fprintln(c.Stdout, "exec:", name)
		_, _ = fmt.Fprintln(c.Stdout, "name:", c.Exec.Name)
		_, _ = fmt.Fprintln(c.Stdout, "path:", c.Exec.Path())
		_, _ = fmt.Fprintln(c.Stdout, "args:", args)
		_, _ = fmt.Fprintln(c.Stdout, "vars:", c.Vars)
	}
}
