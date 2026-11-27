package ox

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

// TestAsValueSlice holds the conversion of a value to a []byte and a []rune
// (B12).
func TestAsValueSlice(t *testing.T) {
	type myBytes []byte
	tests := []struct {
		name string
		v    any
		val  any
		exp  any
		err  error
	}{
		{"bytes", new([]byte), "abc", []byte("abc"), nil},
		{"bytes from bytes", new([]byte), []byte("x"), []byte("x"), nil},
		{"runes", new([]rune), "añb", []rune("añb"), nil},
		{"named bytes", new(myBytes), "abc", myBytes("abc"), nil},
		{"ints", new([]int), "abc", []int(nil), ErrInvalidConversion},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := reflect.ValueOf(test.v)
			err := asValue(v, test.val)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			}
			if got := v.Elem().Interface(); !reflect.DeepEqual(got, test.exp) {
				t.Errorf("expected %#v, got: %#v", test.exp, got)
			}
		})
	}
}

// TestBindBytes holds that Bind sets a []byte, a [][]byte and a map of
// []byte from the command line (B12).
func TestBindBytes(t *testing.T) {
	var (
		b   []byte
		bs  [][]byte
		m   map[string][]byte
		r   [][]rune
		got = func() any { return []any{b, bs, m, r} }
	)
	tests := []struct {
		name string
		args []string
		exp  any
	}{
		{
			"each", []string{"--b", "abc", "--bs", "x", "--bs", "y", "--m", "k=v", "--r", "añ"},
			[]any{[]byte("abc"), [][]byte{[]byte("x"), []byte("y")}, map[string][]byte{"k": []byte("v")}, [][]rune{[]rune("añ")}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				Flags().
					String("b", "bytes", Bind(&b)).
					Slice("bs", "slice of bytes", Bind(&bs)).
					Map("m", "map of bytes", Bind(&m)).
					Slice("r", "slice of runes", Bind(&r)),
				Exec(func() {}),
				Args(test.args...),
			)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if !reflect.DeepEqual(got(), test.exp) {
				t.Errorf("expected %#v, got: %#v", test.exp, got())
			}
		})
	}
}

// TestBindTypes holds that Bind and From set the types of B49.
func TestBindTypes(t *testing.T) {
	var v struct {
		Timeout time.Duration   `ox:"a timeout,short:t"`
		Size    Size            `ox:"a size"`
		Rate    Rate            `ox:"a rate,type:rate"`
		Rate2   Rate            `ox:"a rate"`
		When    time.Time       `ox:"a date,type:date"`
		Stamp   time.Time       `ox:"a timestamp"`
		Verbose int             `ox:"verbosity,type:count,short:v"`
		C       complex64       `ox:"a complex"`
		Waits   []time.Duration `ox:"waits"`
	}
	var d time.Duration
	var n uint
	err := RunError(
		context.Background(),
		Usage("app", "an app"),
		Flags().
			Duration("d", "a duration", Bind(&d)).
			Count("count", "a count", Short("n"), Bind(&n)),
		From(&v),
		Exec(func() {}),
		Args(
			"-t", "5s", "--size", "2 MiB", "--rate", "1 MiB/s", "--rate2", "2 MiB/s",
			"--when", "2026-11-27", "--stamp", "2026-11-27T10:11:12Z", "-vvv", "--c", "1+2i",
			"--waits", "1s", "--waits", "2m", "--d", "1h", "-nn",
		),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		name string
		v    any
		exp  any
	}{
		{"duration", v.Timeout, 5 * time.Second},
		{"size", v.Size, Size(2 * MiB)},
		{"rate", v.Rate, NewRate(1*MiB, time.Second)},
		{"rate without a type", v.Rate2, NewRate(2*MiB, time.Second)},
		{"date", v.When, time.Date(2026, 11, 27, 0, 0, 0, 0, time.UTC)},
		{"timestamp", v.Stamp, time.Date(2026, 11, 27, 10, 11, 12, 0, time.UTC)},
		{"count", v.Verbose, 3},
		{"complex64", v.C, complex64(1 + 2i)},
		{"durations", v.Waits, []time.Duration{time.Second, 2 * time.Minute}},
		{"bound duration", d, time.Hour},
		{"bound count", n, uint(2)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.v, test.exp) {
				t.Errorf("expected %v, got: %v", test.exp, test.v)
			}
		})
	}
}

// testText is a type whose pointer has the text marshal pair, and that no
// program registers.
type testText struct {
	s string
}

func (v *testText) UnmarshalText(b []byte) error {
	if string(b) == "bad" {
		return errors.New("bad text")
	}
	v.s = "T:" + string(b)
	return nil
}

func (v *testText) MarshalText() ([]byte, error) {
	return []byte(v.s), nil
}

// TestFromMarshaler holds that From and Bind take a pointer type with the
// text marshal pair that no program registers (B19 item 7).
func TestFromMarshaler(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  string
		err  error
	}{
		{"a value", []string{"--text", "x"}, "T:x", nil},
		{"no value", nil, "", nil},
		{"a bad value", []string{"--text", "bad"}, "", ErrInvalidValue},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var v struct {
				Text *testText `ox:"a text"`
			}
			var s string
			err := RunError(
				context.Background(),
				Usage("app", "an app"),
				From(&v),
				Exec(func(ctx context.Context) {
					c, _ := Ctx(ctx)
					if val, ok := c.Vars["text"]; ok {
						s = val.String()
					}
				}),
				Args(test.args...),
			)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.err != nil:
				return
			case s != test.exp:
				t.Errorf("expected %q from the value, got: %q", test.exp, s)
			case test.exp != "" && (v.Text == nil || v.Text.s != test.exp):
				t.Errorf("expected %q in the field, got: %v", test.exp, v.Text)
			}
		})
	}
}
