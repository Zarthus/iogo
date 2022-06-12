package reader

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestInMemoryReader_Read(t *testing.T) {
	r := NewInMemoryReader(test.NullFile())

	var b []byte
	_, err := r.Read(b)
	if err != nil {
		t.Fatalf("got err: %s", err.Error())
	}
}

func TestInMemoryReader_Readln(t *testing.T) {
	r := NewInMemoryReader(test.NullFile())

	if _, err := r.Readln(iogo.Options{}); err != nil {
		t.Fatalf("paniced: %v", err)
	}
}

func TestInMemoryReader_Close(t *testing.T) {
	err := NewInMemoryReader(test.NullFile()).Close()
	if err != nil {
		t.Fatalf("got err: %v", err)
	}
}
