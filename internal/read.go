package internal

import (
	"bufio"
	"io"
)

func Read(rd io.Reader) (*string, error) {
	// TODO: In the future, if TTY, we may want to support
	// reading from key up and key down and deal with history.
	r := bufio.NewReader(rd)

	inp, _, err := r.ReadLine()
	s := string(inp)

	if err != nil {
		if err == io.EOF {
			return &s, nil
		}
		return &s, err
	}

	return &s, nil
}
