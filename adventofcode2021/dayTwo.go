package adventofcode2021

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// takes string numbers returns the parsed number
func StrToInt(strNum string) (number int) {
	number, err := strconv.Atoi(strNum)
	if err != nil {
		log.Fatal("String numbers an only be parsed ", err)
	}
	return
}

func CalPosition(input []string) int {
	//extract value in int
	horizontal := 0
	depth := 0

	for _, value := range input {
		valList := strings.Split(value, " ")
		fmt.Printf("%q", valList)
		operation, operand := valList[0], valList[1]
		switch operation {
		case "forward":
			horizontal += StrToInt(operand)
		case "up":
			depth -= StrToInt(operand)
		case "down":
			depth += StrToInt(operand)
		}
	}

	return horizontal * depth
}
