package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strconv"
)

func main() {
	storedInput := helpers.GetInput("../data/data.txt")

	prevNum := 0
	increasedCount := 0
	// Iterate over each line
	for _, stringValue := range storedInput {
		num, err := strconv.Atoi(stringValue)
		if err != nil {
			fmt.Println("Conversion error:", err)
			return
		}

		if (prevNum != 0) && (num > prevNum) {
			increasedCount++
		}

		prevNum = num
	}

	fmt.Println("Increased count:", increasedCount)
}
