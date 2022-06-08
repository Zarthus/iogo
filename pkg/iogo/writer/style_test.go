package writer

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/term/col"
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestNewWriterStyle(t *testing.T) {
	NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
}

func TestWriterStyle_Title(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
	style.Title("foo")
}

func TestWriterStyle_Section(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
	style.Section("foo")
}

func TestWriterStyle_Block(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())

	c := col.Red
	style.Block("foo", iogo.Options{BgColour: &c})
}

func TestWriterStyle_Info(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
	style.Info("foo")
}

func TestWriterStyle_Success(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
	style.Success("foo")
}

func TestWriterStyle_Warning(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
	style.Warning("foo")
}

func TestWriterStyle_Error(t *testing.T) {
	style := NewWriterStyle(test.NewNullReader("foo"), test.NewNullWriter())
	style.Error("foo")
}
