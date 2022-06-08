package writer

import (
	"os"
)

type defaultWriter struct {
	file *os.File
}

func NewDefaultWriter(file *os.File) *defaultWriter {
	return &defaultWriter{
		file: file,
	}
}

func (writer defaultWriter) Write(msg string) (int, error) {
	return writer.file.WriteString(msg)
}

func (writer defaultWriter) Writeln(msg string) (int, error) {
	return writer.file.WriteString(msg + "\n")
}
