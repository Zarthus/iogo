package test

import "github.com/zarthus/iogo/v2/pkg/iogo"

type nullReader struct {
	returns string
}

func NewNullReader(ret string) *nullReader {
	return &nullReader{
		returns: ret,
	}
}

func (reader nullReader) Readln(_ iogo.Options) (string, error) {
	return reader.returns, nil
}

func (reader nullReader) Reset() {
}
