package writer

import (
	"os"
	"testing"
)

func TestNewDefaultWriter(t *testing.T) {
	NewDefaultWriter(os.Stdin)
}

func TestDefaultWriter_Close(t *testing.T) {
	err := NewDefaultWriter(&os.File{}).Close()

	if err == nil {
		t.Fatalf("Expected panic, got nil")
	}
}
