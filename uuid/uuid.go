// Package uuid provides a ox type for uuid processing.
package uuid

import (
	"context"

	"github.com/google/uuid"
	"github.com/xo/ox"
)

func init() {
	ox.RegisterType(ox.UUIDT, ox.NewTypeDesc(ox.NewText(func() (any, error) {
		return new(uuid.UUID), nil
	}, ox.UUIDT)))
}

// UUID returns the uuid var from the context.
func UUID(ctx context.Context, name string) *uuid.UUID {
	return ox.Get[*uuid.UUID](ctx, name)
}
