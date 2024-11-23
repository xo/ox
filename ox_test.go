package ox

import (
	"strings"
	"testing"
)

func TestLdist(t *testing.T) {
	tests := []struct {
		a   string
		b   string
		exp int
	}{
		{"", "", 0},
		{"", "a", 1},
		{"a", "", 1},
		{"foo", "", 3},
		{"foo", "bar", 3},
		{"foo", "food", 1},
		{"Food", "foo", 1},
		{"fod", "Food", 1},
		{"br", "bar", 1},
		{"bar", "br", 1},
		{"test", "t", 3},
		{"book", "back", 2},
		{"kitten", "sitting", 3},
		{"Kitten", "Sitting", 3},
	}
	for _, test := range tests {
		t.Run(test.a+"::"+test.b, func(t *testing.T) {
			a, b := []rune(strings.ToLower(test.a)), []rune(strings.ToLower(test.b))
			if i := ldist(a, b); i != test.exp {
				t.Errorf("expected %d, got: %d", test.exp, i)
			}
		})
	}
}
