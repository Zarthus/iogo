package reader

import (
	"github.com/zarthus/iogo/v2/internal"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"os"
)

type inMemoryReader struct {
	file *os.File
}

func NewInMemoryReader(file *os.File) *inMemoryReader {
	return &inMemoryReader{
		file,
	}
}

func (r inMemoryReader) Read(p []byte) (n int, err error) {
	return r.file.Read(p)
}

func (r inMemoryReader) Readln(options iogo.Options) (string, error) {
	input, err := internal.Read(r.file)

	if err != nil {
		return *r.fallback(input, &options), err
	}

	return *r.fallback(input, &options), err
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
