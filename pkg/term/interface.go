package term

type TerminalColourSupport int8

const (
	NoColourSupport     TerminalColourSupport = 0 // no colour support or opted out of colours
	ColourSupport       TerminalColourSupport = 1 // 16 colour
	ModernColourSupport TerminalColourSupport = 2 // 256 colour
)

type TerminalInfo struct {
	Env      string
	Colours  TerminalColourSupport
	Size     TerminalSize
	Unicode  bool
	Attached bool
}

type TerminalSize struct {
	Columns uint
	Lines   uint
}
