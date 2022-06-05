package term

import "testing"

func TestColourize(t *testing.T) {
	res := Colourize(Red, false)
	if "\033[31m" != res {
		t.Fatalf("Colour red did not match expectations (%s)", res)
	}
	if "\033[91m" != Colourize(Red, true) {
		t.Fatal("Colour bright red did not match expectations")
	}
}

func TestBackgroundColourize(t *testing.T) {
	res := BackgroundColourize(Red, false)
	if "\033[41m" != res {
		t.Fatalf("BG Colour red did not match expectations (%s)", res)
	}
	if "\033[101m" != BackgroundColourize(Red, true) {
		t.Fatal("BG Colour bright red did not match expectations")
	}
}

func TestWrapColour(t *testing.T) {
	res := WrapColour(Red, "foo", false)
	if "\033[31mfoo\033[0m" != res {
		t.Fatalf("Colour red did not match expectations (%s)", res)
	}
}

func TestWrapBackgroundColour(t *testing.T) {
	res := WrapBackgroundColour(Red, "foo", false)
	if "\033[41mfoo\033[0m" != res {
		t.Fatalf("BG Colour red did not match expectations (%s)", res)
	}
}
