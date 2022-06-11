package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
)

func WriteStyle() {
	rw := style.NewStdReadWriter()

	rw.OutputStyle().Title("Welcome to iogo!")
	rw.OutputStyle().Success("You have installed the software correctly!")
}
