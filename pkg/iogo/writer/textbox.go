package writer

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/term"
	"strings"
)

type textbox struct {
	Info term.TerminalInfo
}

func (tb *textbox) Format(msg string, tbf *textboxfmt, wrapper iogo.StringWrapper) string {
	return tb.FormatMultipleLines(wrapper.Wrap(msg, tb.Info.Size.Columns-4), tbf)
}

func (tb *textbox) FormatMultipleLines(msg []string, textboxformat *textboxfmt) string {
	var tbf textboxfmt
	if textboxformat == nil {
		tbf = tb.detectFormat()
	} else {
		tbf = *textboxformat
	}

	maxLineLen := 0
	for _, line := range msg {
		l := len(line)
		if l > maxLineLen {
			maxLineLen = l
		}
	}

	return tb.Draw(msg, maxLineLen, tbf)
}

func (tb *textbox) Draw(msg []string, maxlen int, tbf textboxfmt) string {
	header := tbf.CornerTopLeft + strings.Repeat(tbf.WallHeader, maxlen+2) + tbf.CornerTopRight
	footer := tbf.CornerBottomLeft + strings.Repeat(tbf.WallFooter, maxlen+2) + tbf.CornerBottomRight
	tpl := tbf.WallVertical + " %s%s " + tbf.WallVertical
	var drawing []string
	drawing = append(drawing, header)
	for _, line := range msg {
		padding := maxlen - len(line)

		if padding == 0 {
			drawing = append(drawing, fmt.Sprintf(tpl, line, ""))
		} else {
			drawing = append(drawing, fmt.Sprintf(tpl, line, strings.Repeat(" ", padding)))
		}
	}
	drawing = append(drawing, footer)
	return strings.Join(drawing, "\n")
}

func (tb *textbox) detectFormat() textboxfmt {
	if tb.Info.Unicode {
		return tbfUnicodeNonSolid()
	}

	return tbfSimple()
}
