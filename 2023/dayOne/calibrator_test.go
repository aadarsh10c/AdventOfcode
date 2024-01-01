package dayOne

import (
	"reflect"
	"testing"
)

func TestCalibrator(t *testing.T) {
	t.Run("extract number from a string", func(t *testing.T) {
		input := "treb7uchet"
		got,_ := ExtractFromString(input)
		want := []int{7}

		compareValue(got, want, t)

	})
	t.Run("non-empty string is send,expected test to fail", func(t *testing.T) {
		input := "1"
		_,got := ExtractFromString(input)
		want := errEmptyString
		if got == nil{
			t.Errorf("Expected %v errror, but got %v",want,got)
		}
	})
	t.Run("empty string is send,expected test to pass", func(t *testing.T) {
		input := ""
		_,got := ExtractFromString(input)
		want := errEmptyString
		if got.Error() != want{
			t.Errorf("Expected %v errror, but got %v",want,got)
		}
	})
	t.Run("empty string is sent ,expect error",func(t *testing.T){
		input := ""
		_,got := DoubleDigitExtractor(input)
		want := errEmptyString
		if got != want {
			t.Errorf("Expected %v errror, but got %v",want,got)
		}
	})
	t.Run("expect 2 digits",func(t *testing.T) {
		input := "1abc2"
		got,_ := DoubleDigitExtractor(input)
		want := []int{1,2}
		compareValue(got, want, t)
	})
	t.Run("expect 2 digits case 2",func(t *testing.T) {
		input := "a1b2c3d4e5f"
		got,_ := DoubleDigitExtractor(input)
		want := []int{1,5}
		compareValue(got, want, t)
	})
	t.Run("expect 2 digits case 3",func(t *testing.T) {
		input := "treb7uchet"
		got,_ := DoubleDigitExtractor(input)
		want := []int{7,7}
		compareValue(got, want, t)
	})
}

func compareValue(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v , wantted %v", got, want)
	}
}
