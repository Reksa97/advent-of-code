package main

import (
	"aoc2023/common"
	"fmt"
	"os"
)

func main() {
	day := 1
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
		// Add your logic here
	}

	// Add your problem-solving logic here and print the result.
	fmt.Println("Result: [Your Solution Here]")
}
