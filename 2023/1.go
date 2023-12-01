package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day := 1
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	part := common.GetPart(os.Args)
	debug := common.IsTestMode(os.Args)
	fmt.Printf("Part: %v\n", part)

	sum := 0
	for _, line := range lines {
		firstNumber := ""
		lastNumber := ""
		for i, char := range line {

			if char >= '0' && char <= '9' {
				if firstNumber == "" {
					firstNumber = string(char)
				}
				lastNumber = string(char)
			} else if part == 2 {
				// if substring starting from this char including next four chars includes any string in digits array
				// then add the number to the sum
				endIndex := i + 5
				if endIndex >= len(line) {
					endIndex = len(line)
				}
				if (endIndex - i) < 3 {
					continue
				}
				inputSlice := line[i:endIndex]

				if debug {
					fmt.Printf("Compare outer: %v %v %v\n", inputSlice, i, endIndex)
				}
				for digitInt, digitString := range digits {
					if len(inputSlice) < len(digitString) {
						continue
					}

					if debug {
						fmt.Printf("Compare: %v %v, %v %v\n", len(digitString), len(inputSlice), inputSlice[:(len(digitString))], digitString)
					}
					if inputSlice[:(len(digitString))] == digitString {

						if debug {
							fmt.Printf("Digit found in '%v': %v %v\n", inputSlice, digitInt, digitString)
						}
						if firstNumber == "" {
							firstNumber = strconv.Itoa(digitInt)
						}
						lastNumber = strconv.Itoa(digitInt)
						break
					}
				}

			}
		}
		if debug {
			fmt.Printf("Full row: %v First number: %v, last number: %v\n", line, firstNumber, lastNumber)
		}
		numberString := firstNumber + lastNumber
		number, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Printf("Error converting string to int: %s\n", err)
			return
		}
		sum += number
	}

	// Add your problem-solving logic here and print the result.
	fmt.Println(sum)
}
