package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func parseInput(lines []string) ([]string, int, int) {
	newLines := make([]string, 0)
	for _, line := range lines {
		newLines = append(newLines, line)
	}
	for y, line := range newLines {
		for x, char := range line {
			if char == 'S' {
				newLines = append(newLines[:y], append([]string{line[:x] + "." + line[x+1:]}, newLines[y+1:]...)...)
				return newLines, x, y
			}
		}
	}
	panic("Start not found")
}

type Coordinate struct {
	x int
	y int
}

type MapShard struct {
	x int
	y int
}

func partOne(originalLines []string) {
	steps := 64
	lines, startX, startY := parseInput(originalLines)
	if debug {
		fmt.Println(startX, startY)
		for _, line := range lines {
			fmt.Println(line)
		}
		for _, line := range originalLines {
			fmt.Println(line)
		}
	}

	possiblePlaces := make(map[Coordinate]bool)
	possiblePlaces[Coordinate{startX, startY}] = true
	for step := 1; step <= steps; step++ {
		nextPlaces := make(map[Coordinate]bool, 0)
		for coordinate := range possiblePlaces {
			// South
			if coordinate.y+1 < len(lines) && lines[coordinate.y+1][coordinate.x] == '.' {
				nextPlaces[Coordinate{coordinate.x, coordinate.y + 1}] = true
			}
			// North
			if coordinate.y-1 >= 0 && lines[coordinate.y-1][coordinate.x] == '.' {
				nextPlaces[Coordinate{coordinate.x, coordinate.y - 1}] = true
			}
			// East
			if coordinate.x+1 < len(lines[coordinate.y]) && lines[coordinate.y][coordinate.x+1] == '.' {
				nextPlaces[Coordinate{coordinate.x + 1, coordinate.y}] = true
			}
			// West
			if coordinate.x-1 >= 0 && lines[coordinate.y][coordinate.x-1] == '.' {
				nextPlaces[Coordinate{coordinate.x - 1, coordinate.y}] = true
			}
		}
		if debug {
			fmt.Println("\nStep", step)
			for y, line := range lines {
				for x, char := range line {
					if _, ok := nextPlaces[Coordinate{x, y}]; ok {
						fmt.Print("O")
					} else {
						fmt.Print(string(char))
					}
				}
				fmt.Println()
			}
		}
		possiblePlaces = nextPlaces
	}

	fmt.Println("Possible plots:", len(possiblePlaces))

}

