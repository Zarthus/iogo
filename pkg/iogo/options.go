package iogo

import (
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
)

type KeyInputs rune

const (
	KeyUp   KeyInputs = 38
	KeyDown KeyInputs = 40
)

type Options struct {
	// If no input is inserted, default to this value
	Default string

	// Avoid appending input to the internal history list
	// useful e.g. if you do not want to have passwords stored in memory for longer than necessary.
	DoNotTrack bool

	// Some output can be styled with text colour
	FgColor term.Colour
	BgColor term.Colour
}
