package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/stringtools"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"github.com/zarthus/iogo/v2/pkg/iogo/term"
	"github.com/zarthus/iogo/v2/pkg/iogo/term/col"
	"math"
	"strings"
)

type writerStyle struct {
	info term.TerminalInfo

	reader iogo.Reader
	writer iogo.Writer
}

const blockPadding = 2

func NewWriterStyle(r iogo.Reader, w iogo.Writer) iogo.WriterStyle {
	return &writerStyle{
		info:   term.Detect(),
		reader: r,
		writer: w,
	}
}

func (style writerStyle) Title(msg string) {
	style.writer.Writeln("")
	style.Section(msg)
	style.writer.Writeln("")
}

func (style writerStyle) Section(msg string) {
	line := strings.Repeat("=", len(msg))

	style.writer.Writeln(msg + "\n" + line)
}

func (style writerStyle) Block(msg string, options iogo.Options) {
	cols := style.info.Size.Columns
	blocklen := math.Min(float64(cols), float64(cols-(blockPadding*2)))
	space := strings.Repeat(" ", int(blocklen))
	padding := strings.Repeat(" ", blockPadding)
	msglen := len(msg)
	needsWrapping := msglen > int(cols)

	var prefix string
	var suffix string

	if options.BgColour != nil {
		prefix = term.BackgroundColourize(*options.BgColour, true)
		suffix = string(term.Reset)
	} else {
		prefix = ""
		suffix = ""
	}

	style.writer.Writeln("")
	style.writer.Writeln(prefix + padding + space + padding + suffix)
	if !needsWrapping {
		msgpadding := strings.Repeat(" ", int(cols)-((blockPadding)+len(msg)))
		style.writer.Writeln(prefix + padding + msg + msgpadding + suffix)
	} else {
		for _, m := range stringtools.Wrap(msg, cols-blockPadding*2) {
			msgpadding := strings.Repeat(" ", int(cols)-((blockPadding)+len(m)))
			style.writer.Writeln(prefix + padding + m + msgpadding + suffix)
		}
	}
	style.writer.Writeln(prefix + padding + space + padding + suffix)
	style.writer.Writeln("")
}

func (style writerStyle) Info(msg string) {
	style.doBlock(msg, col.Cyan, "INFO")
}

func (style writerStyle) Success(msg string) {
	style.doBlock(msg, col.Green, "SUCCESS")
}

func (style writerStyle) Warning(msg string) {
	style.doBlock(msg, col.Yellow, "WARNING")
}

func (style writerStyle) Error(msg string) {
	style.doBlock(msg, col.Red, "ERROR")
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
	style.writer.Write("\n")
}

func (style writerStyle) autodetectProgressFormatter() iogo.ProgressBarFormatter {
	return formatter.NewSimpleProgressBarFormatter(nil)
}

func (style writerStyle) doBlock(msg string, col col.Colour, oftype string) {
	if style.info.Colours != term.NoColourSupport {
		style.Block(msg, iogo.Options{BgColour: &col})
	} else {
		style.Block(oftype+": "+msg, iogo.Options{})
	}
}
