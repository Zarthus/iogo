package examples

import (
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestWriteTextBoxWriteTextBox(t *testing.T) {
	reset := test.StdReset{}
	reset.Init()
	defer reset.Reset()

	WriteTextBox()
}
