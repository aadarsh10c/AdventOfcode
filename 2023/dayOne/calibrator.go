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

func DoubleDigitExtractor(line string) (int, string) {
	digits,isError := 	ExtractFromString(line)
	if isError != nil{
		return 0,isError.Error()
	}
	dig_len := len(digits)
	if dig_len > 2 {
		dblDigit  := (digits[0] * 10 ) +  digits[dig_len-1]
		return dblDigit,""
	}else if dig_len == 1 {
		dblDigit := (digits[0] * 10 ) +  digits[0]
		return dblDigit,""
	}
	dblDigit := (digits[0] * 10) + digits[1]
	return dblDigit,""
}

func TotalCalibarationValue (calibList []string) int {
	var sum int = 0
	for _,line := range calibList{
		value,err := DoubleDigitExtractor(line)
		if err != ""{
			return 0
		}else{
			sum += value
		}
	}
	return sum  
}
