package glob

import (
	"testing"

	"github.com/xo/ox"
)

func TestGLOB(t *testing.T) {
	for _, exp := range globTests() {
		t.Run(exp, func(t *testing.T) {
			v, err := ox.GlobT.New()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err := v.Set(exp); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			val, err := ox.As[*GlobValue](v)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if val == nil {
				t.Fatalf("expected non-nil value")
			}
			if s := val.String(); s != exp {
				t.Errorf("expected %s, got: %s", exp, s)
			}
			t.Logf("u: %v", val)
		})
	}
}

func globTests() []string {
	return []string{
		"",
		"config.{toml,yaml}",
		"file.txt",
		"[A-Z]*{,_test}.go",
	}
}
