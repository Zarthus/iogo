package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader"
	"github.com/zarthus/iogo/v2/pkg/iogo/writer"
	"os"
)

type defaultIo struct {
	reader iogo.Reader
	writer iogo.Writer
	style  iogo.Style
}

func CreateDefaultReadWriter(read *os.File, write *os.File) iogo.Iogo {
	var r iogo.Reader
	var w iogo.Writer

	if &read == nil {
		r = reader.NewInMemoryReader(os.Stdin)
	} else {
		r = reader.NewInMemoryReader(read)
	}
	if &write == nil {
		w = writer.NewDefaultWriter(os.Stdout)
	} else {
		w = writer.NewDefaultWriter(write)
	}

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
