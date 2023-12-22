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

func partOne(lines []string) {
	bricks := make([]Brick, 0)
	for _, line := range lines {
		split := strings.Split(line, "~")
		startSplit := common.ConvertToInt(strings.Split(split[0], ","))
		endSplit := common.ConvertToInt(strings.Split(split[1], ","))
		start := Coordinate{startSplit[0], startSplit[1], startSplit[2]}
		end := Coordinate{endSplit[0], endSplit[1], endSplit[2]}
		brick := Brick{start, end, false, make([]int, 0), make([]int, 0)}
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
	for i, brick := range bricks {
		if debug {
			fmt.Println("go through bricks", i, brick)
		}
		if len(brick.directlyAbove) == 0 {
			couldBeRemovedCount++
			if debug {
				fmt.Println("could be removed, no above bricks", i, brick)
			}
		} else {
			allAboveBricksHaveAnotherSupport := true
			for _, aboveIndex := range brick.directlyAbove {
				hasAnotherSupport := false
				for _, belowIndex := range bricks[aboveIndex].directlyBelow {
					if belowIndex != i {
						hasAnotherSupport = true
						break
					}
				}

				if !hasAnotherSupport {
					allAboveBricksHaveAnotherSupport = false
					break
				}
			}
			if allAboveBricksHaveAnotherSupport {
				couldBeRemovedCount++
				if debug {
					fmt.Println("could be removed, all above have support", i, brick)
				}
			}
		}
	}
	fmt.Println("couldBeRemovedCount", couldBeRemovedCount)
	if debug {
		fmt.Println("bricks = [")
		for _, brick := range bricks {
			fmt.Printf("    [(%v, %v, %v), (%v, %v, %v), np.random.rand(3,)],\n", brick.start.x, brick.start.y, brick.start.z, brick.end.x, brick.end.y, brick.end.z)
		}
		fmt.Println("]")
	}
}

func partTwo(lines []string) {

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

	if part == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
