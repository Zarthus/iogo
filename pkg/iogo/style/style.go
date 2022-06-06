package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	readr "github.com/zarthus/iogo/v2/pkg/iogo/reader/style"
	write "github.com/zarthus/iogo/v2/pkg/iogo/writer/style"
)

type defaultStyle struct {
	input  iogo.ReaderStyle
	output iogo.WriterStyle
}

func createDefaultStyle(writer iogo.Writer, reader iogo.Reader) iogo.Style {
	return defaultStyle{
		input:  readr.NewReaderStyle(writer, reader),
		output: write.NewWriterStyle(writer, reader),
	}
}

func (s defaultStyle) Input() iogo.ReaderStyle {
	return s.input
}

func (s defaultStyle) Output() iogo.WriterStyle {
	return s.output
}
