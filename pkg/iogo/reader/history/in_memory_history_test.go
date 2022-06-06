package history

import (
	"strings"
	"testing"
)

func TestTracker_Get(t *testing.T) {
	hist := NewHistoryTracker([]string{"foo", "bar"})

	assertLength(t, 2, len(hist.Get()))
}

func TestTracker_Reset(t *testing.T) {
	hist := NewHistoryTracker([]string{"foo", "bar"})
	hist.Reset()

	assertLength(t, 0, len(hist.Get()))
}

func TestTracker_Track(t *testing.T) {
	hist := NewHistoryTracker([]string{})

	assertLength(t, 0, len(hist.Get()))

	hist.Track("foo")
	assertLength(t, 1, len(hist.Get()))

	hist.Track("bar")
	assertLength(t, 2, len(hist.Get()))
}

func TestTracker_Untrack(t *testing.T) {
	hist := NewHistoryTracker([]string{"foo", "bar", "baz"})

	assertLength(t, 3, len(hist.Get()))

	hist.Untrack("bar")
	assertLength(t, 2, len(hist.Get()))

	if strings.Contains(strings.Join(hist.Get(), ","), "bar") {
		t.Fatalf("Expected element 'bar' not to be in list: %s", strings.Join(hist.Get(), ","))
	}
}

func assertLength(t *testing.T, length int, actual int) {
	if length != actual {
		t.Fatalf("Expected history count of %d, got %d", length, actual)
	}
}
