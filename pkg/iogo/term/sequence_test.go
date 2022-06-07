package term

import (
	"github.com/zarthus/iogo/v2/pkg/iogo/term/col"
	"testing"
)

func TestColourize(t *testing.T) {
	res := Colourize(col.Red, false)
	if res != "\033[31m" {
		t.Fatalf("Colour red did not match expectations (%s)", res)
	}
	if Colourize(col.Red, true) != "\033[91m" {
		t.Fatal("Colour bright red did not match expectations")
	}
}

func TestBackgroundColourize(t *testing.T) {
	res := BackgroundColourize(col.Red, false)
	if res != "\033[41m" {
		t.Fatalf("BG Colour red did not match expectations (%s)", res)
	}
	if BackgroundColourize(col.Red, true) != "\033[101m" {
		t.Fatal("BG Colour bright red did not match expectations")
	}
}

func TestWrapColour(t *testing.T) {
	res := WrapColour(col.Red, "foo", false)
	if res != "\033[31mfoo\033[0m" {
		t.Fatalf("Colour red did not match expectations (%s)", res)
	}
}

func TestWrapBackgroundColour(t *testing.T) {
	res := WrapBackgroundColour(col.Red, "foo", false)
	if res != "\033[41mfoo\033[0m" {
		t.Fatalf("BG Colour red did not match expectations (%s)", res)
	}
}
