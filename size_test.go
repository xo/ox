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
