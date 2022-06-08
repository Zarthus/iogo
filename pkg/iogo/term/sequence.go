package term

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo/term/col"
)

type EscapeSequence rune
type ControlSequence string

// Escape Sequences
const (
	Bell           EscapeSequence = 0x7
	Backspace      EscapeSequence = 0x8
	Tab            EscapeSequence = 0x9
	LineFeed       EscapeSequence = 0x0A
	FormFeed       EscapeSequence = 0x0C
	CarriageReturn EscapeSequence = 0x0D
	Escape         EscapeSequence = 0x1B
	Control        EscapeSequence = 0x21
)

// note: terminals vary wildely, not all terminals may interpret the same control sequence as to mean the same thing.
// Resources:
// - https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
// - https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
const (
	Reset         ControlSequence = "\033[0m"
	Bold          ControlSequence = "\033[1m"
	Italics       ControlSequence = "\033[3m"
	Underline     ControlSequence = "\033[4m"
	Alert         ControlSequence = "\033[6m" // Some sort of blink in some terminals
	Invert        ControlSequence = "\033[7m" // aka "Reversed", "Reverse Video", "Negative"
	Concealed     ControlSequence = "\033[8m"
	Strikethrough ControlSequence = "\033[9m"
	colour        ControlSequence = "\033[%dm"
	colourLight   ControlSequence = "\033[%dm"
	bgColour      ControlSequence = "\033[%dm"
	bgColourLight ControlSequence = "\033[%dm"
	Redraw        ControlSequence = "\033c"
)

func Colourize(c col.Colour, bright bool) string {
	if bright {
		return fmt.Sprintf(string(colourLight), 60+c)
	} else {
		return fmt.Sprintf(string(colour), c)
	}
}

func BackgroundColourize(c col.Colour, bright bool) string {
	c += 10

	if bright {
		return fmt.Sprintf(string(bgColourLight), 60+c)
	} else {
		return fmt.Sprintf(string(bgColour), int(c))
	}
}

func WrapColour(colour col.Colour, s string, bright bool) string {
	return Colourize(colour, bright) + s + string(Reset)
}

func WrapBackgroundColour(colour col.Colour, s string, bright bool) string {
	return BackgroundColourize(colour, bright) + s + string(Reset)
}
