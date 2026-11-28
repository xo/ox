package ox

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestStrftimeLayout(t *testing.T) {
	tests := []struct {
		layout string
		exp    string
		err    string
	}{
		{"%Y-%m-%d", "2006-01-02", ""},
		{"%d/%m/%Y %H:%M", "02/01/2006 15:04", ""},
		{"%y%m%d", "060102", ""},
		{"%b %B", "Jan January", ""},
		{"%a %A", "Mon Monday", ""},
		{"%e", "_2", ""},
		{"%j", "002", ""},
		{"%Y-%j", "2006-002", ""},
		{"%I:%M %p", "03:04 PM", ""},
		{"%H:%M:%S.%f", "15:04:05.000000", ""},
		{"%H:%M:%S,%f", "15:04:05,000000", ""},
		{"%H:%M:%S.%L", "15:04:05.000", ""},
		{"%z %Z", "-0700 MST", ""},
		{"%T", "15:04:05", ""},
		{"%R", "15:04", ""},
		{"%D", "01/02/06", ""},
		{"%F", "2006-01-02", ""},
		{"%F %T %z", "2006-01-02 15:04:05 -0700", ""},
		{"%%", "%", ""},
		{"%H%% %Y", "15% 2006", ""},
		{"100%% %Y", "", "Go reads the text"},
		{"%%Y", "%Y", ""},
		{"at %H h", "at 15 h", ""},
		{"%a day", "Mon day", ""},
		{"%Aday", "Mondayday", ""},
		{"%Y.%m", "2006.01", ""},
		{"", "", "empty layout"},
		{"%", "", "a % ends the layout"},
		{"%Y-%", "", "a % ends the layout"},
		{"%q", "", "%q is not supported"},
		{"%c", "", "%c is not supported"},
		{"%x", "", "%x is not supported"},
		{"%l", "", "%l is not supported"},
		{"%-d", "", "%- is not supported"},
		{"%E", "", "%E is not supported"},
		{"%é", "", "%é is not supported"},
		{"%S%f", "", "%f must follow a . or a ,"},
		{"%L", "", "%L must follow a . or a ,"},
		{"%Y-1", "", "Go reads the text"},
		{"%m1", "", "Go reads the text"},
		{"Jan %d", "", "Go reads the text"},
		{"%aday", "", "Go reads the text"},
		{"%buary", "", "Go reads the text"},
		{"_%e", "", "Go reads the text"},
		{"%H.0", "", "Go reads the text"},
		{"%S.9", "", "Go reads the text"},
		{"%H PM", "", "Go reads the text"},
		{"MST %H", "", "Go reads the text"},
		{"%Y 2006", "", "Go reads the text"},
	}
	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			s, err := strftimeLayout(test.layout)
			switch {
			case test.err != "" && !errors.Is(err, ErrInvalidLayout):
				t.Fatalf("expected error %v, got: %v", ErrInvalidLayout, err)
			case test.err != "" && !strings.Contains(err.Error(), test.err):
				t.Fatalf("expected error to contain %q, got: %v", test.err, err)
			case test.err == "" && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case s != test.exp:
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

// TestStrftimeFormat tests that each directive writes what strftime writes,
// and that the text reads back to the same time.
func TestStrftimeFormat(t *testing.T) {
	// 2026-03-05 is a Thursday, and the 64th day of the year
	tm := time.Date(2026, time.March, 5, 7, 8, 9, 12345678, time.FixedZone("IST", 5*3600+30*60))
	pm := time.Date(2026, time.December, 31, 23, 59, 59, 0, time.UTC)
	tests := []struct {
		layout string
		t      time.Time
		exp    string
	}{
		{"%Y %y %m %d", tm, "2026 26 03 05"},
		{"%b %B", tm, "Mar March"},
		{"%a %A", tm, "Thu Thursday"},
		{"[%e]", tm, "[ 5]"},
		{"%j", tm, "064"},
		{"%j", pm, "365"},
		{"%H %I %p", tm, "07 07 AM"},
		{"%H %I %p", pm, "23 11 PM"},
		{"%M %S", tm, "08 09"},
		{"%S.%f", tm, "09.012345"},
		{"%S,%L", tm, "09,012"},
		{"%z %Z", tm, "+0530 IST"},
		{"%z %Z", pm, "+0000 UTC"},
		{"%T", tm, "07:08:09"},
		{"%R", tm, "07:08"},
		{"%D", tm, "03/05/26"},
		{"%F", tm, "2026-03-05"},
		{"%H%%", tm, "07%"},
	}
	for _, test := range tests {
		t.Run(test.layout+"/"+test.exp, func(t *testing.T) {
			layout, err := strftimeLayout(test.layout)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			s := test.t.Format(layout)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			if _, err := time.Parse(layout, s); err != nil {
				t.Errorf("expected %q to read back, got: %v", s, err)
			}
		})
	}
}

func TestStrftimeFlag(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		args []string
		exp  string
		typ  Type
		err  error
	}{
		{
			"default",
			[]Option{Strftime("%d/%m/%Y"), Default("05/03/2026")},
			nil,
			"2026-03-05T00:00:00Z", TimestampT, nil,
		},
		{
			"command line",
			[]Option{Strftime("%d/%m/%Y %H:%M"), Default("05/03/2026 10:00")},
			[]string{"--at", "31/12/2026 23:59"},
			"2026-12-31T23:59:00Z", TimestampT, nil,
		},
		{
			"date keeps its type",
			[]Option{DateT, Strftime("%Y%m%d")},
			[]string{"--at", "20260305"},
			"2026-03-05T00:00:00Z", DateT, nil,
		},
		{
			"time keeps its type",
			[]Option{TimeT, Strftime("%I:%M %p")},
			[]string{"--at", "11:30 PM"},
			"0000-01-01T23:30:00Z", TimeT, nil,
		},
		{
			"zone",
			[]Option{Strftime("%F %T %z")},
			[]string{"--at", "2026-03-05 07:08:09 +0530"},
			"2026-03-05T07:08:09+05:30", TimestampT, nil,
		},
		{
			"bad value",
			[]Option{Strftime("%d/%m/%Y")},
			[]string{"--at", "2026-03-05"},
			"", "", ErrInvalidValue,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bound time.Time
			var got Vars
			var typ Type
			err := RunError(
				context.Background(),
				Usage("app", ""),
				Flags().Var("at", "a time", append(test.opts, Bind(&bound))...),
				Exec(func(ctx context.Context) {
					c, _ := Ctx(ctx)
					got = c.Vars
					typ = c.Root.Flags.Flags[0].Type
				}),
				Args(test.args...),
			)
			switch {
			case test.err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.err != nil:
				return
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			}
			if typ != test.typ {
				t.Errorf("expected type %s, got: %s", test.typ, typ)
			}
			v := To[time.Time](got["at"])
			if s := v.Format(time.RFC3339); s != test.exp {
				t.Errorf("expected %s, got: %s", test.exp, s)
			}
			if !bound.Equal(v) {
				t.Errorf("expected bound %s, got: %s", v, bound)
			}
		})
	}
}

func TestStrftimeFlagError(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		err  error
	}{
		{"bad layout", []Option{Strftime("%d/%m/%Q")}, ErrInvalidLayout},
		{"slice", []Option{SliceT, Strftime("%Y")}, ErrAppliedToInvalidType},
		{"array", []Option{ArrayT, Strftime("%Y")}, ErrAppliedToInvalidType},
		{"map", []Option{MapT, Strftime("%Y")}, ErrAppliedToInvalidType},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := NewFlag("at", "a time", test.opts...); !errors.Is(err, test.err) {
				t.Errorf("expected error %v, got: %v", test.err, err)
			}
		})
	}
}

func TestStrftimeHelp(t *testing.T) {
	var buf bytes.Buffer
	err := RunError(
		context.Background(),
		Usage("app", ""),
		Help(),
		Pipe(nil, &buf, &buf),
		Flags().Var("since", "the first day", Strftime("%d/%m/%Y"), Default("05/03/2026")),
		Args("--help"),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if exp := "--since timestamp  the first day (default: 05/03/2026)"; !strings.Contains(buf.String(), exp) {
		t.Errorf("expected help to contain %q, got:\n%s", exp, buf.String())
	}
}
