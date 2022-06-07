package term

import (
	"os"
	"strconv"
)

func Detect() TerminalInfo {
	env := os.Getenv("TERM")

	return TerminalInfo{
		Env:      env,
		Colours:  detectColours(env),
		Size:     detectTermSize(),
		Unicode:  detectUnicodeSupport(),
		Attached: true, // TODO: determine if attached to stdin
	}
}

func (info *TerminalInfo) UpdateSize() {
	info.Size = detectTermSize()
}

func detectTermSize() TerminalSize {
	return TerminalSize{
		Columns: envInt("COLUMNS", 80),
		Lines:   envInt("LINES", 40),
	}
}

func detectUnicodeSupport() bool {
	// TODO: impl
	return true
}

func envInt(envname string, fallback uint) uint {
	e := os.Getenv(envname)

	if e == "" {
		return fallback
	}

	if c, err := strconv.Atoi(e); err != nil {
		return fallback
	} else {
		return uint(c)
	}
}
