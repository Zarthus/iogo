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

func createDefaultStyle(w iogo.Writer, r iogo.Reader) iogo.Style {
	return defaultStyle{
		input:  reader.NewReaderStyle(w, r),
		output: writer.NewWriterStyle(w, r),
	}
}

func (style defaultStyle) Input() iogo.ReaderStyle {
	return style.input
}

func (style defaultStyle) Output() iogo.WriterStyle {
	return style.output
}
