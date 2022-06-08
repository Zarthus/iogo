package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader"
	"github.com/zarthus/iogo/v2/pkg/iogo/writer"
)

type defaultIo struct {
	reader iogo.Reader
	writer iogo.Writer
	style  iogo.Style
}

func CreateDefaultReadWriter() iogo.Iogo {
	w, r := writer.NewDefaultWriter(), reader.NewInMemoryReader()
	return defaultIo{
		reader: r,
		writer: w,
		style:  createDefaultStyle(r, w),
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
