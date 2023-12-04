package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(games []string, debug bool) {
	redCubesAtStart := 12
	greenCubesAtStart := 13
	blueCubesAtStart := 14
	sum := 0
	for _, game := range games {
		split := strings.Split(game, ": ")
		gameId, _ := strconv.Atoi(strings.Split(split[0], "Game ")[1])
		output := split[1]

		draws := strings.Split(output, "; ")
		validGame := true
		for _, draw := range draws {
			drawCubesAmountsAndColors := strings.Split(draw, ", ")
			redCubes := redCubesAtStart
			greenCubes := greenCubesAtStart
			blueCubes := blueCubesAtStart
			for _, amountAndColor := range drawCubesAmountsAndColors {
				amountAndColorSplit := strings.Split(amountAndColor, " ")
				amount, _ := strconv.Atoi(amountAndColorSplit[0])
				color := amountAndColorSplit[1]
				switch color {
				case "red":
					redCubes -= amount
				case "green":
					greenCubes -= amount
				case "blue":
					blueCubes -= amount
				}
			}
			if redCubes < 0 || greenCubes < 0 || blueCubes < 0 {
				if debug {
					fmt.Println(gameId, output)
					fmt.Printf("Game %v: %v\n", gameId, "invalid")
				}
				validGame = false
				break
			}
		}
		if validGame {
			if debug {
				fmt.Println(gameId, output)
				fmt.Printf("Game %v: %v\n", gameId, "valid")
			}
			sum += gameId
		}
	}
	fmt.Printf("Sum: %v\n", sum)
}

func partTwo(games []string, debug bool) {
	sumOfPowers := 0
	for _, game := range games {
		minRedCubes := 0
		minGreenCubes := 0
		minBlueCubes := 0
		split := strings.Split(game, ": ")
		output := split[1]

		draws := strings.Split(output, "; ")
		for _, draw := range draws {
			drawCubesAmountsAndColors := strings.Split(draw, ", ")

			for _, amountAndColor := range drawCubesAmountsAndColors {
				amountAndColorSplit := strings.Split(amountAndColor, " ")
				amount, _ := strconv.Atoi(amountAndColorSplit[0])
				color := amountAndColorSplit[1]
				switch color {
				case "red":
					if amount > minRedCubes {
						minRedCubes = amount
					}
				case "green":
					if amount > minGreenCubes {
						minGreenCubes = amount
					}
				case "blue":
					if amount > minBlueCubes {
						minBlueCubes = amount
					}
				}
			}
		}
		power := minRedCubes * minGreenCubes * minBlueCubes
		sumOfPowers += power
	}
	fmt.Printf("Sum of powers: %v\n", sumOfPowers)
}

func main() {
	day := 2
	games, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	part := common.GetPart(os.Args)
	debug := common.IsTestMode(os.Args)
	fmt.Printf("Part: %v\n", part)
	if part == 1 {
		partOne(games, debug)
	} else {
		partTwo(games, debug)
	}
}
