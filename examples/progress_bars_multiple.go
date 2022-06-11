package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"time"
)

func ProgressBarsMultiple() {
	io := style.NewStdReadWriter()
	bf1, bf2, bf3 := formatter.NewSimpleProgressBarFormatter("Bar 1"),
		formatter.NewSimpleProgressBarFormatter("Bar 2"),
		formatter.NewSimpleProgressBarFormatter("Bar 3")

	io.OutputStyle().ProgressBar([]iogo.ProgressBarContainer{
		{
			Bar: progress.NewDefaultProgressBar(100),
			Runnable: func(bar iogo.ProgressBar) {
				bar.Advance(1)
				time.Sleep(time.Millisecond * 33)
			},
			Formatter: &bf1,
		},
		{
			Bar: progress.NewDefaultProgressBar(125),
			Runnable: func(bar iogo.ProgressBar) {
				bar.Advance(2)
				time.Sleep(time.Millisecond * 50)
			},
			Formatter: &bf2,
		},
		{
			Bar: progress.NewDefaultProgressBar(10),
			Runnable: func(bar iogo.ProgressBar) {
				bar.Advance(1)
				time.Sleep(time.Millisecond * 250)
			},
			Formatter: &bf3,
		},
	})
}
