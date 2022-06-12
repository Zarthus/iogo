package internal

import (
	"bufio"
	"bytes"
	"io"
)

func Read(rd io.Reader) (*string, error) {
	// TODO: In the future, if TTY, we may want to support
	// reading from key up and key down and deal with history.
	r := bufio.NewReader(rd)

	var buf bytes.Buffer
	for {
		inp, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buf.WriteString(inp)
	}

	s := buf.String()
	return &s, nil
}
