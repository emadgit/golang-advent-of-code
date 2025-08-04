package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strconv"
)

func main() {
	storedInput := helpers.GetInput("../data/day3-data.txt")

	byteLength := len(storedInput[0])

	zeroBitCount := 0
	oneBitCount := 0
	gammaRate := ""
	epsilonRate := ""
	for i := 0; i < byteLength; i++ {
		for _, stringValue := range storedInput {
			if string(stringValue[i]) == "0" {
				zeroBitCount++
			}

			if string(stringValue[i]) == "1" {
				oneBitCount++
			}

		}
		if oneBitCount > zeroBitCount {
			gammaRate += "1"
			epsilonRate += "0"
		} else if oneBitCount < zeroBitCount {
			gammaRate += "0"
			epsilonRate += "1"
		}

		oneBitCount = 0
		zeroBitCount = 0
	}

	fmt.Println("gammaRate: ", gammaRate)
	fmt.Println("epsilonRate: ", epsilonRate)

	gammaDecimal, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Println("gammaDecimal: ", gammaDecimal)
	fmt.Println("epsilonDecimal: ", epsilonDecimal)
	fmt.Println("Result: ", gammaDecimal*epsilonDecimal)
}

// Day 3 - Part 1 -> https://adventofcode.com/2021/day/3
