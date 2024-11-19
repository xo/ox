package kobra

import (
	"io"
)

// HelpDesc holds help command descriptions.
type HelpDesc struct {
	Sort func(*Command, *Command) bool
}

// NewHelp
func NewHelp(opts ...Option) (*HelpDesc, error) {
	return nil, nil
}

func DefaultHelp(opts ...Option) Option {
	return nil
}

// Writer returns
func (h *HelpDesc) Writer() func(io.Writer) (int64, error) {
	return nil
}
