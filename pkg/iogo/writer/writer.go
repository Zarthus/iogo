package writer

import (
	"fmt"
)

type defaultWriter struct {
}

func NewDefaultWriter() *defaultWriter {
	return &defaultWriter{}
}

func (writer defaultWriter) Write(message string) {
	fmt.Print(message)
}

func (writer defaultWriter) WriteLine(message string) {
	fmt.Println(message)
}
