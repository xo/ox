//go:build go1.27

package ox

import (
	"reflect"
	"uuid"
)

// The standard uuid package is new in Go 1.27. A program that builds with an
// older Go has no [UUIDT] type. See D24.
func init() {
	RegisterTypeName(UUIDT, "*uuid.UUID")
	valueTypes[reflect.TypeFor[uuid.UUID]()] = UUIDT
	RegisterTextType(func() (*uuid.UUID, error) {
		return new(uuid.UUID), nil
	})
}
