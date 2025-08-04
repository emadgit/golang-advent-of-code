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
	aim := 0

	for _, stringValue := range storedInput {
		valueParts := strings.Fields(stringValue)

		command := valueParts[0]
		value, _ := strconv.Atoi(valueParts[1])

		switch command {
		case "up":
			aim -= value
		case "down":
			aim += value

		case "forward":
			horizental += value
			depth += aim * value
		}
	}

	fmt.Println("Horizontal: ", horizental)
	fmt.Println("Depth: ", depth)
	fmt.Println("Aim: ", aim)
	fmt.Println("Result: ", horizental*depth)
}
