package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	reader "github.com/zarthus/iogo/v2/pkg/iogo/reader/style"
	writer "github.com/zarthus/iogo/v2/pkg/iogo/writer/style"
)

type defaultStyle struct {
	input  iogo.ReaderStyle
	output iogo.WriterStyle
}

func createDefaultStyle(r iogo.Reader, w iogo.Writer) iogo.Style {
	return defaultStyle{
		input:  reader.NewReaderStyle(r, w),
		output: writer.NewWriterStyle(r, w),
	}
}

func (style defaultStyle) Input() iogo.ReaderStyle {
	return style.input
}

func (style defaultStyle) Output() iogo.WriterStyle {
	return style.output
}
