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

}
