package writer

import (
	"github.com/zarthus/iogo/v2/internal"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress"
	"github.com/zarthus/iogo/v2/pkg/iogo/style/progress/formatter"
	"github.com/zarthus/iogo/v2/pkg/term"
	"github.com/zarthus/iogo/v2/pkg/term/col"
	"math"
	"strings"
	"time"
)

type writerStyle struct {
	info term.TerminalInfo

	reader  iogo.Reader
	writer  iogo.Writer
	wrapper iogo.StringWrapper
}

const blockPadding = 2

func NewWriterStyle(r iogo.Reader, w iogo.Writer) iogo.WriterStyle {
	return &writerStyle{
		info:    term.Detect(),
		reader:  r,
		writer:  w,
		wrapper: &internal.SimpleStringWrapper{},
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

func (style writerStyle) TextBox(msg string, options iogo.Options) error {
	tb := textbox{
		Info: style.info,
	}

	box := tb.Format(msg, nil, style.wrapper)
	if options.BgColour != nil || options.FgColour != nil {
		split := strings.Split(box, "\n")
		if options.BgColour != nil {
			split = style.rewriteLines(split, term.BackgroundColourize(*options.BgColour, false), string(term.Reset))
		}
		if options.FgColour != nil {
			split = style.rewriteLines(split, term.Colourize(*options.FgColour, false), string(term.Reset))
		}

		box = strings.Join(split, "\n")
	}

	_, err := style.writer.Writeln(box)
	return err
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
		for _, m := range style.wrapper.Wrap(msg, cols-blockPadding*2) {
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
	if style.info.Colours != term.NoColourSupport {
		style.Block(msg, iogo.Options{BgColour: &col})
	} else {
		style.Block(oftype+": "+msg, iogo.Options{})
	}
}

func (style writerStyle) rewriteLines(msg []string, prefix string, suffix string) []string {
	var message []string

	for _, m := range msg {
		message = append(message, prefix+m+suffix)
	}

	return message
}
