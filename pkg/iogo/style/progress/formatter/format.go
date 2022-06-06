package formatter

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"strings"
)

type simpleProgressBarFormatter struct {
	descriptor *string
}

func NewSimpleProgressBarFormatter(descriptor *string) iogo.ProgressBarFormatter {
	return &simpleProgressBarFormatter{
		descriptor: descriptor,
	}
}

func (formatter simpleProgressBarFormatter) Format(bar iogo.ProgressBar) string {
	cur, max := bar.Current(), bar.Maximum()
	percentage := float32(cur) / float32(max) * 100
	bars := int(percentage / 10)

	desc := ""
	if formatter.descriptor != nil {
		desc = " " + *formatter.descriptor
	}

	return fmt.Sprintf(
		"[%s%s]%s %d/%d (%.2f%%)",
		strings.Repeat("#", bars),
		strings.Repeat(" ", 10-bars),
		desc,
		bar.Current(),
		bar.Maximum(),
		percentage,
	)
}
