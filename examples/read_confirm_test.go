package examples

import (
	"github.com/zarthus/iogo/v2/test"
	"os"
	"testing"
)

func TestReadConfirm(t *testing.T) {
	reset := test.StdReset{}
	reset.Init()
	defer reset.Reset()

	os.Stdin.WriteString("y")
	ReadConfirm()
}
