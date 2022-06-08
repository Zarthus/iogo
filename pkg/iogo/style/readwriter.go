package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader"
	"github.com/zarthus/iogo/v2/pkg/iogo/writer"
	"os"
)

type readwriter struct {
	reader      iogo.Reader
	readerStyle iogo.ReaderStyle

	writer      iogo.Writer
	writerStyle iogo.WriterStyle
}

// CreateStdReadWriter initializes an iogo.Iogo instance with os.Stdin and os.Stdout
func CreateStdReadWriter() iogo.Iogo {
	r := reader.NewInMemoryReader(os.Stdin)
	w := writer.NewDefaultWriter(os.Stdout)

	return readwriter{
		reader:      r,
		readerStyle: reader.NewReaderStyle(r, w),

		writer:      w,
		writerStyle: writer.NewWriterStyle(r, w),
	}
}

func CreateReadWriter(read *os.File, write *os.File) iogo.Iogo {
	r := reader.NewInMemoryReader(read)
	w := writer.NewDefaultWriter(write)

	return readwriter{
		reader:      r,
		readerStyle: reader.NewReaderStyle(r, w),

		writer:      w,
		writerStyle: writer.NewWriterStyle(r, w),
	}
}

func (io readwriter) Reader() iogo.Reader {
	return io.reader
}

func (io readwriter) Read(p []byte) (int, error) {
	return io.reader.Read(p)
}

func (io readwriter) Writer() iogo.Writer {
	return io.writer
}

func (io readwriter) Write(p []byte) (int, error) {
	return io.writer.Write(p)
}

func (io readwriter) InputStyle() iogo.ReaderStyle {
	return io.readerStyle
}

func (io readwriter) OutputStyle() iogo.WriterStyle {
	return io.writerStyle
}

func (io readwriter) Close() error {
	rerr, werr := io.Reader().Close(), io.Writer().Close()

	if rerr != nil {
		return rerr
	}
	if werr != nil {
		return werr
	}
	return nil
}
