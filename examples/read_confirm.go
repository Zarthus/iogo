package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
)

func ReadConfirm() {

	rw := style.NewStdReadWriter()

	def := "y"
	confirmed, err := rw.InputStyle().Confirm("Do you want to go swimming today?", iogo.Options{Default: &def})
	if err != nil {
		panic(err)
	}

	if confirmed {
		// go swimming
		rw.Writer().Writeln("You went swimming!")
	} else {
		rw.Writer().Writeln("You did not go swimming")
	}
}
