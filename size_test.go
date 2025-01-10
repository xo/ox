package ox

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		v   Size
		as  string
		exp string
	}{
		{1000, "%d", "1000"},
		{1000, "%s", "1000 B"},
		{1000, "%S", "1 kB"},
		{-1000, "%s", "-1000 B"},
		{-1000, "%S", "-1 kB"},
		{1000, "%f", "1000 B"},
		{1024, "%d", "1024"},
		{1024, "%s", "1 KiB"},
		{1024, "%f", "1 KiB"},
		{1024, "%.6z", "1.000000KiB"},
		{1024, "%.0z", "1KiB"},
		{1024, "%.1z", "1.0KiB"},
		{1024, "%.2z", "1.00KiB"},
		{1024, "%.3z", "1.000KiB"},
		{12345678, "%d", "12345678"},
		{12345678, "%D", "12345678"},
		{12345678, "%s", "11.77 MiB"},
		{12345678, "%m", "11.77MiB"},
		{12345678, "%.3m", "11.774MiB"},
		{12345678, "% .3m", "11.774 MiB"},
		{12345678, "%d", "12345678"},
		{12345678, "%s", "11.77 MiB"},
		{12345678, "%f", "11.77 MiB"},
		{3*MiB + 100*KiB, "%d", "3248128"},
		{3*MiB + 100*KiB, "%s", "3.1 MiB"},
		{3*MiB + 100*KiB, "%F", "3.25 MB"},
		{2 * GiB, "%d", "2147483648"},
		{2 * GiB, "%s", "2 GiB"},
		{2 * GiB, "%f", "2 GiB"},
		{4 * TiB, "%d", "4398046511104"},
		{4 * TiB, "%s", "4 TiB"},
		{5*PiB + 200*TiB, "%d", "5849401859768320"},
		{5*PiB + 200*TiB, "% m", "5578424320 MiB"},
		{5*PiB + 200*TiB, "% g", "5447680 GiB"},
		{5*PiB + 200*TiB, "% t", "5320 TiB"},
		{5*PiB + 200*TiB, "% b", "5.2 PiB"},
		{5*PiB + 200*TiB, "%v", "5.2 PiB"},
		{5*PiB + 200*TiB, "%V", "5.85 PB"},
		{5*PiB + 200*TiB, "% x", "%x(error=unknown size verb)"},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := fmt.Sprintf(test.as, test.v)
			t.Logf("%d %q :: %q == %q", test.v, test.as, test.exp, s)
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}
