package main

import (
	"fmt"

	"github.com/aadarsh10c/AdventOfCode/adventofcode2021"
)

func main() {
	input := adventofcode2021.ReadInput("./adventofcode2021/input.txt")
	value := adventofcode2021.DayOne(input)
	fmt.Printf("Answer: %d", value)
}