package progress

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"github.com/zarthus/iogo/v2/test"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	bar := NewDefaultProgressBar(10)
	bar.Advance(3)

	writer := test.NewNullWriter()
	err := Render(writer, bar, formatter.NewSimpleProgressBarFormatter(""))
	if err != nil {
		t.Fatalf("unexpected err, %v", err)
	}

	if len(writer.Get()) != 1 {
		t.Fatalf("Expected writer to receive one message, got %d", len(writer.Get()))
	}

	msg := writer.Get()[0]
	expected := "[###       ]  3/10 ( 30.0%)"
	if !strings.Contains(msg, expected) {
		t.Fatalf("Expected message to look like a specific format, \nexpected: %s\n   given: %s", expected, msg)
	}
}

func TestRenderMultipleWithFirst(t *testing.T) {
	bar := NewDefaultProgressBar(10)
	bar.Advance(3)
	bar2 := NewDefaultProgressBar(50)
	bar2.Advance(12)
	fmter := formatter.NewSimpleProgressBarFormatter("")

	writer := test.NewNullWriter()
	err := RenderMultiple(writer, []iogo.ProgressBarContainer{
		{
			Bar:       bar,
			Runnable:  nil,
			Formatter: &fmter,
		},
		{
			Bar:       bar2,
			Runnable:  nil,
			Formatter: &fmter,
		},
	}, true)

	if err != nil {
		t.Fatalf("unexpected err, %v", err)
	}

	if len(writer.Get()) != 3 {
		t.Fatalf("Expected writer to receive 3 messages, got %d", len(writer.Get()))
	}

	msg := writer.Get()[1]
	expected := "[###       ]  3/10 ( 30.0%)"
	if !strings.Contains(msg, expected) {
		t.Fatalf("Expected message to look like a specific format, \nexpected: %s\n   given: %s", expected, msg)
	}

	msg2 := writer.Get()[2]
	expected2 := "[##        ] 12/50 ( 24.0%)"
	if !strings.Contains(msg2, expected2) {
		t.Fatalf("Expected message to look like a specific format, \nexpected: %s\n   given: %s", expected2, msg2)
	}
}

func TestRenderMultipleWithoutFirst(t *testing.T) {
	bar := NewDefaultProgressBar(10)
	fmter := formatter.NewSimpleProgressBarFormatter("")
	writer := test.NewNullWriter()

	err := RenderMultiple(writer, []iogo.ProgressBarContainer{
		{
			Bar:       bar,
			Runnable:  nil,
			Formatter: &fmter,
		},
	}, false)

	if err != nil {
		t.Fatalf("unexpected err, %v", err)
	}

	if len(writer.Get()) != 3 {
		t.Fatalf("Expected writer to receive 3 messages, got %d", len(writer.Get()))
	}

	msg := writer.Get()[2]
	expected := "[          ]  0/10 (  0.0%)"
	if !strings.Contains(msg, expected) {
		t.Fatalf("Expected message to look like a specific format, \nexpected: %s\n   given: %s", expected, msg)
	}
}
