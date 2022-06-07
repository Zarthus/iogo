package term

import "testing"

func TestColourize(t *testing.T) {
	res := Colourize(Red, false)
	if res != "\033[31m" {
		t.Fatalf("Colour red did not match expectations (%s)", res)
	}
	if Colourize(Red, true) != "\033[91m" {
		t.Fatal("Colour bright red did not match expectations")
	}
}

func TestBackgroundColourize(t *testing.T) {
	res := BackgroundColourize(Red, false)
	if res != "\033[41m" {
		t.Fatalf("BG Colour red did not match expectations (%s)", res)
	}
	if BackgroundColourize(Red, true) != "\033[101m" {
		t.Fatal("BG Colour bright red did not match expectations")
	}
}

func TestWrapColour(t *testing.T) {
	res := WrapColour(Red, "foo", false)
	if res != "\033[31mfoo\033[0m" {
		t.Fatalf("Colour red did not match expectations (%s)", res)
	}
}

func TestWrapBackgroundColour(t *testing.T) {
	res := WrapBackgroundColour(Red, "foo", false)
	if res != "\033[41mfoo\033[0m" {
		t.Fatalf("BG Colour red did not match expectations (%s)", res)
	}
}
