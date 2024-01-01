package dayOne

import "errors"

const errEmptyString = "string is empty"

func ExtractFromString(line string) ([]int, error) {
	var digits []int
	if line == "" {
		return digits , errors.New(errEmptyString)
	}
	for _, runeString := range line {
		if runeString >= 48 && runeString <= 57 {
			digits = append(digits, int(runeString)-48)
		}
	}
	return digits,nil
}

func DoubleDigitExtractor(line string) ([]int, string) {
	digits,isError := 	ExtractFromString(line)
	if isError != nil{
		return digits,isError.Error()
	}
	dig_len := len(digits)
	if dig_len > 2 {
		newList := []int{digits[0], digits[dig_len-1]}
		return newList,""
	}else if dig_len == 1 {
		newList := []int{digits[0], digits[0]}
		return newList,""
	}

	return digits,"	"
}
