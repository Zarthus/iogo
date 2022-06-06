package raw

import (
	"bufio"
	"bytes"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"os"
)

func Read(tracker iogo.HistoryTracker) (string, error) {
	r := bufio.NewReader(os.Stdin)
	keyUp, keyDown := byte(iogo.KeyUp), byte(iogo.KeyDown)

	var buf bytes.Buffer
	for {
		if b, err := r.ReadByte(); err != nil {
			// if err != io.EOF {
			// 	return "", err
			// }
			break
		} else if b == keyUp {
			// TODO: modify input to history next
		} else if b == keyDown {
			// TODO: modify input to history previous
		} else if b == '\n' {
			break
		} else {
			buf.WriteByte(b)
		}
	}
	return buf.String(), nil
}
