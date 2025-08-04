package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var storedInput []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		storedInput = append(storedInput, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return storedInput
}
