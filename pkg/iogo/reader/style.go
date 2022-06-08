package reader

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"regexp"
	"strings"
)

// TODO: find a good way to handle the multitude of rare errors if cannot write but can receive
// in most* cases this should go fine, but e.g. writing to a closed file is a realistic scenario
// and may lead to bad UX.

type readerStyle struct {
	writer iogo.Writer
	reader iogo.Reader

	confirmRegexp *regexp.Regexp
}

func NewReaderStyle(reader iogo.Reader, writer iogo.Writer) iogo.ReaderStyle {
	return &readerStyle{
		writer:        writer,
		reader:        reader,
		confirmRegexp: regexp.MustCompile("^[yY]"),
	}
}

func (style readerStyle) Prompt(prompt string, options iogo.Options) (string, error) {
	style.writer.Writeln(prompt)

	return style.reader.Readln(options)
}

func (style readerStyle) Confirm(prompt string, options iogo.Options) (bool, error) {
	defaultYes := options.Default == nil || style.confirmRegexp.MatchString(*options.Default)

	var yes string
	var no string

	if defaultYes {
		yes = "Y"
		no = "n"
	} else {
		yes = "y"
		no = "N"
	}

	style.writer.Writeln(prompt + fmt.Sprintf(" (%s/%s)", yes, no))
	result, err := style.reader.Readln(options)

	if result == "" {
		return defaultYes, err
	}

	return style.confirmRegexp.MatchString(result), err
}

func (style readerStyle) Select(prompt string, valid []string, options iogo.Options) (string, error) {
	var safeValid []string
	for _, value := range valid {
		safeValid = append(safeValid, regexp.QuoteMeta(value))
	}
	selectRegexp := regexp.MustCompile("^" + strings.Join(safeValid, "|") + "$")

	style.writer.Writeln(prompt)
	style.writer.Writeln("Valid options: " + strings.Join(valid, ", "))

	for {
		result, err := style.reader.Readln(options)
		if err != nil {
			continue
		}

		if selectRegexp.MatchString(result) {
			return result, err
		} else {
			style.writer.Writeln("Your input did not match the valid selection.")
			style.writer.Writeln(prompt)
			style.writer.Writeln("Valid options: " + strings.Join(valid, ", "))
		}
	}
}
