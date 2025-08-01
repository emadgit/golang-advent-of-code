package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../data/data.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
