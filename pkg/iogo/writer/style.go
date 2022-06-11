package writer

import (
	"github.com/zarthus/iogo/v2/internal"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	term2 "github.com/zarthus/iogo/v2/pkg/term"
	"github.com/zarthus/iogo/v2/pkg/term/col"
	"math"
	"strings"
	"time"
)

type writerStyle struct {
	info term2.TerminalInfo

	reader iogo.Reader
	writer iogo.Writer
}

const blockPadding = 2

func NewWriterStyle(r iogo.Reader, w iogo.Writer) iogo.WriterStyle {
	return &writerStyle{
		info:   term2.Detect(),
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
		prefix = term2.BackgroundColourize(*options.BgColour, true)
		suffix = string(term2.Reset)
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
		for _, m := range internal.Wrap(msg, cols-blockPadding*2) {
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

func (style *writerStyle) Progress(max uint, runnable func(bar iogo.ProgressBar), barFormatter *iogo.ProgressBarFormatter) error {
	var bf *iogo.ProgressBarFormatter
	if barFormatter == nil {
		bf = style.autodetectProgressFormatter()
	} else {
		bf = barFormatter
	}

	return style.ProgressBar([]iogo.ProgressBarContainer{{
		Bar:       progress.NewDefaultProgressBar(max),
		Runnable:  runnable,
		Formatter: bf,
	}})
}

func (style *writerStyle) ProgressBar(bars []iogo.ProgressBarContainer) (err error) {
	err = progress.RenderMultiple(style.writer, bars, true)
	if err != nil {
		return err
	}

	barLen := len(bars)

	for _, bar := range bars {
		go func(b iogo.ProgressBarContainer) {
			for !b.Bar.IsFinished() {
				b.Runnable(b.Bar)
			}
			barLen--
		}(bar)
	}

	for barLen > 0 {
		progress.RenderMultiple(style.writer, bars, false)
		time.Sleep(time.Millisecond * 10)
	}

	err = progress.RenderMultiple(style.writer, bars, false)
	return
}

func (style writerStyle) autodetectProgressFormatter() *iogo.ProgressBarFormatter {
	bf := formatter.NewSimpleProgressBarFormatter("")
	return &bf
}

func (style writerStyle) doBlock(msg string, col col.Colour, oftype string) {
	if style.info.Colours != term2.NoColourSupport {
		style.Block(msg, iogo.Options{BgColour: &col})
	} else {
		style.Block(oftype+": "+msg, iogo.Options{})
	}
}
