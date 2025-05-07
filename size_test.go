package ox

import (
	"fmt"
	"strconv"
	"testing"
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
		{1024, "%.6z", "1.000000KiB", true},
		{1024, "%.0z", "1KiB", true},
		{1024, "%.1z", "1.0KiB", true},
		{1024, "%.2z", "1.00KiB", true},
		{1024, "%.3z", "1.000KiB", true},
		{12345678, "%d", "12345678", true},
		{12345678, "%D", "12345678", true},
		{12345678, "%s", "11.77 MiB", false},
		{12345678, "%m", "11.77MiB", false},
		{12345678, "%.3m", "11.774MiB", false},
		{12345678, "% .3m", "11.774 MiB", false},
		{12345678, "%d", "12345678", true},
		{12345678, "%s", "11.77 MiB", false},
		{12345678, "%f", "11.77 MiB", false},
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
