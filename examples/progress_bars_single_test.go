package examples

import (
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestProgressBarsSingle(t *testing.T) {
	reset := test.StdReset{}
	reset.Init()
	defer reset.Reset()

	ProgressBarsSingle()
}
