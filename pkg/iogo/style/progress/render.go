package progress

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
)

func Render(w iogo.Writer, bar iogo.ProgressBar, formatter iogo.ProgressBarFormatter) {
	formatted := formatter.Format(bar)

	w.Write(string(term.CarriageReturn) + formatted)
}
