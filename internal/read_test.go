package internal

import (
	"github.com/zarthus/iogo/v2/pkg/iogo/reader/history"
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestRead(t *testing.T) {
	nr := test.NewNullReader("foo")

	read, err := Read(nr, history.NewHistoryTracker([]string{}))
	if err != nil {
		t.Fatalf("Did not expect err, got %s", err.Error())
	}

	if read == nil {
		t.Fatalf("received nil read")
	}

	if *read != "foo" {
		t.Fatalf("Expected input to equal 'foo', got '%s'", *read)
	}
}
