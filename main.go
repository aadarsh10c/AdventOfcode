package main

import (
	"fmt"

	"github.com/aadarsh10c/AdventOfCode/2023/util"
	"github.com/aadarsh10c/AdventOfCode/2023/dayOne"
)

func main() {
	//read input rom the text file
	faultyCalibrations := util.ReadInput("./2023/input/dayOneInput.txt")
	// fmt.Printf("%v",faultyCalibrations)
	//calculate the sum of all calibration values
	calibrationValue := dayOne.TotalCalibarationValue(faultyCalibrations)
	fmt.Printf("Final value : %d",calibrationValue)
}
