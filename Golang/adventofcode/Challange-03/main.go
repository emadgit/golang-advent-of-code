package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../data/day2-data.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizental := 0
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Fields(line)

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

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
