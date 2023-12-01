package common

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func IsTestMode(args []string) bool {
	return len(args) > 1 && args[1] == "--test"
}

func ReadInput(day int, args []string) ([]string, error) {
	isTest := IsTestMode(args)
	var inputFile string
	if isTest {
		inputFile = filepath.Join(".", "test_inputs", fmt.Sprintf("%d.input", day))
	} else {
		inputFile = filepath.Join(".", "inputs", fmt.Sprintf("%d.input", day))
	}

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
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
