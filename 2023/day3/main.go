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
	maxX := len(lines[0]) - 1
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
		line := lines[y]
		for x := fromX; x <= toX; x++ {
			char := line[x]
			// if char is not a digit or a '.' then it's a symbol
			if (char < '0' || char > '9') && char != '.' {
				return true
			}
		}
	}
	return false
}

func findNumberAtCoordinate(lines []string, x int, y int) (int, int) {
	number := string(lines[y][x])
	for xx := x - 1; xx >= 0; xx-- {
		char := lines[y][xx]
		if char >= '0' && char <= '9' {
			number = string(char) + number
		} else {
			break
		}
	}
	for x = x + 1; x < len(lines[y]); x++ {
		char := lines[y][x]
		if char >= '0' && char <= '9' {
			number = number + string(char)
		} else {
			break
		}
	}
	numberInt, _ := strconv.Atoi(number)
	return numberInt, x
}

func gearRatio(lines []string, x int, y int) int {
	numberOne := -1
	numberTwo := -1
	startY := y - 1
	if startY < 0 {
		startY = 0
	}
	startX := x - 1
	if startX < 0 {
		startX = 0
	}
	for yy := startY; yy <= y+1 && yy < len(lines); yy++ {
		for xx := startX; xx <= x+1 && xx < len(lines[0]); xx++ {
			if xx == x && yy == y {
				continue
			}
			char := lines[yy][xx]
			if char >= '0' && char <= '9' {
				if numberOne == -1 {
					numberOne, xx = findNumberAtCoordinate(lines, xx, yy)
				} else if numberTwo == -1 {
					numberTwo, xx = findNumberAtCoordinate(lines, xx, yy)
					fmt.Printf("Number one: %v, number two: %v\n", numberOne, numberTwo)
					// convert to int and return numbers
					return numberOne * numberTwo
				}
			}
		}
	}

	return -1
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
	sum := 0
	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				gearRatio := gearRatio(lines, x, y)
				if gearRatio >= 0 {
					fmt.Printf("Gear at %v,%v %v\n", x, y+1, gearRatio)
					sum += gearRatio
				}
			}
		}
	}
	fmt.Printf("Sum: %v\n", sum)
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
