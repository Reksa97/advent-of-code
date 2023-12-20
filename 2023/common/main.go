package common

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Function to calculate GCD
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate LCM of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate LCM of an array of integers
func LcmArray(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func ConvertToInt(values []string) []int {
	var result []int
	for _, value := range values {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Sprintf("Error converting %v to int: %v", value, err))
		}
		result = append(result, intValue)
	}
	return result
}

func GetPart(args []string) int {
	// Determine which part to run, default to Part 1
	part := 1
	if len(os.Args) > 1 {
		var err error
		part, err = strconv.Atoi(os.Args[1])
		if err != nil || (part != 1 && part != 2) {
			fmt.Println("Invalid part argument. Defaulting to 1.")
			return 1
		}
	}
	return part
}

func IsTestMode(args []string) bool {
	return len(args) > 2 && args[2] == "--test"
}

func ReadInput(day int, args []string) ([]string, error) {
	isTest := IsTestMode(args)
	if isTest {
		fmt.Println("Test mode")
	}
	part := GetPart(args)
	var inputFile string
	if isTest {
		if part == 1 {
			inputFile = filepath.Join(".", "test_inputs", fmt.Sprintf("%d.1.input", day))
		} else {
			inputFile = filepath.Join(".", "test_inputs", fmt.Sprintf("%d.2.input", day))
		}
	} else {
		inputFile = filepath.Join(".", "inputs", fmt.Sprintf("%d.input", day))
	}

	file, err := os.Open(inputFile)
	if err != nil {
		inputFile = filepath.Join(".", "test_inputs", fmt.Sprintf("%d.input", day))
		file, err = os.Open(inputFile)
		if err != nil {
			return nil, fmt.Errorf("error opening input file: %w", err)
		}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file: %w", err)
	}

	return lines, nil
}
