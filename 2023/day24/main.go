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

type Hailstone struct {
	x         int
	y         int
	z         int
	xVelocity int
	yVelocity int
	zVelocity int
	slope     float64
	yAtX0     float64
}

func isInFuture(hailstone Hailstone, x float64, y float64) bool {
	if debug {
		fmt.Printf("isInFuture: %v, %v, %v\n", hailstone, x, y)
	}
	if hailstone.xVelocity > 0 {
		if float64(hailstone.x) > x {
			return false
		}
	}
	if hailstone.xVelocity < 0 {
		if float64(hailstone.x) < x {
			return false
		}
	}
	if hailstone.yVelocity > 0 {
		if float64(hailstone.y) > y {
			return false
		}
	}
	if hailstone.yVelocity < 0 {
		if float64(hailstone.y) < y {
			return false
		}
	}
	return true
}

func parseHailstone(line string) Hailstone {
	split := strings.Split(line, " @ ")
	coordinateSplit := strings.Split(split[0], ", ")
	velocitySplit := strings.Split(split[1], ", ")
	hailstone := Hailstone{
		x:         common.MustAtoi(coordinateSplit[0]),
		y:         common.MustAtoi(coordinateSplit[1]),
		z:         common.MustAtoi(coordinateSplit[2]),
		xVelocity: common.MustAtoi(velocitySplit[0]),
		yVelocity: common.MustAtoi(velocitySplit[1]),
		zVelocity: common.MustAtoi(velocitySplit[2]),
	}
	hailstone.slope = float64(hailstone.yVelocity) / float64(hailstone.xVelocity)
	hailstone.yAtX0 = float64(hailstone.y) - (hailstone.slope * float64(hailstone.x))
	return hailstone
}

func partOne(lines []string) {
	minXY := 200000000000000
	maxXY := 400000000000000
	if debug {
		minXY = 7
		maxXY = 27
	}
	hailstones := make([]Hailstone, len(lines))
	for i, line := range lines {
		hailstones[i] = parseHailstone(line)
		if debug {
			fmt.Println(hailstones)
		}
	}

	intersectsInArea := 0
	for i, hailstone := range hailstones {
		if debug {
			fmt.Printf("slope: %v\n", hailstone)
		}
		for j, otherHailstone := range hailstones {
			if j <= i {
				continue
			}
			if debug {
				fmt.Printf("compare: %v %v\n", i, j)
			}
			if hailstone.slope == otherHailstone.slope {
				if debug {
					fmt.Printf("Parallel lines\n")
				}
				continue
			}
			// a*x+c = b*x+d
			// x = (d-c)/(a-b)
			x := (otherHailstone.yAtX0 - hailstone.yAtX0) / (hailstone.slope - otherHailstone.slope)
			y := hailstone.slope*x + hailstone.yAtX0

			if isInFuture(hailstone, x, y) && isInFuture(otherHailstone, x, y) &&
				x > float64(minXY) && x < float64(maxXY) &&
				y > float64(minXY) && y < float64(maxXY) {
				if debug {
					fmt.Printf("Intersection in area at %v, %v\n", x, y)
				}
				intersectsInArea++
			} else if debug {
				if !isInFuture(hailstone, x, y) || !isInFuture(otherHailstone, x, y) {
					fmt.Printf("Intersection in past at %v, %v\n", x, y)
				} else {
					fmt.Printf("Intersection outside area at %v, %v\n", x, y)
				}
			}
		}
	}
	fmt.Printf("Intersections in area: %v\n", intersectsInArea)
}

func partTwo(lines []string) {
	fileContent := make([]string, 0)
	fileContent = append(fileContent, "var('x y z vx vy vz t1 t2 t3 ans')")
	for i, line := range lines[0:3] {
		hailstone := parseHailstone(line)
		eqIndex := i*3 + 1
		fileContent = append(fileContent, fmt.Sprintf("eq%v = x + vx * t%v == %v + (%v * t%v)", eqIndex, i+1, hailstone.x, hailstone.xVelocity, i+1))
		eqIndex++
		fileContent = append(fileContent, fmt.Sprintf("eq%v = y + vy * t%v == %v + (%v * t%v)", eqIndex, i+1, hailstone.y, hailstone.yVelocity, i+1))
		eqIndex++
		fileContent = append(fileContent, fmt.Sprintf("eq%v = z + vz * t%v == %v + (%v * t%v)", eqIndex, i+1, hailstone.z, hailstone.zVelocity, i+1))
	}
	fileContent = append(fileContent, "eq10 = ans == x + y + z")
	fileContent = append(fileContent, "print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9,eq10],x,y,z,vx,vy,vz,t1,t2,t3,ans))")

	file, err := os.Create("day24/part2.sage")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, str := range fileContent {
		_, err := file.WriteString(str + "\n") // Adding a newline for readability
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Run 'load(\"day24/part2.sage\")' in SageMath")
}

func main() {
	day := 24
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
