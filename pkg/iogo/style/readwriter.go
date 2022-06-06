package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader"
	"github.com/zarthus/iogo/v2/pkg/iogo/writer"
)

type defaultIo struct {
	writer iogo.Writer
	reader iogo.Reader
	style  iogo.Style
}

func CreateDefaultReadWriter() iogo.Iogo {
	w, r := writer.NewDefaultWriter(), reader.NewInMemoryReader()
	return defaultIo{
		writer: w,
		reader: r,
		style:  createDefaultStyle(w, r),
	}
}

func (io defaultIo) Reader() iogo.Reader {
	return io.reader
}

func (io defaultIo) Writer() iogo.Writer {
	return io.writer
}

func (io defaultIo) Style() iogo.Style {
	return io.style
}
