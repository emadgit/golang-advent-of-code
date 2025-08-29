package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strconv"
	"strings"
)

type Board struct {
	cells       []int
	marked      [25]bool
	rowCount    [5]int
	colCount    [5]int
	sumUnmarked int
}

func main() {
	storedInput := helpers.GetInput("../data/day4-data.txt")

	drawNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	fmt.Println(drawNumbers)

	joined := strings.Join(storedInput, " ")
	fields := strings.Fields(joined)

	numbers := make([]int, len(fields))

	for i, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		numbers[i] = n
	}

	boards := [][]int{}
	for i := 0; i < len(numbers); i += 25 {
		boards = append(boards, numbers[i:i+25])
	}

	fmt.Println(boards)
}
