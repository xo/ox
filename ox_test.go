package ox

import (
	"bytes"
	"context"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestSuggestions(t *testing.T) {
	for i, test := range []struct {
		args []string
		exp  string
	}{
		{
			ss(`subfoo`), `error: unknown command "subfoo" for "cmd"`,
		},
		{
			ss(`on`), `error: unknown command "on" for "cmd"

Did you mean this?
  one
`,
		},
		{
			ss(`remove`), `error: unknown command "remove" for "cmd"

Did you mean this?
  one
`,
		},
		{
			ss(`rmove`), `error: unknown command "rmove" for "cmd"

Did you mean this?
  one
`,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var code int
			c := testContext(t, &code, test.args...)
			err := c.Run(context.Background())
			if err == nil {
				t.Fatalf("expected non-nil error")
			}
			if !c.Handler(err) {
				t.Fatalf("expected Handler to return true")
			}
			if code == 0 {
				t.Fatalf("expected Handler to set non-zero code")
			}
			if s := strings.TrimSuffix(c.Stderr.(*bytes.Buffer).String(), "\n"); s != test.exp {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", test.exp, s)
			}
		})
	}
}

func TestLdist(t *testing.T) {
	tests := []struct {
		a   string
		b   string
		exp int
	}{
		{"", "", 0},
		{"", "a", 1},
		{"ab", "aa", 1},
		{"ab", "aaa", 2},
		{"ab", "ba", 2},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
		{"bar", "br", 1},
		{"bbb", "a", 3},
		{"book", "back", 2},
		{"br", "bar", 1},
		{"distance", "difference", 5},
		{"fod", "Food", 1},
		{"foo", "", 3},
		{"foo", "bar", 3},
		{"foo", "Food", 1},
		{"Hafþór Júlíus Björnsson", "Hafþor Julius Bjornsson", 4},
		{"", "hello", 5},
		{"hello", "hello", 0},
		{"kitten", "sitting", 3},
		{"levenshtein", "frankenstein", 6},
		{"mississippi", "swiss miss", 8},
		{"resume and cafe", "resumé and café", 2},
		{"resume and cafe", "resumes and cafes", 2},
		{"resumé and café", "resumés and cafés", 2},
		{"rosettacode", "raisethysword", 8},
		{"saturday", "sunday", 3},
		{"sitten", "sittin", 1},
		{"sittin", "sitting", 1},
		{"sleep", "fleeting", 5},
		{"stop", "tops", 2},
		{"test", "t", 3},
		{"།་གམ་འས་པ་་མ།", "།་གམའས་པ་་མ", 2},
		{"👀😀", "😀", 1},
	}
	for _, test := range tests {
		t.Run(test.a+"::"+test.b, func(t *testing.T) {
			a, b := []rune(strings.ToLower(test.a)), []rune(strings.ToLower(test.b))
			if i := Ldist(a, b); i != test.exp {
				t.Errorf("expected %d, got: %d", test.exp, i)
			}
			if i := Ldist(b, a); i != test.exp {
				t.Errorf("expected %d, got: %d", test.exp, i)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		size int64
		prec int
		iec  bool
		exp  string
	}{
		{0, DefaultPrec, false, "0 B"},
		{0, DefaultPrec, true, "0 B"},
		{int64(0.754 * float64(MB)), DefaultPrec, false, "754 kB"},
		{int64(1.5 * float64(GB)), 2, false, "1.50 GB"},
		{int64(1.5 * float64(GB)), 1, false, "1.5 GB"},
		{int64(1.54 * float64(GB)), 2, false, "1.54 GB"},
		{int64(1.5 * float64(GiB)), 2, true, "1.50 GiB"},
		{int64(1.5 * float64(PiB)), 2, true, "1.50 PiB"},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := FormatSize(test.size, test.prec, test.iec)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			size, prec, iec, err := ParseSize(test.exp)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if size != test.size {
				t.Errorf("expected size %d, got: %d", test.size, size)
			}
			if prec != test.prec {
				t.Errorf("expected prec %d, got: %d", test.prec, prec)
			}
			if test.size != 0 && iec != test.iec {
				t.Errorf("expected iec %t, got: %t", test.iec, iec)
			}
		})
	}
}

func TestFormatRate(t *testing.T) {
	tests := []struct {
		rate int64
		prec int
		iec  bool
		unit time.Duration
		exp  string
	}{
		{0, DefaultPrec, false, time.Second, "0 B/s"},
		{0, DefaultPrec, true, time.Second, "0 B/s"},
		{int64(0.754 * float64(MB)), DefaultPrec, false, time.Hour, "754 kB/h"},
		{int64(1.5 * float64(GB)), 1, false, time.Microsecond, "1.5 GB/us"},
		{int64(1.54 * float64(GB)), 2, false, time.Microsecond, "1.54 GB/us"},
		{int64(1.5 * float64(GiB)), 2, true, time.Millisecond, "1.50 GiB/ms"},
		{int64(1.5 * float64(PiB)), 2, true, time.Second, "1.50 PiB/s"},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := FormatRate(test.rate, test.prec, test.iec, test.unit)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			rate, prec, iec, unit, err := ParseRate(test.exp)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if rate != test.rate {
				t.Errorf("expected rate %d, got: %d", test.rate, rate)
			}
			if prec != test.prec {
				t.Errorf("expected prec %d, got: %d", test.prec, prec)
			}
			if test.rate != 0 && iec != test.iec {
				t.Errorf("expected iec %t, got: %t", test.iec, iec)
			}
			if unit != test.unit {
				t.Errorf("expected unit %v, got: %v", test.unit, unit)
			}
		})
	}
}

func testContext(t *testing.T, code *int, args ...string) *Context {
	t.Helper()
	cmd := testCommand(t)
	cmd.Exec = nil
	c := &Context{
		Exit: func(c int) {
			*code = c
		},
		Panic: func(v any) {
			t.Fatalf("context panic: %v", v)
		},
		Root:   cmd,
		Stderr: new(bytes.Buffer),
		Args:   args,
		Vars:   make(Vars),
	}
	c.Handler = DefaultErrorHandler(c)
	if err := c.Parse(); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return c
}