func partTwoFindThreeValuesFromQuadraticFunction(originalLines []string) {
	steps := 26501365
	lines, startX, startY := parseInput(originalLines)
	if debug {
		fmt.Println(startX, startY)
		for _, line := range lines {
			fmt.Println(line)
		}
		for _, line := range originalLines {
			fmt.Println(line)
		}
	}

	possiblePlaces := make(map[MapShard]map[Coordinate]bool)
	possiblePlaces[MapShard{0, 0}] = make(map[Coordinate]bool)
	possiblePlaces[MapShard{0, 0}][Coordinate{startX, startY}] = true
	startTime := time.Now()
	possiblePlacesCounts := make([]int, 0)
	for step := 0; step <= steps; step++ {
		nextPlaces := make(map[MapShard]map[Coordinate]bool, 0)
		totalPossiblePlacesCount := 0
		for mapShard := range possiblePlaces {
			for coordinate := range possiblePlaces[mapShard] {
				totalPossiblePlacesCount++
				if _, ok := nextPlaces[mapShard]; !ok {
					nextPlaces[mapShard] = make(map[Coordinate]bool)
				}
				// South
				if coordinate.y+1 < len(lines) && lines[coordinate.y+1][coordinate.x] == '.' {
					nextPlaces[mapShard][Coordinate{coordinate.x, coordinate.y + 1}] = true
				} else if coordinate.y+1 >= len(lines) && lines[0][coordinate.x] == '.' {
					// Add new map shard if not already added
					if _, ok := nextPlaces[MapShard{mapShard.x, mapShard.y + 1}]; !ok {
						nextPlaces[MapShard{mapShard.x, mapShard.y + 1}] = make(map[Coordinate]bool)
					}
					nextPlaces[MapShard{mapShard.x, mapShard.y + 1}][Coordinate{coordinate.x, 0}] = true
				}
				// North
				if coordinate.y-1 >= 0 && lines[coordinate.y-1][coordinate.x] == '.' {
					nextPlaces[mapShard][Coordinate{coordinate.x, coordinate.y - 1}] = true
				} else if coordinate.y-1 < 0 && lines[len(lines)-1][coordinate.x] == '.' {
					// Add new map shard if not already added
					if _, ok := nextPlaces[MapShard{mapShard.x, mapShard.y - 1}]; !ok {
						nextPlaces[MapShard{mapShard.x, mapShard.y - 1}] = make(map[Coordinate]bool)
					}
					nextPlaces[MapShard{mapShard.x, mapShard.y - 1}][Coordinate{coordinate.x, len(lines) - 1}] = true
				}
				// East
				if coordinate.x+1 < len(lines[coordinate.y]) && lines[coordinate.y][coordinate.x+1] == '.' {
					nextPlaces[mapShard][Coordinate{coordinate.x + 1, coordinate.y}] = true
				} else if coordinate.x+1 >= len(lines[coordinate.y]) && lines[coordinate.y][0] == '.' {
					// Add new map shard if not already added
					if _, ok := nextPlaces[MapShard{mapShard.x + 1, mapShard.y}]; !ok {
						nextPlaces[MapShard{mapShard.x + 1, mapShard.y}] = make(map[Coordinate]bool)
					}
					nextPlaces[MapShard{mapShard.x + 1, mapShard.y}][Coordinate{0, coordinate.y}] = true
				}
				// West
				if coordinate.x-1 >= 0 && lines[coordinate.y][coordinate.x-1] == '.' {
					nextPlaces[mapShard][Coordinate{coordinate.x - 1, coordinate.y}] = true
				} else if coordinate.x-1 < 0 && lines[coordinate.y][len(lines[coordinate.y])-1] == '.' {
					// Add new map shard if not already added
					if _, ok := nextPlaces[MapShard{mapShard.x - 1, mapShard.y}]; !ok {
						nextPlaces[MapShard{mapShard.x - 1, mapShard.y}] = make(map[Coordinate]bool)
					}
					nextPlaces[MapShard{mapShard.x - 1, mapShard.y}][Coordinate{len(lines[coordinate.y]) - 1, coordinate.y}] = true
				}
			}
		}
		if step == (65 + 131*2) {
			stepTime := time.Now().Sub(startTime)
			timeUntilReady := int(stepTime.Microseconds()) * (steps - step)
			if timeUntilReady > 10000000 {
				fmt.Printf("Calculating step %v took %v ms, if the rest took as long, it would take %v minutes until answer was brute forced (actually time per step will grow exponentially so it would probably never finish)\n", step, stepTime.Milliseconds(), timeUntilReady/1000000/60)
			}
		}

		// 65 steps to reach the border of the first map shard
		// 131 steps from there to reach the border of the second map shard
		if step == 65 || step == (65+131) || step == (65+131*2) {
			possiblePlacesCounts = append(possiblePlacesCounts, totalPossiblePlacesCount)
			if step == 327 {
				break
			}
		}
		possiblePlaces = nextPlaces
		startTime = time.Now()
	}
	fmt.Println()
	fmt.Println("Go to Wolfram Alpha and search for quadratic fit calculator")
	fmt.Println("https://www.wolframalpha.com/input?i=quadratic+fit+calculator")
	fmt.Println("data set of x values: {0, 1, 2}")
	fmt.Printf("data set of y values: {%v, %v, %v}\n", possiblePlacesCounts[0], possiblePlacesCounts[1], possiblePlacesCounts[2])
	// steps = 65 + 131 * x
	// x = (steps - 65) / 131
	useX := (steps - 65) / 131
	fmt.Printf("Click 'Least-squares best fit' to open it and add ' where x = %v' at the end of the input\n\n", useX)
}

func main() {
	day := 21
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	if part == 1 {
		partOne(lines)
	} else {
		partTwoFindThreeValuesFromQuadraticFunction(lines)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
