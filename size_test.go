package ox

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestSize(t *testing.T) {
	tests := []struct {
		v     Size
		as    string
		exp   string
		parse bool
	}{
		{1000, "%d", "1000", true},
		{1000, "%s", "1000 B", true},
		{1000, "%S", "1 kB", true},
		{-1000, "%s", "-1000 B", true},
		{-1000, "%S", "-1 kB", true},
		{1000, "%f", "1000 B", true},
		{1024, "%d", "1024", true},
		{1024, "%s", "1 KiB", true},
		{1024, "%f", "1 KiB", true},
		{1024, "%e", "1Ki", true},
		{1024, "%z", "1KiB", true},
		{1024, "% e", "1 Ki", true},
		{1024, "% z", "1 KiB", true},
		{1024, "%.6z", "1.000000KiB", true},
		{1024, "%.0z", "1KiB", true},
		{1024, "%.1z", "1.0KiB", true},
		{1024, "%.2z", "1.00KiB", true},
		{1024, "%.3z", "1.000KiB", true},
		{1024, "%.6e", "1.000000Ki", true},
		{1024, "%.0e", "1Ki", true},
		{1024, "%.1e", "1.0Ki", true},
		{1024, "%.2e", "1.00Ki", true},
		{1024, "%.3e", "1.000Ki", true},
		{12345678, "%d", "12345678", true},
		{12345678, "%D", "12345678", true},
		{12345678, "%S", "12.35 MB", false},
		{12345678, "%m", "11.77MiB", false},
		{12345678, "%.3m", "11.774MiB", false},
		{12345678, "% .3m", "11.774 MiB", false},
		{12345678, "%d", "12345678", true},
		{12345678, "%s", "11.77 MiB", false},
		{12345678, "%f", "11.77 MiB", false},
		{12345678, "%e", "11.77Mi", false},
		{12345678, "%E", "12.35M", false},
		{12345678, "%z", "11.77MiB", false},
		{3*MiB + 100*KiB, "%d", "3248128", true},
		{3*MiB + 100*KiB, "%s", "3.1 MiB", false},
		{3*MiB + 100*KiB, "%F", "3.25 MB", false},
		{2 * GiB, "%d", "2147483648", true},
		{2 * GiB, "%s", "2 GiB", true},
		{2 * GiB, "%f", "2 GiB", true},
		{4 * TiB, "%d", "4398046511104", true},
		{4 * TiB, "%s", "4 TiB", true},
		{5*PiB + 200*TiB, "%d", "5849401859768320", true},
		{5*PiB + 200*TiB, "% m", "5578424320 MiB", true},
		{5*PiB + 200*TiB, "% g", "5447680 GiB", true},
		{5*PiB + 200*TiB, "% t", "5320 TiB", true},
		{5*PiB + 200*TiB, "% b", "5.2 PiB", false},
		{5*PiB + 200*TiB, "%v", "5.2 PiB", false},
		{5*PiB + 200*TiB, "%V", "5.85 PB", false},
		{5*PiB + 200*TiB, "% x", "%x(error=unknown size verb)", false},
		{100, "%e", "100", true},
		{100, "% e", "100", true},
		{100, "% E", "100", true},
		{-100, "% e", "-100", true},
		{0, "% e", "0", true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := fmt.Sprintf(test.as, test.v)
			t.Logf("%d %q :: %q == %q", test.v, test.as, test.exp, s)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			if !test.parse {
				return
			}
			sz, err := ParseSize(s)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if sz != test.v {
				t.Errorf("expected %s, got: %s", test.v, sz)
			}
		})
	}
}

