package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strconv"
)

func hasSymbol(lines []string, fromX int, toX int, fromY int, toY int) bool {
	// change coordinates to be within the bounds of the lines
	if fromX < 0 {
		fromX = 0
	}
	maxX := len([]rune(lines[0])) - 1
	if toX > maxX {
		toX = maxX
	}
	if fromY < 0 {
		fromY = 0
	}
	if toY >= len(lines) {
		toY = len(lines) - 1
	}

	for y := fromY; y <= toY; y++ {
		for x := fromX; x <= toX; x++ {
			char := []rune(lines[y])[x]
			// if char is not a digit or a '.' then it's a symbol
			if (char < '0' || char > '9') && char != '.' {
				return true
			}
		}
	}
	return false
}

func partOne(lines []string, debug bool) {
	sum := 0
	for y, line := range lines {
		fmt.Println("Line", y)
		currentNumber := ""
		for x, char := range line {
			checkForSymbol := false
			eol := x == len(line)-1
			charIsDigit := char >= '0' && char <= '9'
			if charIsDigit {
				currentNumber += string(char)
				if eol {
					checkForSymbol = true
				}

			} else if currentNumber != "" {
				checkForSymbol = true
			}
			if checkForSymbol {
				number, _ := strconv.Atoi(currentNumber)
				if hasSymbol(lines, x-len(currentNumber)-1, x, y-1, y+1) {
					fmt.Printf("Number: %v %v \n", number, "has symbol")
					sum += number
				} else {
					fmt.Printf("Number: %v %v \n", number, "does not have symbol")
				}
				currentNumber = ""
			}
		}
	}
	fmt.Printf("Sum: %v\n", sum)
}

func partTwo(lines []string, debug bool) {

}

func main() {
	day := 3
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	part := common.GetPart(os.Args)
	debug := common.IsTestMode(os.Args)
	fmt.Printf("Part: %v\n", part)
	if part == 1 {
		partOne(lines, debug)
	} else {
		partTwo(lines, debug)
	}
}
