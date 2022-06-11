package style

import (
	"github.com/zarthus/iogo/v2/test"
	"os"
	"testing"
)

func TestCreateStdReadWriter(t *testing.T) {
	NewStdReadWriter()
}

func TestCreateReadWriter(t *testing.T) {
	NewReadWriter(os.Stdin, os.Stdout)
}

func TestReadwriter_InputStyle(t *testing.T) {
	NewStdReadWriter().InputStyle()
}

func TestReadwriter_OutputStyle(t *testing.T) {
	NewStdReadWriter().OutputStyle()
}

func TestReadwriter_Close(t *testing.T) {
	f1, f2 := test.NullFile(), test.NullFile()

	err := NewReadWriter(f1, f2).Close()
	if err != nil {
		t.Fatalf("expected nil error, got %s", err.Error())
	}
}
