package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strings"
)

func partOne(lines []string, debug bool) {
	totalPoints := 0
	for i, line := range lines {
		points := 0
		split := strings.Split(line, ": ")
		cards := split[1]

		fmt.Println(i)
		cardsSplit := strings.Split(cards, " | ")
		winningNumbers := strings.Fields(cardsSplit[0])
		winningNumbersMap := make(map[string]bool)
		for _, winningNumber := range winningNumbers {
			winningNumbersMap[winningNumber] = true
		}
		myNumbers := strings.Fields(cardsSplit[1])
		for _, myNumber := range myNumbers {
			if _, ok := winningNumbersMap[myNumber]; ok {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
				fmt.Println("Winning number", myNumber, points)
			}
		}
		totalPoints += points
	}
	fmt.Printf("Total points: %v\n", totalPoints)
}

func partTwo(lines []string, debug bool) {
	amountsOfCards := make([]int, len(lines)+1)

	for i, line := range lines {
		amountsOfCards[i+1] += 1
		amountOfCards := amountsOfCards[i+1]
		cardsWon := 0
		split := strings.Split(line, ": ")
		cards := split[1]

		if debug {
			fmt.Println(i + 1)
		}
		cardsSplit := strings.Split(cards, " | ")
		winningNumbers := strings.Fields(cardsSplit[0])
		winningNumbersMap := make(map[string]bool)
		for _, winningNumber := range winningNumbers {
			winningNumbersMap[winningNumber] = true
		}
		myNumbers := strings.Fields(cardsSplit[1])
		for _, myNumber := range myNumbers {
			if _, ok := winningNumbersMap[myNumber]; ok {
				cardsWon++
			}
		}

		if debug {
			fmt.Println("Cards won", cardsWon, amountOfCards, cardsWon*amountOfCards)
		}
		for ii := i + 1; ii <= i+cardsWon; ii++ {
			if debug {
				fmt.Println("Card", ii+1, amountsOfCards[ii+1]+1)
			}
			amountsOfCards[ii+1] += amountOfCards
		}
	}
	totalCards := 0
	for _, amountOfCards := range amountsOfCards {
		totalCards += amountOfCards
	}
	fmt.Printf("Total cards: %v\n", totalCards)
}

func main() {
	day := 4
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
