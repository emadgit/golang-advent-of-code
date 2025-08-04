package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../data/data.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	groupCount := 0
	groupSum := 0

	var storedInput []int
	var groupSumInSlice []int

	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		storedInput = append(storedInput, num)
	}

	for outerIndex, _ := range storedInput {
		for _, num := range storedInput[outerIndex : outerIndex+3] {
			groupSum = groupSum + num
			groupCount = groupCount + 1

			if groupCount == 3 {
				groupSumInSlice = append(groupSumInSlice, groupSum)
				groupCount = 0
				groupSum = 0
			}
		}
	}

	prevGroupSum := 0
	increasedGroupSum := 0

	for _, sumForEachGroup := range groupSumInSlice {
		if (prevGroupSum != 0) && (sumForEachGroup > prevGroupSum) {
			increasedGroupSum++
		}

		prevGroupSum = sumForEachGroup
	}

	fmt.Println("increasedGroupSum: ", increasedGroupSum)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// Solution for the second part https://adventofcode.com/2021/day/1
