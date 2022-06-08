package test

import (
	"os"
)

func NullFile() *os.File {
	file, err := os.Open(os.DevNull)

	if err != nil {
		panic(err)
	}

	return file
}
