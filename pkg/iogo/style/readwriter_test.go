package style

import (
	"github.com/zarthus/iogo/v2/test"
	"os"
	"testing"
)

func TestCreateStdReadWriter(t *testing.T) {
	CreateStdReadWriter()
}

func TestCreateReadWriter(t *testing.T) {
	CreateReadWriter(os.Stdin, os.Stdout)
}

func TestReadwriter_InputStyle(t *testing.T) {
	CreateStdReadWriter().InputStyle()
}

func TestReadwriter_OutputStyle(t *testing.T) {
	CreateStdReadWriter().OutputStyle()
}

func TestReadwriter_Close(t *testing.T) {
	f1, f2 := test.NullFile(), test.NullFile()

	err := CreateReadWriter(f1, f2).Close()
	if err != nil {
		t.Fatalf("expected nil error, got %s", err.Error())
	}
}
