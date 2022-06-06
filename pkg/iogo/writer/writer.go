package writer

import (
	"fmt"
)

type DefaultWriter struct{}

func NewDefaultWriter() *DefaultWriter {
	return &DefaultWriter{}
}

func (w DefaultWriter) Write(msg string) {
	fmt.Print(msg)
}

func (w DefaultWriter) Writeln(msg string) {
	fmt.Println(msg)
}
