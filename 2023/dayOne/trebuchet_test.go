package dayOne

import (
	"reflect"
	"testing"
)

func TestTrebuchet(t *testing.T) {
	t.Run("extract number from a string", func(t *testing.T) {
		input := "1abc2"
		got := DoubleDigitExtractor(input)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , wantted %v", got, want)
		}

	})
}
