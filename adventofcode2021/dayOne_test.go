package adventofcode2021

import (
	"testing"
)

func TestDayOne(t *testing.T) {

	t.Run("count number of increased measurements", func(t *testing.T) {
		input := ReadInput("input.txt")
		got := DayOne(input)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("three-measurement slide window", func(t *testing.T) {
		input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := GetSlideList(input)
		want := []int{607, 618, 618, 617, 647, 716, 769, 792}
		comapreSlice(got, want, t)
	})

}

func comapreSlice(got, want []int, t *testing.T) {
	t.Helper()
	gotLen, wantLen := len(got), len(want)
	if gotLen == wantLen {
		for i, _ := range got {
			if got[i] != want[i] {
				t.Fatalf("slices are unequal, got : %v ,wanted: %v", got, want)
			}
		}
	} else {
		t.Errorf("lengths are unequal, got length: %d, want length: %d", gotLen, wantLen)
	}
}
