package progress

import (
	"fmt"
	"strings"
	"testing"
)

func TestDefaultProgressBar_Render(t *testing.T) {
	var render string
	current, maximum := uint(0), uint(10)

	bar := newDefaultProgressBar(maximum)

	for {
		render = bar.Render()

		if !strings.Contains(render, fmt.Sprintf("%d/%d", current, maximum)) {
			t.Fatalf("Expected render output to contain '%d/%d', instead we got: '%s'", current, maximum, render)
		}

		if bar.IsFinished() {
			break
		}

		bar.Advance(2)
		current += 2
	}

	bar.Advance(50)
	if !strings.Contains(render, fmt.Sprintf("%d/%d", maximum, maximum)) {
		t.Fatalf("Expected render output to contain '%d/%d', instead we got: '%s'", maximum, maximum, render)
	}
}

func TestDefaultProgressBar_Advance(t *testing.T) {
	bar := newDefaultProgressBar(20)

	bar.Advance(10)
	assertEquals(10, bar.Current(), "Advanced by 10", t)

	bar.Advance(10)
	assertEquals(20, bar.Current(), "Advanced to Maximum", t)

	bar.Advance(10)
	assertEquals(20, bar.Current(), "Cannot exceed maximum", t)
}

func TestDefaultProgressBar_Current(t *testing.T) {
	bar := newDefaultProgressBar(3)

	assertEquals(0, bar.Current(), "Mismatched Current", t)
	bar.Advance(1)
	assertEquals(1, bar.Current(), "Mismatched Current", t)
	bar.Finish()
	assertEquals(3, bar.Current(), "Mismatched Current", t)
}

func TestDefaultProgressBar_Maximum(t *testing.T) {
	bar := newDefaultProgressBar(3)

	assertEquals(3, bar.Maximum(), "Mismatched Maximum", t)
}

func TestDefaultProgressBar_SetMaximum(t *testing.T) {
	bar := newDefaultProgressBar(3)

	assertEquals(3, bar.Maximum(), "Mismatched Maximum", t)

	bar.SetMaximum(4)
	assertEquals(4, bar.Maximum(), "Mismatched Maximum", t)
}

func TestDefaultProgressBar_Finish(t *testing.T) {
	bar := newDefaultProgressBar(1)
	bar.Finish()

	if bar.IsFinished() != true {
		t.Fatal("Expected bar to be finished")
	}
}

func TestDefaultProgressBar_IsFinished(t *testing.T) {
	bar := newDefaultProgressBar(3)

	bar.Advance(3)
	if bar.IsFinished() != true {
		t.Fatal("Expected bar to be finished")
	}
}

func assertEquals(expected uint, given uint, message string, t *testing.T) {
	if expected != given {
		t.Fatalf("Expected %d to be %d: %s", given, expected, message)
	}
}
