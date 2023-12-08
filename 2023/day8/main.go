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

func parseInput(lines []string) (string, map[string]Node, []string) {
	instructions := lines[0]
	nodesMap := map[string]Node{}
	var startNodes []string
	nodes := lines[2:]
	for _, node := range nodes {
		nodeParts := strings.Split(node, " ")
		left := nodeParts[2][1 : len(nodeParts[2])-1]
		right := nodeParts[3][0 : len(nodeParts[3])-1]
		name := nodeParts[0]
		nodesMap[nodeParts[0]] = Node{left: left, right: right, name: name}
		if name[len(name)-1] == 'Z' {
			startNodes = append(startNodes, name)
		}
	}
	return instructions, nodesMap, startNodes
}

func partOne(lines []string) {
	instructions, nodesMap, _ := parseInput(lines)

	currentNode := "AAA"

	steps := 0
	for i := 1; true; i++ {
		if debug {
			fmt.Printf("Iteration: %v\n", i)
		}
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

// Function to calculate GCD
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate LCM of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate LCM of an array of integers
func lcmArray(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func partTwo(lines []string) {
	instructions, nodesMap, nodes := parseInput(lines)

	steps := 0
	for i := 1; true; i++ {
		if debug {
			fmt.Printf("Iteration: %v\n", i)
		} else if i%1000000 == 0 {
			fmt.Printf("Iteration: %v\n", i)
		}
		for _, instruction := range instructions {
			steps++
			for nodeIndex, node := range nodes {
				if debug {
					fmt.Printf("nodeIndex: %v, node: %v, instruction: %v\n", nodeIndex, node, instruction)
				}
				switch instruction {
				case 'L':
					if debug {
						fmt.Println("left")
					}
					nodes[nodeIndex] = nodesMap[node].left
				case 'R':
					if debug {
						fmt.Println("right")
					}
					nodes[nodeIndex] = nodesMap[node].right
				default:
					fmt.Println("error")
				}
			}
		}

		everythingInItsRightPlace := true
		for _, node := range nodes {
			if node[len(node)-1] != 'Z' {
				everythingInItsRightPlace = false
			}
		}
		if everythingInItsRightPlace {
			fmt.Printf("Found it! Iteration: %v, steps: %v\n", i, steps)
			break
		}
	}
}

func partTwoOptimised(lines []string) {
	instructions, nodesMap, nodes := parseInput(lines)

	nodeIterations := make([]int, len(nodes))
	for nodeIndex, node := range nodes {
		currentNode := node
		steps := 0
		for i := 1; true; i++ {
			if debug {
				fmt.Printf("Iteration: %v\n", i)
			}
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
			if currentNode[len(currentNode)-1] == 'Z' {
				fmt.Printf("Found it! Node %v, Iteration: %v, steps: %v, %v\n", node, i, steps, steps/i)
				nodeIterations[nodeIndex] = i
				break
			}
		}
	}
	fmt.Printf("nodeIterations: %v\n", nodeIterations)
	leastCommonMultiple := lcmArray(nodeIterations)
	fmt.Printf("LCM: %v\n", leastCommonMultiple)
	fmt.Printf("Steps: %v\n", leastCommonMultiple*len(instructions))
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
		partTwoOptimised(lines)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
