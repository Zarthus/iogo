package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/stringtools"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
	"math"
	"strings"
)

type writerStyle struct {
	termcol string
	unicode bool
	width   uint
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
		colours: true,    // TODO: Detect colour support based on `termcol`, allow switching colours off
		reader:  reader,
		writer:  writer,
	}
}

func (style writerStyle) Title(message string) {
	style.writer.WriteLine("")
	style.Section(message)
	style.writer.WriteLine("")
}

func (style writerStyle) Section(message string) {
	line := strings.Repeat("=", len(message))

	style.writer.WriteLine(message + "\n" + line)
}

func (style writerStyle) Block(message string, options iogo.Options) {
	blocklen := math.Min(80, float64(style.width-(blockPadding*2)))
	space := strings.Repeat(" ", int(blocklen))
	padding := strings.Repeat(" ", blockPadding)
	msglen := len(message)
	needsWrapping := msglen > int(style.width)

	var prefix string
	var suffix string

	if &options.BgColor != nil {
		prefix = term.BackgroundColourize(options.BgColor, true)
		suffix = string(term.Reset)
	} else {
		prefix = ""
		suffix = ""
	}

	style.writer.WriteLine("")
	style.writer.WriteLine(prefix + padding + space + padding + suffix)
	if !needsWrapping {
		msgpadding := strings.Repeat(" ", int(style.width)-((blockPadding)+len(message)))
		style.writer.WriteLine(prefix + padding + message + msgpadding + suffix)
	} else {
		for _, msg := range stringtools.Wrap(message, style.width-blockPadding*2) {
			msgpadding := strings.Repeat(" ", int(style.width)-((blockPadding)+len(msg)))
			style.writer.WriteLine(prefix + padding + msg + msgpadding + suffix)
		}
	}
	style.writer.WriteLine(prefix + padding + space + padding + suffix)
	style.writer.WriteLine("")
}

func (style writerStyle) Info(message string) {
	if style.colours {
		style.Block(message, iogo.Options{BgColor: term.Cyan})
	} else {
		style.Block("INFO: "+message, iogo.Options{})
	}
}

func (style writerStyle) Success(message string) {
	if style.colours {
		style.Block(message, iogo.Options{BgColor: term.Green})
	} else {
		style.Block("SUCCESS: "+message, iogo.Options{})
	}
}

func (style writerStyle) Warning(message string) {
	if style.colours {
		style.Block(message, iogo.Options{BgColor: term.Yellow})
	} else {
		style.Block("WARNING: "+message, iogo.Options{})
	}
}

func (style writerStyle) Error(message string) {
	if style.colours {
		style.Block(message, iogo.Options{BgColor: term.Red})
	} else {
		style.Block("ERROR: "+message, iogo.Options{})
	}
}

func (style writerStyle) Progress(bar iogo.ProgressBar, runnable func(progressBar iogo.ProgressBar), barFormatter *iogo.ProgressBarFormatter) {
	var bf iogo.ProgressBarFormatter
	if barFormatter == nil {
		bf = style.autodetectProgressFormatter()
	} else {
		bf = *barFormatter
	}

	for !bar.IsFinished() {
		progress.Render(style.writer, bar, bf)
		runnable(bar)
	}
	progress.Render(style.writer, bar, bf)
}

func (style writerStyle) autodetectProgressFormatter() iogo.ProgressBarFormatter {
	return formatter.NewSimpleProgressBarFormatter(nil)
}
