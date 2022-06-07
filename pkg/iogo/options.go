package iogo

import (
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
)

type KeyInputs rune

const (
	KeyUp   KeyInputs = 38
	KeyDown KeyInputs = 40
)

// Options contain a list of options - not all options are respected at every time or for every method,
// this depends on the implementation of the Reader or Writer
type Options struct {
	// If no input is inserted, default to this value
	Default *string

	// Avoid appending input to the internal history list
	// useful e.g. if you do not want to have passwords stored in memory for longer than necessary.
	DoNotTrack bool

	// Some output can be styled with text colour
	FgColour *term.Colour
	BgColour *term.Colour
}
