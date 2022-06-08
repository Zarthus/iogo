package term

import (
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestCursorInstruction(t *testing.T) {
	mockWriter := test.NewNullWriter()
	instructor := CursorInstruction{writer: mockWriter}

	expectMessage := "Foo!"
	expectCount := 7

	instructor.SavePosition().Down(1).Up(1).Write(expectMessage).Forward(1).Backward(1).RestorePosition()

	if len(mockWriter.Get()) != expectCount {
		t.Fatalf("Count was %d, not %d", len(mockWriter.Get()), expectCount)
	}

	if mockWriter.Get()[3] != expectMessage {
		t.Fatalf("Fourth key in writer | expect: %s | given: %s", expectMessage, mockWriter.Get()[3])
	}
}
