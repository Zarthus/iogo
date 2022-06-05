package raw

import (
	"bufio"
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"os"
)

func Read(tracker iogo.HistoryTracker) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	keyUp, keyDown := byte(iogo.KeyUp), byte(iogo.KeyDown)

	var bytes []byte
	for {
		b, err := reader.ReadByte()

		if err != nil {
			//if err == io.EOF {
			//	break
			//}
			break
		} else if b == keyUp {
			// TODO: modify input to history next
		} else if b == keyDown {
			// TODO: modify input to history previous
		} else if b == '\n' {
			break
		} else {
			bytes = append(bytes, b)
		}
	}

	return string(bytes), nil
}
