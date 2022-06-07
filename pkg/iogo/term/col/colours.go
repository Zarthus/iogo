package col

type Colour uint

// xterm colours
const (
	Black   Colour = 30 // 0x000000
	Red     Colour = 31 // 0xff0000
	Green   Colour = 32 // 0x00ff00
	Yellow  Colour = 33 // 0xffff00
	Blue    Colour = 34 // 0x0000ff
	Fuchsia Colour = 35 // 0xff00ff, aka magenta
	Cyan    Colour = 36 // 0x00ffff, aka aqua
	White   Colour = 37 // 0xffffff
)
