//go:build go1.27

package ox

import (
	"reflect"
	"testing"
	"uuid"
)

func TestTypeNewUUID(t *testing.T) {
	testTypeNew(t, []typeTest{
		{
			UUIDT, []test{
				{"", "00000000-0000-0000-0000-000000000000"},
				{"", mustUUID(t, "00000000-0000-0000-0000-000000000000")},
				{"00000000-0000-0000-0000-000000000000", mustUUID(t, "00000000-0000-0000-0000-000000000000")},
				{"00000000000000000000000000000000", mustUUID(t, "00000000-0000-0000-0000-000000000000")},
				{"f47ac10b-58cc-0372-8567-0e02b2c3d479", mustUUID(t, "f47ac10b-58cc-0372-8567-0e02b2c3d479")},
				{"F47AC10B-58CC-1372-8567-0E02B2C3D479", mustUUID(t, "f47ac10b-58cc-1372-8567-0e02b2c3d479")},
				{"f47ac10b-58cc-2372-8567-0e02b2c3d479", mustUUID(t, "f47ac10b-58cc-2372-8567-0e02b2c3d479")},
				{"f47ac10b58cc337285670e02b2c3d479", mustUUID(t, "f47ac10b-58cc-3372-8567-0e02b2c3d479")},
				{"{f47ac10b-58cc-4372-8567-0e02b2c3d479}", mustUUID(t, "f47ac10b-58cc-4372-8567-0e02b2c3d479")},
				{"{F47AC10B-58CC-4372-8567-0E02B2C3D479}", mustUUID(t, "f47ac10b-58cc-4372-8567-0e02b2c3d479")},
				{"0", ErrInvalidValue},
				{":", ErrInvalidValue},
			},
		},
	})
}

// TestUUIDTypeLookup tests that type_go127.go registers each lookup of the
// uuid type that type.go used to hold.
func TestUUIDTypeLookup(t *testing.T) {
	typ, _, _, err := defaultType(reflect.TypeFor[uuid.UUID]())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		name string
		typ  Type
	}{
		{"defaultType", typ},
		{"typeRef", typeRef(new(uuid.UUID))},
		{"reflectType", reflectType(reflect.TypeFor[*uuid.UUID]())},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.typ != UUIDT {
				t.Errorf("expected %s, got: %s", UUIDT, test.typ)
			}
		})
	}
}

func mustUUID(t *testing.T, s string) uuid.UUID {
	t.Helper()
	u, err := uuid.Parse(s)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return u
}
