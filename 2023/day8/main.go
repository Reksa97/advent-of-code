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

type Node struct {
	left  string
	right string
	name  string
}

func parseInput(lines []string) (string, map[string]Node) {
	instructions := lines[0]
	nodesMap := map[string]Node{}
	nodes := lines[2:]
	for _, node := range nodes {
		nodeParts := strings.Split(node, " ")
		left := nodeParts[2][1 : len(nodeParts[2])-1]
		right := nodeParts[3][0 : len(nodeParts[3])-1]
		nodesMap[nodeParts[0]] = Node{left: left, right: right, name: nodeParts[0]}
	}
	return instructions, nodesMap
}

func partOne(lines []string) {
	instructions, nodesMap := parseInput(lines)

	currentNode := "AAA"

	steps := 0
	for i := 1; true; i++ {
		fmt.Printf("Iteration: %v\n", i)
		for _, instruction := range instructions {
			steps++
			if debug {
				fmt.Printf("currentNode: %v, instruction: %v\n", currentNode, instruction)
			}
			switch instruction {
			case 'L':
				if debug {
					fmt.Println("left")
				}
				currentNode = nodesMap[currentNode].left
			case 'R':
				if debug {
					fmt.Println("right")
				}
				currentNode = nodesMap[currentNode].right
			default:
				fmt.Println("error")
			}
		}
		if currentNode == "ZZZ" {
			fmt.Printf("Found it! Iteration: %v, steps: %v\n", i, steps)
			break
		}
	}

}

func partTwo(lines []string) {

}

func main() {
	day := 8
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
