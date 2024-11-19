// Package uuid provides a kobra type for uuid processing.
package uuid

import (
	"context"

	"github.com/google/uuid"
	"github.com/xo/kobra"
)

func init() {
	kobra.RegisterType(kobra.UUIDT, kobra.NewTypeDesc(kobra.NewText(func() (any, error) {
		return new(uuid.UUID), nil
	}, kobra.UUIDT)))
}

// UUID returns the uuid var from the context.
func UUID(ctx context.Context, name string) *uuid.UUID {
	return kobra.Get[*uuid.UUID](ctx, name)
}
