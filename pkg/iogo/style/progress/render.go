package progress

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	term2 "github.com/zarthus/iogo/v2/pkg/term"
)

func Render(w iogo.Writer, bar iogo.ProgressBar, formatter iogo.ProgressBarFormatter) error {
	formatted := formatter.Format(bar)

	_, err := w.WriteString(string(term2.CarriageReturn) + formatted)
	return err
}

func RenderMultiple(w iogo.Writer, bars []iogo.ProgressBarContainer, first bool) (err error) {
	barLen := len(bars)

	if !first {
		cin := term2.CursorInstruction{Writer: w}
		cin.Up(barLen)
	}

	_, err = w.WriteString(string(term2.CarriageReturn))
	if err != nil {
		return
	}

	for _, bar := range bars {
		formatted := (*bar.Formatter).Format(bar.Bar)

		_, err = w.WriteString(formatted + "\n")

		if err != nil {
			return
		}
	}
	return
}
