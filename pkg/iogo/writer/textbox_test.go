package writer

import (
	"github.com/zarthus/iogo/v2/internal"
	"github.com/zarthus/iogo/v2/pkg/term"
	"strings"
	"testing"
)

func TestTextbox_Format(t *testing.T) {
	tb := textbox{Info: term.TerminalInfo{Unicode: false}}
	doTest(tb.Format(createGiven(), nil, &internal.SimpleStringWrapper{}), t)
}

func TestTextbox_FormatMultipleLines(t *testing.T) {
	tb := textbox{Info: term.TerminalInfo{Unicode: false}}
	doTest(tb.FormatMultipleLines([]string{createGiven()}, nil), t)
}

func TestTextbox_Draw(t *testing.T) {
	tb := textbox{Info: term.TerminalInfo{Unicode: false}}
	doTest(tb.Draw([]string{createGiven()}, len(createGiven()), tbfSimple()), t)
}

func createGiven() string {
	return "foo"
}

func createExpected() [3]string {
	var expected [3]string
	expected[0] = "+-----+"
	expected[1] = "| foo |"
	expected[2] = "+_____+"
	return expected
}

func doTest(givenunfmted string, t *testing.T) {
	given := strings.Split(givenunfmted, "\n")
	expected := createExpected()

	for i := 0; i < len(expected); i++ {
		if given[i] != expected[i] {
			t.Fatalf("Expected format [%d]: %s, got %s", i, expected[i], given[i])
		}
	}
}
