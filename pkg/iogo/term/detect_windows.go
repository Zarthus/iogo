//go:build windows

package term

import (
	"strings"
)

func detectColours(env string) TerminalColourSupport {
	if strings.Contains(env, "256color") {
		return ModernColourSupport
	}
	return NoColourSupport
}
