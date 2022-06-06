package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/stringtools"
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
	"strings"
)

type writerStyle struct {
	termcol string
	unicode bool
	width   int
	colours bool

	reader iogo.Reader
	writer iogo.Writer
}

const blockPadding = 2

func NewWriterStyle(writer iogo.Writer, reader iogo.Reader) iogo.WriterStyle {
	return &writerStyle{
		termcol: "xterm", // TODO: Detect terminal, use 256 colour radius if possible
		unicode: false,   // TODO: Detect terminal unicode support
		width:   80,      // TODO: Detect terminal window width
		colours: true,    // TODO: Detect colour supprot based on `termcol`, allow switching colours off
		reader:  reader,
		writer:  writer,
	}
}

func (s writerStyle) Title(msg string) {
	s.writer.Writeln("")
	s.Section(msg)
	s.writer.Writeln("")
}

func (s writerStyle) Section(msg string) {
	line := strings.Repeat("=", len(msg))
	s.writer.Writeln(msg + "\n" + line)
}

func (s writerStyle) Block(msg string, opts *iogo.Options) {
	blockLen := min(80, s.width-(blockPadding*2))
	space := strings.Repeat(" ", blockLen)
	padding := strings.Repeat(" ", blockPadding)
	needsWrapping := len(msg) > s.width

	var prefix, suffix string
	if opts.BgColour != 0 {
		prefix = term.BackgroundColourize(opts.BgColour, true)
		suffix = string(term.Reset)
	}

	s.writer.Writeln("")
	s.writer.Writeln(prefix + padding + space + padding + suffix)
	if needsWrapping {
		for _, msg := range stringtools.Wrap(msg, s.width-blockPadding*2) {
			msgPadding := strings.Repeat(" ", s.width-(blockPadding+len(msg)))
			s.writer.Writeln(prefix + padding + msg + msgPadding + suffix)
		}
	} else {
		msgPadding := strings.Repeat(" ", s.width-(blockPadding+len(msg)))
		s.writer.Writeln(prefix + padding + msg + msgPadding + suffix)
	}
	s.writer.Writeln(prefix + padding + space + padding + suffix)
	s.writer.Writeln("")
}

func (s writerStyle) Info(msg string) {
	if s.colours {
		s.Block(msg, &iogo.Options{BgColour: term.Cyan})
	} else {
		s.Block("INFO: "+msg, &iogo.Options{})
	}
}

func (s writerStyle) Success(msg string) {
	if s.colours {
		s.Block(msg, &iogo.Options{BgColour: term.Green})
	} else {
		s.Block("SUCCESS: "+msg, &iogo.Options{})
	}
}

func (s writerStyle) Warning(msg string) {
	if s.colours {
		s.Block(msg, &iogo.Options{BgColour: term.Yellow})
	} else {
		s.Block("WARNING: "+msg, &iogo.Options{})
	}
}

func (s writerStyle) Error(msg string) {
	if s.colours {
		s.Block(msg, &iogo.Options{BgColour: term.Red})
	} else {
		s.Block("ERROR: "+msg, &iogo.Options{})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
