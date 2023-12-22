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

type Coordinate struct {
	x int
	y int
	z int
}

type Brick struct {
	start         Coordinate
	end           Coordinate
	foundItsPlace bool
	directlyAbove []int
	directlyBelow []int
	deleted       bool
}

/* func isBlocking(fallingBrick Brick, otherBrick Brick) bool {
	return fallingBrick.start.x <= otherBrick.end.x &&
		fallingBrick.end.x >= otherBrick.start.x &&
		fallingBrick.start.y <= otherBrick.end.y &&
		fallingBrick.end.y >= otherBrick.start.y &&
		fallingBrick.start.z-1 == otherBrick.end.z
} */

func isBlocking(fallingBrick Brick, otherBrick Brick) bool {
	if fallingBrick.start.z-1 == otherBrick.end.z {
		for x := fallingBrick.start.x; x <= fallingBrick.end.x; x++ {
			for y := fallingBrick.start.y; y <= fallingBrick.end.y; y++ {
				if x >= otherBrick.start.x && x <= otherBrick.end.x && y >= otherBrick.start.y && y <= otherBrick.end.y {
					return true
				}
			}
		}
	}
	return false
}

func solve(lines []string) {
	bricks := make([]Brick, 0)
	for _, line := range lines {
		split := strings.Split(line, "~")
		startSplit := common.ConvertToInt(strings.Split(split[0], ","))
		endSplit := common.ConvertToInt(strings.Split(split[1], ","))
		start := Coordinate{startSplit[0], startSplit[1], startSplit[2]}
		end := Coordinate{endSplit[0], endSplit[1], endSplit[2]}
		brick := Brick{start, end, false, make([]int, 0), make([]int, 0), false}
		if start.x > end.x || start.y > end.y || start.z > end.z {
			panic("wrong expected input")
		}
		bricks = append(bricks, brick)
	}
	bricksAreFalling := true
	for bricksAreFalling {
		allBricksFoundTheirPlace := true
		for i, fallingBrick := range bricks {
			if fallingBrick.foundItsPlace {
				continue
			}
			falling := fallingBrick.start.z > 1
			if !falling {
				fallingBrick.foundItsPlace = true
				bricks[i] = fallingBrick
				continue
			}
			for j, otherBrick := range bricks {
				if i == j {
					continue
				}
				// if any coordinate of the falling brick is just over any coordinate in the other brick, it's not falling
				if isBlocking(fallingBrick, otherBrick) {
					falling = false // not falling in this round, below brick might still be falling
					if otherBrick.foundItsPlace {
						if debug {
							fmt.Println("found its place", fallingBrick, "i", i)
						}
						fallingBrick.foundItsPlace = true
						found := false
						for _, index := range otherBrick.directlyAbove {
							if index == i {
								found = true
								break
							}
						}
						if !found {
							otherBrick.directlyAbove = append(otherBrick.directlyAbove, i)
						}
						found = false
						for _, index := range fallingBrick.directlyBelow {
							if index == j {
								found = true
								break
							}
						}
						if !found {
							fallingBrick.directlyBelow = append(fallingBrick.directlyBelow, j)
						}
						bricks[i] = fallingBrick
						bricks[j] = otherBrick
					} else {
						allBricksFoundTheirPlace = false
					}
				}
			}
			if falling {
				if debug {
					fmt.Println("still falling", fallingBrick, i)
				}
				fallingBrick.start.z--
				fallingBrick.end.z--
				bricks[i] = fallingBrick
				if debug {
					fmt.Println("fell", bricks[i])
				}
				allBricksFoundTheirPlace = false
			}
		}
		bricksAreFalling = !allBricksFoundTheirPlace
		if debug {
			fmt.Println("bricksAreFalling", bricksAreFalling)
			for _, brick := range bricks {
				fmt.Println(brick)
			}
			fmt.Println()
		}
	}
	fmt.Println()
	couldBeRemovedCount := 0
	couldNotBeRemoved := make(map[int][]int, 0)
	for i, brick := range bricks {
		if len(brick.directlyAbove) == 0 {
			couldBeRemovedCount++
			if debug {
				fmt.Println("could be removed, no above bricks", i, brick)
			}
			continue
		}
		wouldFall := make([]int, 0)
		for _, aboveIndex := range brick.directlyAbove {
			hasAnotherSupport := false
			for _, belowIndex := range bricks[aboveIndex].directlyBelow {
				if belowIndex != i {
					hasAnotherSupport = true
					break
				}
			}

			if !hasAnotherSupport {
				wouldFall = append(wouldFall, aboveIndex)
			}
		}
		if len(wouldFall) == 0 {
			couldBeRemovedCount++
			if debug {
				fmt.Println("could be removed, all above have support", i, brick)
			}
		} else {
			couldNotBeRemoved[i] = wouldFall
			if debug {
				fmt.Println("could not be removed", i, brick, wouldFall)
			}
		}
	}
	if part == 1 {
		fmt.Println("couldBeRemovedCount", couldBeRemovedCount)
		return
	}

	sumOfFallenBricks := 0
	for index := range couldNotBeRemoved {
		copyBricks := make([]Brick, len(bricks))
		copy(copyBricks, bricks)
		copyBricks[index].deleted = true
		simulateRemovingBrick(copyBricks, index)
		wouldFall := -1
		for _, brick := range copyBricks {
			if brick.deleted {
				wouldFall++
			}
		}

		if debug {
			fmt.Println(index, "wouldFall", wouldFall)
		}
		sumOfFallenBricks += wouldFall
	}
	fmt.Println("sumOfFallenBricks", sumOfFallenBricks)
}

var cache = make(map[int]int)

func simulateRemovingBrick(bricks []Brick, index int) {
	fallen := make([]int, 0)
	for _, aboveIndex := range bricks[index].directlyAbove {
		if debug {
			fmt.Println("aboveWouldFallIndex", aboveIndex)
		}
		stillHasSupport := false
		for _, belowIndex := range bricks[aboveIndex].directlyBelow {
			if belowIndex != index && !bricks[belowIndex].deleted {
				stillHasSupport = true
				break
			}
		}
		if !stillHasSupport {
			bricks[aboveIndex].deleted = true
			fallen = append(fallen, aboveIndex)
		}
	}
	for _, aboveIndex := range fallen {
		simulateRemovingBrick(bricks, aboveIndex)
	}
}

func main() {
	day := 22
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
