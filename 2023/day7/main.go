package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	cards    string
	handType HandType
	bid      int
}

func calculateHandType(cards string, debug bool) HandType {
	cardMap := make(map[string]int)
	for _, card := range cards {
		cardMap[string(card)]++
	}
	if debug {
		fmt.Println(cardMap)
	}

	maxOfAKind := 0
	for _, amount := range cardMap {
		if amount > maxOfAKind {
			maxOfAKind = amount
		}
	}

	if maxOfAKind == 5 {
		return FiveOfAKind
	}
	if maxOfAKind == 4 {
		return FourOfAKind
	}
	if maxOfAKind == 3 {
		if len(cardMap) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if maxOfAKind == 2 {
		if len(cardMap) == 3 {
			return TwoPair
		}
		return OnePair
	}

	return HighCard
}

func calculateHandTypePart2(cards string, debug bool) HandType {
	cardMap := make(map[string]int)
	for _, card := range cards {
		cardMap[string(card)]++
	}
	if debug {
		fmt.Println(cardMap)
	}

	jokerCount := cardMap["J"]
	delete(cardMap, "J")

	maxOfAKind := 0
	maxOfAKindCard := ""
	for card, amount := range cardMap {
		if amount > maxOfAKind {
			maxOfAKind = amount
			maxOfAKindCard = card
		}
	}

	if maxOfAKind == 0 && jokerCount == 5 {
		maxOfAKind = 5
		maxOfAKindCard = "A"
	} else if jokerCount > 0 {
		for _, amount := range cardMap {
			if amount == maxOfAKind {
				maxOfAKind += jokerCount
				cardMap[maxOfAKindCard] += jokerCount
				break
			}
		}
	}

	if maxOfAKind == 5 {
		return FiveOfAKind
	}
	if maxOfAKind == 4 {
		return FourOfAKind
	}
	if maxOfAKind == 3 {
		if len(cardMap) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if maxOfAKind == 2 {
		if len(cardMap) == 3 {
			return TwoPair
		}
		return OnePair
	}

	return HighCard

}

type ByHand []Hand

var part = common.GetPart(os.Args)
var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var cardValuesPart2 = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func (a ByHand) Len() int { return len(a) }
func (a ByHand) Less(i, j int) bool {
	iCards := a[i].cards
	jCards := a[j].cards

	iHandType := a[i].handType
	jHandType := a[j].handType

	if iHandType != jHandType {
		return iHandType > jHandType
	}

	for cardIndex := 0; cardIndex < 5; cardIndex++ {
		iCard := string(iCards[cardIndex])
		jCard := string(jCards[cardIndex])

		if iCard != jCard {
			if part == 2 {
				return cardValuesPart2[iCard] < cardValuesPart2[jCard]
			}
			return cardValues[iCard] < cardValues[jCard]
		}
	}
	panic(fmt.Sprintf("Cards shouldn't be equal: %v, %v", iCards, jCards))
}
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func partOne(lines []string, debug bool) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		values := strings.Fields(line)
		cards := values[0]
		bid, _ := strconv.Atoi(values[1])
		hands[i] = Hand{cards, calculateHandType(cards, debug), bid}
	}

	sort.Sort(ByHand(hands))
	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		winnings := hand.bid * rank
		totalWinnings += winnings
		if debug {
			fmt.Printf("Hand %v: %v, %v, %v, %v\n", rank, hand.cards, hand.handType, hand.bid, winnings)
		}
	}
	fmt.Printf("Total winnings: %v\n", totalWinnings)
}

func partTwo(lines []string, debug bool) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		values := strings.Fields(line)
		cards := values[0]
		bid, _ := strconv.Atoi(values[1])
		hands[i] = Hand{cards, calculateHandTypePart2(cards, debug), bid}
	}

	sort.Sort(ByHand(hands))
	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		winnings := hand.bid * rank
		totalWinnings += winnings
		if debug {
			fmt.Printf("Hand %v: %v, %v, %v, %v\n", rank, hand.cards, hand.handType, hand.bid, winnings)
		}
	}
	fmt.Printf("Total winnings: %v\n", totalWinnings)
}

func main() {
	day := 7
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	debug := common.IsTestMode(os.Args)
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	if part == 1 {
		partOne(lines, debug)
	} else {
		partTwo(lines, debug)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
