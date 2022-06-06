package reader

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader/history"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader/raw"
)

// Maintains an active list of history based in memory.
// The moment the software closes, the history is gone.
type inMemoryReader struct {
	history iogo.HistoryTracker
}

func NewInMemoryReader() *inMemoryReader {
	return &inMemoryReader{
		history.NewHistoryTracker([]string{}),
	}
}

func (r inMemoryReader) ReadLine(opts *iogo.Options) (string, error) {
	input, err := raw.Read(r.history)
	if err == nil {
		r.trackHistory(input, opts)
	}
	return r.fallback(input, opts), err
}

func (r inMemoryReader) Reset() {
	r.history.Reset()
}

func (r inMemoryReader) fallback(input string, opts *iogo.Options) string {
	if input == "" {
		return opts.Default
	}
	return input
}

func (r inMemoryReader) trackHistory(input string, opts *iogo.Options) {
	if !opts.DoNotTrack && input != "" {
		r.history.Track(input)
	}
}
