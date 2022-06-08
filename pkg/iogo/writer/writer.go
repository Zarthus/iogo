package writer

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
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

func (w *defaultWriter) Write(p []byte) (int, error) {
	return w.file.Write(p)
}

func (w *defaultWriter) Writeln(msg string) (int, error) {
	return w.WriteString(msg + iogo.OsLineEndings)
}

func (w *defaultWriter) WriteString(msg string) (int, error) {
	return w.file.WriteString(msg)
}

func (w *defaultWriter) Close() error {
	return w.file.Close()
}
