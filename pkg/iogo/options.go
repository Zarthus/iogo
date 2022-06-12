package iogo

import (
	"github.com/zarthus/iogo/v2/pkg/term/col"
)

// Options contain a list of options - not all options are respected at every time or for every method,
// this depends on the implementation of the Reader or Writer
type Options struct {
	// If no input is inserted, default to this value
	Default *string

	// Some output can be styled with text colour
	FgColour *col.Colour
	BgColour *col.Colour
}
