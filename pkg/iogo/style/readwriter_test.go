package style

import (
	"os"
	"testing"
)

func TestCreateDefaultReadWriter(t *testing.T) {
	CreateDefaultReadWriter(os.Stdin, os.Stdout)
}
