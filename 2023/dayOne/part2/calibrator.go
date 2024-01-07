package part2

import 

var digitMap = map[string]int{
	"one": 1,
	"two":2,
	"three":3,
	"four":4,
	"five":5,
	"six":6,
	"seven":7,
	"eight":8,
	"nine":9,
}

func StartsWith(matcher,container string) (bool,error) {
	matcher_len = len(matcher)
	container_len = len(container)
	var flag int
	if (matcher_len > container_len) {
		return false,errors.New("length of the matcher is greater than container")
	}

	for i:= 0 ; i<matcher_len ; i ++ {
		if (matcher[i] != container[i]){
			return false,nil //no match immedialtey return
		}
	}

	return true,nil //match

}

func mapToDigit(matcher string) (bool,int){\
	var digit int
	var found bool

	for key,_ := range digitMap{
		if matcher startwith{
			if len(matcher) === len(key){
				digit = digitdigitMap[key]
				found = true
				return digit,found //break out of loop 
			}else {
				//we found i startwith
				found =true
				return digit,found
			}
		}
	}
}

func CalibratorTwo(line string) int{
	matcher := ""
	var	digitList []int
	for _,char := range line {
		matcher += string(char)
		match,digit := mapToDigit(matcher)
	}

}