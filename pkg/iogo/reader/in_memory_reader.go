package reader

import (
	"github.com/zarthus/iogo/v2/internal"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader/history"
	"os"
)

// Maintains an active list of history based in memory
// The moment the software closes, the history is gone
type inMemoryReader struct {
	file    *os.File
	history iogo.HistoryTracker
}

func NewInMemoryReader(file *os.File) *inMemoryReader {
	return &inMemoryReader{
		file,
		history.NewHistoryTracker([]string{}),
	}
}

func (r inMemoryReader) Read(p []byte) (n int, err error) {
	return r.file.Read(p)
}

func (r inMemoryReader) Readln(options iogo.Options) (string, error) {
	input, err := internal.Read(r.file, r.history)

	if err != nil {
		return *r.fallback(input, &options), err
	}

	fb := *r.fallback(input, &options)
	if input != nil {
		r.trackHistory(*input, &options)
	}
	return fb, err
}

func (r inMemoryReader) Reset() {
	r.history.Reset()
}

func (r inMemoryReader) Close() error {
	return r.file.Close()
}

func (r inMemoryReader) fallback(input *string, opts *iogo.Options) *string {
	if input == nil {
		return opts.Default
	}
	return input
}

func (r inMemoryReader) trackHistory(input string, opts *iogo.Options) {
	if opts != nil && !opts.DoNotTrack && input != "" {
		r.history.Track(input)
	}
}