// TestSizeExact holds that q and Q write the exact size, as Kubernetes writes
// a quantity, and that ParseSize reads it back to the same size.
func TestSizeExact(t *testing.T) {
	tests := []struct {
		v   Size
		as  string
		exp string
	}{
		{0, "%q", "0"},
		{0, "% q", "0"},
		{0, "%Q", "0"},
		{1, "%q", "1"},
		{100, "% q", "100"},
		{100, "% Q", "100"},
		{1000, "%q", "1000"},
		{1000, "%Q", "1k"},
		{1000, "% Q", "1 k"},
		{1024, "%q", "1Ki"},
		{1024, "% q", "1 Ki"},
		{1024, "%Q", "1024"},
		{1536, "%q", "1536"},
		{1536 * KiB, "%q", "1536Ki"},
		{6 * GiB, "%q", "6Gi"},
		{6 * GiB, "% q", "6 Gi"},
		{6 * GiB, "%Q", "6442450944"},
		{GiB + GiB/2, "%q", "1536Mi"},
		{GiB + GiB/2, "%.0q", "1536Mi"},
		{1500 * MiB, "%q", "1500Mi"},
		{6 * GB, "%q", "5859375Ki"},
		{6 * GB, "%Q", "6G"},
		{6 * GB, "% Q", "6 G"},
		{1500 * MB, "%Q", "1500M"},
		{-1536 * MiB, "%q", "-1536Mi"},
		{-6 * GB, "%Q", "-6G"},
		{12345678, "%q", "12345678"},
		{12345678, "%Q", "12345678"},
		{5*PiB + 200*TiB, "%q", "5320Ti"},
		{4 * TB, "%Q", "4T"},
		{3 * PB, "%Q", "3P"},
		{7 * EiB, "%q", "7Ei"},
		{EB, "%Q", "1E"},
		{2048 * PiB, "%q", "2Ei"},
		{1 << 53, "%q", "8Pi"},
		{1<<53 + 1, "%q", "9007199254740993"},
		{1<<53 + 1, "%Q", "9007199254740993"},
		{(1<<52 + 1) * KiB, "%q", "4503599627370497Ki"},
		{math.MaxInt64, "%q", "9223372036854775807"},
		{math.MaxInt64, "%Q", "9223372036854775807"},
		{math.MaxInt64 - KiB + 1, "%q", "9007199254740991Ki"},
		{math.MinInt64, "%q", "-8Ei"},
		{math.MinInt64, "%Q", "-9223372036854775808"},
	}
	for _, test := range tests {
		t.Run(test.as+" "+strconv.FormatInt(int64(test.v), 10), func(t *testing.T) {
			s := fmt.Sprintf(test.as, test.v)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
			sz, err := ParseSize(s)
			switch {
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case sz != test.v:
				t.Errorf("expected %d, got: %d", test.v, sz)
			}
			for _, unit := range []time.Duration{time.Second, time.Minute} {
				rate := Rate{test.v, unit}
				s := fmt.Sprintf(test.as, rate)
				if exp := test.exp + "/" + unitString(unit); s != exp {
					t.Errorf("expected %q, got: %q", exp, s)
				}
				r, err := ParseRate(s)
				switch {
				case err != nil:
					t.Fatalf("expected no error, got: %v", err)
				case r != rate:
					t.Errorf("expected %v, got: %v", rate, r)
				}
			}
		})
	}
}

