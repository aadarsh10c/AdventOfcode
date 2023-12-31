package adventofcode2021

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(fileName string) []int {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	contentStr := string(content)
	//remove \r\n
	contentStrList := strings.Split(contentStr, "\r\n")

	input := make([]int, len(contentStrList))
	for index, strValue := range contentStrList {
		input[index], _ = strconv.Atoi(strValue)
	}
	return input
}

func DayOne(input []int) int {
	counter := 0
	minVal := 0
	for index, currValue := range input {
		if index > 0 {
			minVal = input[index-1]
		} else {
			//because index is 0
			minVal = currValue
		}
		if currValue > minVal {
			counter++
		}
	}
	return counter
}

func GetSlideList(input []int) []int {
	slideListLen := len(input) - 2
	slideList := make([]int, slideListLen)
	for i, _ := range slideList {
		first, second, third := input[i], input[i+1], input[i+2]
		slideList[i] = first + second + third
	}
	return slideList
}
