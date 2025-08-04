package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strconv"
	"strings"
)

func main() {
	storedInput := helpers.GetInput("../data/day2-data.txt")

	horizental := 0
	depth := 0

	for _, stringValue := range storedInput {
		lineParts := strings.Fields(stringValue)

		stringCommandValue := lineParts[0]
		numericCommandValue, err := strconv.Atoi(lineParts[1])

		if err != nil {
			fmt.Println(err)
		}

		if stringCommandValue == "forward" {
			horizental += numericCommandValue
		}

		if stringCommandValue == "up" {
			depth -= numericCommandValue
		}

		if stringCommandValue == "down" {
			depth += numericCommandValue
		}
	}

	fmt.Println("horizental: ", horizental)
	fmt.Println("depth: ", depth)
	fmt.Println("Answer: ", horizental*depth)
}
