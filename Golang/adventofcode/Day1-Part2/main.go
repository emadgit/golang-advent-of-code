package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strconv"
)

func main() {
	storedInput := helpers.GetInput("../data/data.txt")

	groupSum := 0
	groupCount := 0
	var groupSumInSlice []int

	for outerIndex, _ := range storedInput {
		for _, stringNum := range storedInput[outerIndex : outerIndex+3] {
			num, _ := strconv.Atoi(stringNum)
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
}

// Solution for the second part https://adventofcode.com/2021/day/1
