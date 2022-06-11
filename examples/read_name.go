package examples

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style"
)

func ReadName() {
	rw := style.NewStdReadWriter()

	rw.Writer().Writeln("What's your name?") // Or use `name, err := rw.InputStyle().Prompt(...)`
	if name, err := rw.Reader().Readln(iogo.Options{}); err != nil {
		panic(err)
	} else {
		rw.Writer().Writeln("Nice to meet you, " + name)
	}
}
