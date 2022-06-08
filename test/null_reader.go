package test

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"io"
)

type nullReader struct {
	returns string

	readIndex int
}

func NewNullReader(ret string) *nullReader {
	return &nullReader{
		returns:   ret,
		readIndex: 0,
	}
}

func (r *nullReader) Read(p []byte) (n int, err error) {
	if r.readIndex >= len(r.returns) {
		err = io.EOF
		return
	}

	n = copy(p, []byte(r.returns)[r.readIndex:])
	r.readIndex += n
	return
}

func (r nullReader) Readln(_ iogo.Options) (string, error) {
	return r.returns, nil
}

func (r nullReader) Close() error {
	return nil
}

func (r nullReader) Reset() {
}
