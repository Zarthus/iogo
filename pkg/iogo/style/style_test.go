package style

import (
	"github.com/zarthus/iogo/v2/pkg/iogo"
	"github.com/zarthus/iogo/v2/test"
	"testing"
)

func TestDefaultStyle_Input(t *testing.T) {
	r, w := test.NewNullReader("y"), test.NewNullWriter()
	style := createDefaultStyle(r, w)

	style.Input()
}

func TestDefaultStyle_Output(t *testing.T) {
	r, w := test.NewNullReader("y"), test.NewNullWriter()
	style := createDefaultStyle(r, w)

	style.Output()
}

func TestDefaultStyle(t *testing.T) {
	r, w := test.NewNullReader("y"), test.NewNullWriter()
	style := createDefaultStyle(r, w)

	if confirmed, err := style.Input().Confirm("foo", iogo.Options{}); err != nil {
		t.Fatalf("error: %s", err.Error())
	} else if !confirmed {
		t.Fatalf("Expected to be confirmed")
	}

	if len(w.Get()) != 1 {
		t.Fatalf("Expected writer to have 1 element, got %d", len(w.Get()))
	}
	if w.Get()[0] != "foo (Y/n)" {
		t.Fatalf("Expected writer to have 'foo', got %s", w.Get()[0])
	}

	style.Output().Title("foobar!")
	if len(w.Get()) != 4 {
		t.Fatalf("Expected writer to have 4 elements, got %d", len(w.Get()))
	}
}
