package progress

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
)

func Render(w iogo.Writer, bar iogo.ProgressBar, formatter iogo.ProgressBarFormatter) error {
	formatted := formatter.Format(bar)

	_, err := w.WriteString(string(term.CarriageReturn) + formatted)
	return err
}
