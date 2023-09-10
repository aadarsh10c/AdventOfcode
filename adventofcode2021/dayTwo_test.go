package adventofcode2021

import "testing"

func TestCalPosition(t *testing.T) {
	testInput := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	got := CalPosition(testInput)
	want := 150

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
