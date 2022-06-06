package writer

import (
	"fmt"
)

type defaultWriter struct {
}

func NewDefaultWriter() *defaultWriter {
	return &defaultWriter{}
}

func (writer defaultWriter) Write(msg string) {
	fmt.Print(msg)
}

func (writer defaultWriter) Writeln(msg string) {
	fmt.Println(msg)
}
