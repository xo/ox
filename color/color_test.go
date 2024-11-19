package color

import (
	"context"
	"testing"

	"github.com/kenshaw/colors"
	"github.com/xo/kobra"
)

func TestColor(t *testing.T) {
	for _, test := range colorTests() {
		t.Run(test.s, func(t *testing.T) {
			v, err := kobra.ColorT.New()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err := v.Set(context.Background(), test.s); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			c, err := kobra.As[*colors.Color](v)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := c.AsText(); s != test.exp {
				t.Errorf("expected %s, got: %s", test.exp, s)
			}
			t.Logf("c: %v", c)
		})
	}
}

type colorTest struct {
	s   string
	exp string
}

func colorTests() []colorTest {
	return []colorTest{
		{"red", "red"},
		{"BLACK", "black"},
		{"rgb(0,255,0)", "lime"},
		{"rgba( 0, 255, 0, 255)", "lime"},
		{"hex(0,ff,0)", "lime"},
		{"hex(0,ff,0,ff)", "lime"},
		{"#00ff00", "lime"},
		{"#00ff00ff", "lime"},
	}
}
