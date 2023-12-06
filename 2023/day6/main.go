package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput(lines []string, debug bool) ([]int, []int) {
	timeValues := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
	distanceValues := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))
	if debug {
		fmt.Printf("Time: %v\n", timeValues)
		fmt.Printf("Distance: %v\n", distanceValues)
	}

	return common.ConvertToInt(timeValues), common.ConvertToInt(distanceValues)
}

func partOne(lines []string, debug bool) {
	timeValues, distanceValues := parseInput(lines, debug)

	result := 1
	for i := 0; i < len(timeValues); i++ {
		recordDistance := distanceValues[i]
		if debug {
			fmt.Printf("Race: %v, Time: %vms, Record distance: %vmm\n", i+1, timeValues[i], recordDistance)
		}
		recordWinningStrategiesCount := 0
		for holdButtonForMs := 1; holdButtonForMs < timeValues[i]; holdButtonForMs++ {
			speedMmPerMs := holdButtonForMs
			timeToTravel := timeValues[i] - holdButtonForMs
			distanceTravelled := speedMmPerMs * timeToTravel

			if distanceTravelled > recordDistance {
				recordWinningStrategiesCount++
				if debug {
					fmt.Printf("Hold button for %vms, speed: %vmm/ms, time to travel: %vms, distance travelled: %vmm\n", holdButtonForMs, speedMmPerMs, timeToTravel, distanceTravelled)
				}
			}
		}
		result *= recordWinningStrategiesCount
	}
	fmt.Printf("Result: %v\n", result)
}

func partTwo(lines []string, debug bool) {

}

func main() {
	day := 6
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	part := common.GetPart(os.Args)
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
