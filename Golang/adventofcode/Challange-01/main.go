package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open the file
	file, err := os.Open("../data/data.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)

	prevNum := 0
	increasedCount := 0

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
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

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}
