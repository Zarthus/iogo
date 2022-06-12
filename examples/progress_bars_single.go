package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"time"
)

func ProgressBarsSingle() {
	io := style.NewStdReadWriter()
	bf1 := formatter.NewSimpleProgressBarFormatter("My Progress Bar")

	err := io.OutputStyle().Progress(350, func(bar iogo.ProgressBar) {
		bar.Advance(33)
		time.Sleep(time.Millisecond * 150)
	}, &bf1)

	if err != nil {
		panic(err)
	}
}
