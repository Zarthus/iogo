package progress

import (
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"github.com/zarthus/iogo/v2/test"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	bar := NewDefaultProgressBar(10)
	bar.Advance(3)

	writer := test.NewNullWriter()
	Render(writer, bar, formatter.NewSimpleProgressBarFormatter(nil))

	if len(writer.Get()) != 1 {
		t.Fatalf("Expected writer to receive one message, got %d", len(writer.Get()))
	}

	msg := writer.Get()[0]
	expected := "[###       ]  3/10 ( 30.0%)"
	if !strings.Contains(msg, expected) {
		t.Fatalf("Expected message to look like a specific format, \nexpected: %s\n   given: %s", expected, msg)
	}
}
