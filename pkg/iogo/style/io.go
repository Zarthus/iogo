package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader"
	"github.com/zarthus/iogo/v2/pkg/iogo/writer"
)

type defaultReadWriter struct {
	writer iogo.Writer
	reader iogo.Reader
	style  iogo.Style
}

func CreateDefaultReadWriter() iogo.ReadWriter {
	w, r := writer.NewDefaultWriter(), reader.NewInMemoryReader()
	return defaultReadWriter{
		writer: w,
		reader: r,
		style:  createDefaultStyle(w, r),
	}
}

func (rw defaultReadWriter) Reader() iogo.Reader {
	return rw.reader
}

func (rw defaultReadWriter) Writer() iogo.Writer {
	return rw.writer
}

func (rw defaultReadWriter) Style() iogo.Style {
	return rw.style
}
