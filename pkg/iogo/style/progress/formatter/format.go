package formatter

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"math"
	"strconv"
	"strings"
)

type simpleProgressBarFormatter struct {
	descriptor string
}

// NewSimpleProgressBarFormatter empty strings are valid
func NewSimpleProgressBarFormatter(descriptor string) iogo.ProgressBarFormatter {
	return &simpleProgressBarFormatter{
		descriptor: descriptor,
	}
}

func (formatter simpleProgressBarFormatter) Format(bar iogo.ProgressBar) string {
	cur, max := bar.Current(), bar.Maximum()
	percentage := float32(cur) / float32(max) * 100
	bars := int(percentage / 10)
	// determines how many 'zeroes' are in the max number, so we can pad the starting number to always
	// be identical. This may 'glitch' formatting when bar.SetMaximum() is called, but oh well.
	printfLeading := strconv.Itoa(1 + int(math.Log10(float64(max))))

	return fmt.Sprintf(
		"[%-10s] %"+printfLeading+"d/%d (%5.1f%%) %s",
		strings.Repeat("#", bars),
		bar.Current(),
		bar.Maximum(),
		percentage,
		formatter.descriptor,
	)
}
