// Package uuid provides a ox type for uuid processing.
package uuid

import (
	"context"

	"github.com/google/uuid"
	"github.com/xo/ox"
	"github.com/xo/ox/otx"
)

func init() {
	ox.RegisterTypeName(ox.UUIDT, "*uuid.UUID")
	ox.RegisterTextType(New)
}

// New creates a new uuid.
func New() (*uuid.UUID, error) {
	return new(uuid.UUID), nil
}

// UUID returns the uuid var from the context.
func UUID(ctx context.Context, name string) *uuid.UUID {
	return otx.Get[*uuid.UUID](ctx, name)
}
