package main

import (
	"fmt"
	"learn-go-with-adventofcode/helpers"
	"strconv"
)

func main() {
	storedInput := helpers.GetInput("../data/day3-data.txt")

	copyInput := storedInput

	byteLength := len(storedInput[0])
	var oxygenGeneratorRating int64
	var co2ScrubberRating int64

	zeroBitCount := 0
	oneBitCount := 0

	oxygenGroupCreteria := []string{}
	co2GroupCreteria := []string{}

	zeroBitGroupCreteria := []string{}
	oneBitGroupCreteria := []string{}

	for i := 0; i < byteLength; i++ {
		for _, stringValue := range copyInput {
			if string(stringValue[i]) == "0" {
				zeroBitCount++
				zeroBitGroupCreteria = append(zeroBitGroupCreteria, string(stringValue))
			}

			if string(stringValue[i]) == "1" {
				oneBitCount++
				oneBitGroupCreteria = append(oneBitGroupCreteria, string(stringValue))
			}

		}
		// For OxygenGenerator Rating
		if oneBitCount >= zeroBitCount {
			zeroBitGroupCreteria = []string{}
			oxygenGroupCreteria = append(oxygenGroupCreteria, oneBitGroupCreteria...)
			copyInput = oxygenGroupCreteria
		}

		if oneBitCount < zeroBitCount {
			oneBitGroupCreteria = []string{}
			oxygenGroupCreteria = append(oxygenGroupCreteria, zeroBitGroupCreteria...)
			copyInput = oxygenGroupCreteria
		}
		zeroBitCount = 0
		oneBitCount = 0
		zeroBitGroupCreteria = []string{}
		oneBitGroupCreteria = []string{}
		oxygenGroupCreteria = []string{}
		if len(copyInput) == 1 {
			break
		}
	}
	if len(copyInput) == 1 {
		oxygenGeneratorRating, _ = strconv.ParseInt(copyInput[0], 2, 64)
	}

	copyInput = storedInput
	for i := 0; i < byteLength; i++ {
		fmt.Println(copyInput)

		for _, stringValue := range copyInput {
			if string(stringValue[i]) == "0" {
				zeroBitCount++
				zeroBitGroupCreteria = append(zeroBitGroupCreteria, string(stringValue))
			}

			if string(stringValue[i]) == "1" {
				oneBitCount++
				oneBitGroupCreteria = append(oneBitGroupCreteria, string(stringValue))
			}

		}

		// For CO2 Scrubbing Rating
		if zeroBitCount <= oneBitCount {
			oneBitGroupCreteria = []string{}
			co2GroupCreteria = append(co2GroupCreteria, zeroBitGroupCreteria...)
			copyInput = co2GroupCreteria
		}

		if zeroBitCount > oneBitCount {
			zeroBitGroupCreteria = []string{}
			co2GroupCreteria = append(co2GroupCreteria, oneBitGroupCreteria...)
			copyInput = co2GroupCreteria
		}

		zeroBitCount = 0
		oneBitCount = 0
		zeroBitGroupCreteria = []string{}
		oneBitGroupCreteria = []string{}
		co2GroupCreteria = []string{}
		if len(copyInput) == 1 {
			break
		}
	}
	if len(copyInput) == 1 {
		co2ScrubberRating, _ = strconv.ParseInt(copyInput[0], 2, 64)
		fmt.Println(co2ScrubberRating)
	}

	fmt.Println("oxygenGeneratorRating: ", oxygenGeneratorRating)
	fmt.Println("co2ScrubberRating: ", co2ScrubberRating)
	fmt.Println("Result: ", oxygenGeneratorRating*co2ScrubberRating)
}
