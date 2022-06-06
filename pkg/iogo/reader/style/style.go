package style

import (
	"fmt"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"regexp"
	"strings"
)

type readerStyle struct {
	writer iogo.Writer
	reader iogo.Reader

	confirmRegexp *regexp.Regexp
}

func NewReaderStyle(w iogo.Writer, r iogo.Reader) iogo.ReaderStyle {
	return &readerStyle{
		writer:        w,
		reader:        r,
		confirmRegexp: regexp.MustCompile("^[yY]"),
	}
}

func (s readerStyle) Prompt(prompt string, opts *iogo.Options) (string, error) {
	s.writer.Writeln(prompt)
	return s.reader.ReadLine(opts)
}

func (s readerStyle) RequirePrompt(prompt string, opts *iogo.Options) (string, error) {
	s.writer.Writeln(prompt)
	return s.reader.ReadLine(opts)
}

func (s readerStyle) Confirm(prompt string, opts *iogo.Options) (bool, error) {
	defaultYes := &opts.Default == nil || opts.Default == "" || s.confirmRegexp.MatchString(opts.Default)

	var yes, no string
	if defaultYes {
		yes, no = "Y", "n"
	} else {
		yes, no = "y", "N"
	}

	s.writer.Writeln(fmt.Sprintf("%s (%s/%s)", prompt, yes, no))
	if result, err := s.reader.ReadLine(opts); err != nil {
		return false, err
	} else if result == "" {
		return defaultYes, err
	} else {
		return s.confirmRegexp.MatchString(result), nil
	}
}

func (s readerStyle) Select(prompt string, valid []string, opts *iogo.Options) (string, error) {
	var safeValid []string
	for _, value := range valid {
		safeValid = append(safeValid, value)
	}
	selectRegexp := regexp.MustCompile("^" + strings.Join(safeValid, "|") + "$")

	s.writer.Writeln(prompt)
	s.writer.Writeln("Valid options: " + strings.Join(valid, ", "))

	for {
		if result, err := s.reader.ReadLine(opts); err != nil {
			continue
		} else if selectRegexp.MatchString(result) {
			return result, nil
		} else {
			s.writer.Writeln("Your input did not match the valid selection.")
			s.writer.Writeln(prompt)
			s.writer.Writeln("Valid options: " + strings.Join(valid, ", "))
		}
	}
}
