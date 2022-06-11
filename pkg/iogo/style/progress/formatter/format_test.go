package formatter

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	var format string
	current, maximum := uint(0), uint(10)

	bar := progress.NewDefaultProgressBar(maximum)
	formatter := NewSimpleProgressBarFormatter("Fooing")

	for {
		format = formatter.Format(bar)

		if !strings.Contains(format, fmt.Sprintf("%d/%d", current, maximum)) {
			t.Fatalf("Expected format output to contain '%d/%d', instead we got: '%s'", current, maximum, format)
		}

		if bar.IsFinished() {
			break
		}

		bar.Advance(2)
		current += 2
	}

	bar.Advance(50)
	if !strings.Contains(format, fmt.Sprintf("%d/%d", maximum, maximum)) {
		t.Fatalf("Expected format output to contain '%d/%d', instead we got: '%s'", maximum, maximum, format)
	}
}
