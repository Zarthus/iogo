package formatter

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"strings"
)

type simpleProgressBarFormatter struct {
	descriptor *string
}

func NewSimpleProgressBarFormatter(descriptor *string) *simpleProgressBarFormatter {
	return &simpleProgressBarFormatter{
		descriptor: descriptor,
	}
}

func (formatter simpleProgressBarFormatter) Format(bar iogo.ProgressBar) string {
	cur, max := bar.Current(), bar.Maximum()
	bars := int(float32(cur)/float32(max)*100) / 10
	desc := ""
	if formatter.descriptor != nil {
		desc = " " + *formatter.descriptor
	}

	return fmt.Sprintf(
		"[%s%s]%s %d/%d",
		strings.Repeat("#", bars),
		strings.Repeat(" ", 10-bars),
		desc,
		bar.Current(),
		bar.Maximum(),
	)
}
