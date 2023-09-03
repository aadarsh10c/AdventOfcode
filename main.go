package main
import "fmt"

func main() {
	input := ReadInput("input.txt")
	value := DayOne(input)
	fmt.Printf("Answer: %d", value)
}
