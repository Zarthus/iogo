package history

import "testing"

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

func assertLength(t *testing.T, length int, actual int) {
	if length != actual {
		t.Fatalf("Expected history count of %d, got %d", length, actual)
	}
}
