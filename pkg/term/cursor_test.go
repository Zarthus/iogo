package term

import (
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestCursorInstruction(t *testing.T) {
	mockWriter := test.NewNullWriter()

	instructor := CursorInstruction{Writer: mockWriter}

	expectMessage := "Foo!"
	expectCount := 5

	instructor.Down(1).Up(1).Write(expectMessage).Forward(1).Backward(1)

	if len(mockWriter.Get()) != expectCount {
		t.Fatalf("Count was %d, not %d", len(mockWriter.Get()), expectCount)
	}

	key := 2
	if mockWriter.Get()[key] != expectMessage {
		t.Fatalf("%d key in writer | expect: %s | given: %s", key+1, expectMessage, mockWriter.Get()[key])
	}

	assertNoErr(instructor, t)
}

func assertNoErr(instructor CursorInstruction, t *testing.T) {
	if instructor.Err != nil {
		t.Fatalf("CursorInstruction has an error stored: %v", instructor.Err)
	}
}
