package test

import "os"

type StdReset struct {
	oldIn  *os.File
	oldOut *os.File
}

func (reset *StdReset) Init() (err error) {
	r, w, err := os.Pipe()

	reset.oldIn = os.Stdin
	reset.oldOut = os.Stdout

	os.Stdin = r
	os.Stdout = w

	return
}

func (reset StdReset) Reset() {
	reset.oldIn = os.Stdin
	reset.oldOut = os.Stdout

	os.Stdin = reset.oldIn
	os.Stdout = reset.oldOut
}
