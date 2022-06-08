package term

import (
	"os"
	"strconv"
	"testing"
)

func TestDetect(t *testing.T) {
	initEnv("test-256color", 10, 10, t)
	tinfo := Detect()

	if tinfo.Env != "test-256color" {
		t.Fatalf("Expected env to be 'test-256color', was %s", tinfo.Env)
	}

	if tinfo.Size.Columns != 10 || tinfo.Size.Lines != 10 {
		t.Fatalf("Expected size to be (10, 10), was (%d, %d)", tinfo.Size.Columns, tinfo.Size.Lines)
	}

	if tinfo.Colours != ModernColourSupport {
		t.Fatalf("Expected 256 colour support, got %d", tinfo.Colours)
	}

	initEnv("test-256color", 12, 12, t)
	tinfo.UpdateSize()
	if tinfo.Size.Columns != 12 || tinfo.Size.Lines != 12 {
		t.Fatalf("Expected size to be (12, 12), was (%d, %d)", tinfo.Size.Columns, tinfo.Size.Lines)
	}
}

func initEnv(term string, columns int, lines int, t *testing.T) {
	if os.Setenv("TERM", term) != nil {
		t.Fatalf("Could not set env TERM")
	}
	if os.Setenv("COLUMNS", strconv.Itoa(columns)) != nil || os.Setenv("LINES", strconv.Itoa(lines)) != nil {
		t.Fatalf("Could not set env COLUMNS or LINES")
	}
}
