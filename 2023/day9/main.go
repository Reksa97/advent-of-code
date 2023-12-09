package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strings"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func solve(lines []string) {
	sumOfNextValues := 0
	for _, history := range lines {
		differences := [][]int{}
		historyValuesString := strings.Fields(history)
		historyValues := common.ConvertToInt(historyValuesString)
		if part == 2 {
			// reverse array
			for i, j := 0, len(historyValues)-1; i < j; i, j = i+1, j-1 {
				historyValues[i], historyValues[j] = historyValues[j], historyValues[i]
			}
		}
		values := historyValues
		historyLength := len(values)
		for i := 0; i < historyLength; i++ {
			if debug {
				fmt.Printf("i: %v/%v\n", i, len(values))
			}
			differences = append(differences, []int{})
			everythingIsZero := true
			previous := values[0]
			for _, value := range values[1:] {
				difference := value - previous
				differences[i] = append(differences[i], difference)
				if difference != 0 {
					everythingIsZero = false
				}
				previous = value
			}
			if everythingIsZero {
				if debug {
					fmt.Printf("everything is zero, differences[%v]: %v\n", i, differences[i])
				}
				break
			}
			values = differences[i]
		}
		if debug {
			fmt.Printf("differences: %v\n", differences)
		}

		lastIndex := len(differences) - 1
		increment := differences[lastIndex][len(differences[lastIndex])-1]
		for i := len(differences) - 2; i >= 0; i-- {
			increment += differences[i][len(differences[i])-1]
			if debug {
				fmt.Printf("i: %v, increment: %v\n", i, increment)
			}
		}
		nextValue := historyValues[len(historyValues)-1] + increment
		sumOfNextValues += nextValue
		if debug {
			fmt.Printf("sumOfNextValues: %v (+%v)\n", sumOfNextValues, nextValue)
		}
	}
	fmt.Printf("Answer: %v\n", sumOfNextValues)
}

func main() {
	day := 9
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	solve(lines)

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
