package progress

import "github.com/zarthus/iogo/v2/pkg/iogo"

func Render(w iogo.Writer, bar iogo.ProgressBar, formatter iogo.ProgressBarFormatter) {
	formatted := formatter.Format(bar)

	w.Writeln(formatted)
}
