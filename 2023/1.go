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

	sum := 0
	for _, line := range lines {
		firstNumber := ""
		lastNumber := ""
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if firstNumber == "" {
					firstNumber = string(char)
				}
				lastNumber = string(char)
			}
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