// TestParseSize holds the forms that ParseSize reads and refuses.
func TestParseSize(t *testing.T) {
	tests := []struct {
		s   string
		exp Size
		err error
	}{
		{"6Gi", 6 * GiB, nil},
		{"6 Gi", 6 * GiB, nil},
		{"6G", 6 * GB, nil},
		{"1.5Gi", GiB + GiB/2, nil},
		{"+1k", KB, nil},
		{"-1Ki", -KiB, nil},
		{"6e9", 6 * GB, nil},
		{"6E9", 6 * GB, nil},
		{"1E3", KB, nil},
		{"1e+3", KB, nil},
		{"1.5e3", 1500, nil},
		{"-2e3", -2000, nil},
		{"15e-1", 1, nil},
		{"0e0", 0, nil},
		{"6e", 6 * EB, nil},
		{"6E", 6 * EB, nil},
		{"6 e", 6 * EB, nil},
		{"1m", MB, nil},
		{"9223372036854775807", math.MaxInt64, nil},
		{"-9223372036854775808", math.MinInt64, nil},
		{"8191Pi", 8191 * PiB, nil},
		{"-8Ei", math.MinInt64, nil},
		{"6e9Gi", 0, ErrInvalidSize},
		{"6 e9", 0, ErrInvalidSize},
		{"6e9.5", 0, ErrInvalidSize},
		{"e9", 0, ErrInvalidSize},
		{"6x", 0, ErrUnknownSize},
		{"9223372036854775808", 0, strconv.ErrRange},
		{"-9223372036854775809", 0, strconv.ErrRange},
		{"8Ei", 0, strconv.ErrRange},
		{"18446744073709551616", 0, strconv.ErrRange},
		{"9.3e18", 0, strconv.ErrRange},
		{"-9.3e18", 0, strconv.ErrRange},
		{"1e400", 0, strconv.ErrRange},
		{"9Ei", 0, strconv.ErrRange},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			sz, err := ParseSize(test.s)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case test.err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.err != nil && !errors.Is(err, ErrInvalidSize):
				t.Fatalf("expected error %v, got: %v", ErrInvalidSize, err)
			case sz != test.exp:
				t.Errorf("expected %d, got: %d", test.exp, sz)
			}
		})
	}
}

// TestConvOverflow holds that a uint64 that does not fit in an int64 is an
// error, and not a negative duration, size or rate.
func TestConvOverflow(t *testing.T) {
	const big = uint64(math.MaxInt64) + 1
	tests := []struct {
		name string
		f    func(any) (any, error)
		v    any
		exp  any
		err  error
	}{
		{"duration", wrapConv(asDuration), big, time.Duration(0), ErrInvalidConversion},
		{"duration that fits", wrapConv(asDuration), uint64(2), 2 * time.Second, nil},
		{"size", wrapConv(asSize), big, Size(0), ErrInvalidConversion},
		{"size that fits", wrapConv(asSize), uint64(2), Size(2), nil},
		{"rate", wrapConv(asRate), big, Rate{}, ErrInvalidConversion},
		{"rate that fits", wrapConv(asRate), uint64(2), Rate{2, DefaultRateUnit}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v, err := test.f(test.v)
			switch {
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case v != test.exp:
				t.Errorf("expected %v, got: %v", test.exp, v)
			}
		})
	}
}

// wrapConv returns f as a func that returns any.
func wrapConv[T any](f func(any) (T, error)) func(any) (any, error) {
	return func(v any) (any, error) {
		return f(v)
	}
}

// TestSizeOverflow holds that a size or a duration out of the range of an
// int64 is an error, and that the lowest size formats with one sign.
// FuzzTypeValue found that 1e20 gave the lowest int64, which help wrote as
// --9223372036854775808 B.
func TestSizeOverflow(t *testing.T) {
	for _, test := range []struct {
		typ Type
		s   string
	}{
		{SizeT, "1e20"},
		{SizeT, "-1e20"},
		{SizeT, "NaN"},
		{SizeT, "+Inf"},
		{DurationT, "1e20"},
		{DurationT, "NaN"},
		{DurationT, "9999999999999"},
		{DurationT, "-9999999999999"},
		{DurationT, "18446744073709551615"},
	} {
		t.Run(test.typ.String()+" "+test.s, func(t *testing.T) {
			v, err := test.typ.New()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err := v.Set(test.s); !errors.Is(err, ErrInvalidValue) {
				t.Errorf("expected %v, got: %v", ErrInvalidValue, err)
			}
		})
	}
	for _, test := range []struct {
		size Size
		verb rune
		exp  string
	}{
		{math.MinInt64, 'f', "-8 EiB"},
		{math.MinInt64, 'F', "-9.22 EB"},
		{math.MinInt64, 'q', "-8 Ei"},
		{math.MaxInt64, 'f', "8 EiB"},
	} {
		t.Run(fmt.Sprintf("format %d %c", test.size, test.verb), func(t *testing.T) {
			if s := FormatSize(int64(test.size), test.verb, -2, true); s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}
