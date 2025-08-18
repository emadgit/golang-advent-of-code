package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strings"
)

func main() {
	storedInput := helpers.GetInput("../data/day4-data.txt")

	drawNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	fmt.Println(drawNumbers)
	boards := make([][]string, 3)
	boards[0] = make([]string, 25)
	boards[1] = make([]string, 25)
	boards[2] = make([]string, 25)

	fields := []string{}

	for _, stringValue := range storedInput {
		fields = strings.Fields(stringValue)
	}

	counter := 0

	for i := 0; i < 3; i++ {
		boards[i] = fields[counter : counter+25]
		counter += 25
	}

	fmt.Println(boards)
	// OKAY! So far we have the boards in right format, now I need to solve it, I CAN DO IT!

}
