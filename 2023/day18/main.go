package main

import (
	"aoc2023/common"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

type Vertex struct {
	X, Y float64
}

func ShoelaceArea(vertices []Vertex) int {
	n := len(vertices)
	area := 0.0

	for i := 0; i < n-1; i++ {
		area += vertices[i].X*vertices[i+1].Y - vertices[i+1].X*vertices[i].Y
	}

	// Closing the polygon
	area += vertices[n-1].X*vertices[0].Y - vertices[0].X*vertices[n-1].Y

	return int(math.Abs(area) / 2.0)
}

func solve(lines []string) {
	vertices := make([]Vertex, len(lines)+1)
	i := 0
	x := 0
	y := 0
	vertices[i] = Vertex{0.0, 0.0}
	edgeArea := 0
	for _, line := range lines {
		i++
		parts := strings.Fields(line)
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])
		if part == 2 {
			lastField := parts[2][1 : len(parts[2])-1]
			directionInt := lastField[len(lastField)-1:]
			hex := lastField[1 : len(lastField)-1]
			hexAsInt, _ := strconv.ParseInt(hex, 16, 64)
			distance = int(hexAsInt)
			if debug {
				fmt.Printf("Hex: %v (%v), direction: %v\n", hex, hexAsInt, directionInt)
			}
			switch directionInt {
			case "0":
				direction = "R"
			case "1":
				direction = "D"
			case "2":
				direction = "L"
			case "3":
				direction = "U"
			}
		}
		switch direction {
		case "U":
			y += distance
		case "D":
			y -= distance
		case "R":
			x += distance
		case "L":
			x -= distance
		default:
			panic("Unknown direction")
		}
		edgeArea += distance
		vertices[i] = Vertex{float64(x), float64(y)}
	}

	area := ShoelaceArea(vertices)

	interiorPoints := area - (edgeArea / 2) + 1
	if debug {
		fmt.Printf("Vertices: %v\n", vertices)
		fmt.Printf("Area: %v\n", area)
		fmt.Printf("Edge area: %v\n", edgeArea)
		fmt.Printf("Interior points: %v - (%v/2) + 1 = %v\n", area, edgeArea, interiorPoints)
	}
	fmt.Printf("Result: %v\n", interiorPoints+edgeArea)
}

func main() {
	day := 18
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
