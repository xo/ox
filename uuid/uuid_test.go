package uuid

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/xo/ox"
)

func TestUUID(t *testing.T) {
	for _, exp := range uuidTests() {
		t.Run(exp, func(t *testing.T) {
			v, err := ox.UUIDT.New()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if err := v.Set(context.Background(), exp); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			u, err := ox.As[*uuid.UUID](v)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s := u.String(); s != exp {
				t.Errorf("expected %s, got: %s", exp, s)
			}
			t.Logf("u: %v", u)
		})
	}
}

func uuidTests() []string {
	return []string{
		"f47ac10b-58cc-0372-8567-0e02b2c3d479",
		"f47ac10b-58cc-1372-8567-0e02b2c3d479",
		"f47ac10b-58cc-2372-8567-0e02b2c3d479",
		"f47ac10b-58cc-3372-8567-0e02b2c3d479",
		"f47ac10b-58cc-4372-8567-0e02b2c3d479",
		"f47ac10b-58cc-5372-8567-0e02b2c3d479",
		"f47ac10b-58cc-6372-8567-0e02b2c3d479",
		"f47ac10b-58cc-7372-8567-0e02b2c3d479",
		"f47ac10b-58cc-8372-8567-0e02b2c3d479",
		"f47ac10b-58cc-9372-8567-0e02b2c3d479",
		"f47ac10b-58cc-a372-8567-0e02b2c3d479",
		"f47ac10b-58cc-b372-8567-0e02b2c3d479",
		"f47ac10b-58cc-c372-8567-0e02b2c3d479",
		"f47ac10b-58cc-d372-8567-0e02b2c3d479",
		"f47ac10b-58cc-e372-8567-0e02b2c3d479",
		"f47ac10b-58cc-f372-8567-0e02b2c3d479",
	}
}
