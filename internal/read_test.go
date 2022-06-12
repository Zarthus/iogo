package internal

import (
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestRead(t *testing.T) {
	nr := test.NewNullReader("foo")
	read, err := Read(nr)

	if err != nil {
		t.Fatalf("Did not expect err, got %v", err)
	}

	if read == nil {
		t.Fatalf("received nil read")
	}

	if *read != "foo" {
		t.Fatalf("Expected input to equal 'foo', got '%s'", *read)
	}
}
