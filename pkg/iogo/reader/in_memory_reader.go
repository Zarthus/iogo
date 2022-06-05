package reader

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader/history"
	"github.com/zarthus/iogo/v2/pkg/iogo/reader/raw"
)

// Maintains an active list of history based in memory
// The moment the software closes, the history is gone.
type inMemoryReader struct {
	history iogo.HistoryTracker
}

func NewInMemoryReader() *inMemoryReader {
	return &inMemoryReader{
		history.NewHistoryTracker([]string{}),
	}
}

func (reader inMemoryReader) ReadLine(options iogo.Options) (string, error) {
	input, err := raw.Read(reader.history)

	if err != nil {
		return reader.fallback(&input, options), err
	}

	reader.trackHistory(&input, options)
	return reader.fallback(&input, options), err
}

func (reader inMemoryReader) Reset() {
	reader.history.Reset()
}

func (reader inMemoryReader) fallback(input *string, options iogo.Options) string {
	if input != nil {
		return *input
	}
	return options.Default
}

func (reader inMemoryReader) trackHistory(input *string, options iogo.Options) {
	if options.DoNotTrack || input == nil || *input == "" {
		return
	}

	reader.history.Track(*input)
}
