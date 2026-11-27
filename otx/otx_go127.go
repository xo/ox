//go:build go1.27

package otx

import (
	"context"
	"uuid"
)

// UUID returns the [uuid.UUID] variable from the context.
//
// The standard uuid package is new in Go 1.27, so a program that builds with
// an older Go has no UUID func.
func UUID(ctx context.Context, name string) *uuid.UUID {
	return Get[*uuid.UUID](ctx, name)
}
