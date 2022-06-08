package reader

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestNewReaderStyle(t *testing.T) {
	NewReaderStyle(test.NewNullReader("foo"), test.NewNullWriter())
}

func TestReaderStyle_Confirm(t *testing.T) {
	r := NewReaderStyle(test.NewNullReader("y"), test.NewNullWriter())

	confirmed, err := r.Confirm("foo", iogo.Options{})
	if err != nil {
		t.Fatalf("unexpected err, %s", err.Error())
	}

	if !confirmed {
		t.Fatalf("expected confirmed to be true")
	}
}

func TestReaderStyle_Select(t *testing.T) {
	r := NewReaderStyle(test.NewNullReader("foo"), test.NewNullWriter())

	_, err := r.Select("foo", []string{"foo"}, iogo.Options{})
	if err != nil {
		t.Fatalf("unexpected err, %s", err.Error())
	}
}

func TestReaderStyle_Prompt(t *testing.T) {
	r := NewReaderStyle(test.NewNullReader("foo"), test.NewNullWriter())

	if _, err := r.Prompt("foo", iogo.Options{}); err != nil {
		t.Fatalf("unexpected err, %s", err.Error())
	}
}
