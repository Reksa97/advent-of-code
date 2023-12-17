package main

import (
	"aoc2023/common"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

type Path struct {
	fromX            int
	fromY            int
	toX              int
	toY              int
	direction        string
	currentHeatLoss  int
	straightDistance int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    Path
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value Path, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func leastHeatLoss(heatLosses []int) int {
	leastHeatLoss := heatLosses[0]
	for _, heatLoss := range heatLosses {
		if heatLoss > 0 && (leastHeatLoss == -1 || heatLoss < leastHeatLoss) {
			leastHeatLoss = heatLoss
		}
	}
	return leastHeatLoss
}

func findPathWithLeastHeatLoss(pq *PriorityQueue, lines []string, smallestHeatLoss [][]map[string]int, fromX int, fromY int, toX int, toY int, direction string, currentHeatLoss int, straightDistance int) int {
	if fromX < 0 || fromY < 0 || fromY >= len(lines) || fromX >= len(lines[fromY]) {
		return -1
	}
	heatLossHere, _ := strconv.Atoi(string(lines[fromY][fromX]))
	heatLossOnPath := heatLossHere + currentHeatLoss

	key := fmt.Sprintf("%v,%v", direction, straightDistance)

	if smallestHeatLoss[fromY][fromX][key] != 0 && smallestHeatLoss[fromY][fromX][key] <= heatLossOnPath {
		return -1
	}

	smallestHeatLoss[fromY][fromX][key] = heatLossOnPath

	if fromX == toX && fromY == toY {
		return heatLossOnPath
	}

	switch direction {
	case "right":
		// try to go down
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX, fromY: fromY + 1, toX: toX, toY: toY, direction: "down", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		// try to go up
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX, fromY: fromY - 1, toX: toX, toY: toY, direction: "up", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		if straightDistance < 3 {
			// continue straight
			heap.Push(pq, &Item{
				value:    Path{fromX: fromX + 1, fromY: fromY, toX: toX, toY: toY, direction: "right", currentHeatLoss: heatLossOnPath, straightDistance: straightDistance + 1},
				priority: heatLossOnPath,
			})
		}
	case "left":
		// try to go down
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX, fromY: fromY + 1, toX: toX, toY: toY, direction: "down", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		// try to go up
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX, fromY: fromY - 1, toX: toX, toY: toY, direction: "up", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		if straightDistance < 3 {
			// continue straight
			heap.Push(pq, &Item{
				value:    Path{fromX: fromX - 1, fromY: fromY, toX: toX, toY: toY, direction: "left", currentHeatLoss: heatLossOnPath, straightDistance: straightDistance + 1},
				priority: heatLossOnPath,
			})
		}
	case "up":
		// try to go right
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX + 1, fromY: fromY, toX: toX, toY: toY, direction: "right", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		// try to go left
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX - 1, fromY: fromY, toX: toX, toY: toY, direction: "left", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		if straightDistance < 3 {
			// continue straight
			heap.Push(pq, &Item{
				value:    Path{fromX: fromX, fromY: fromY - 1, toX: toX, toY: toY, direction: "up", currentHeatLoss: heatLossOnPath, straightDistance: straightDistance + 1},
				priority: heatLossOnPath,
			})
		}
	case "down":
		// try to go right
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX + 1, fromY: fromY, toX: toX, toY: toY, direction: "right", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		// try to go left
		heap.Push(pq, &Item{
			value:    Path{fromX: fromX - 1, fromY: fromY, toX: toX, toY: toY, direction: "left", currentHeatLoss: heatLossOnPath, straightDistance: 1},
			priority: heatLossOnPath,
		})
		if straightDistance < 3 {
			// continue straight
			heap.Push(pq, &Item{
				value:    Path{fromX: fromX, fromY: fromY + 1, toX: toX, toY: toY, direction: "down", currentHeatLoss: heatLossOnPath, straightDistance: straightDistance + 1},
				priority: heatLossOnPath,
			})
		}
	default:
		panic(fmt.Sprintf("Unknown direction: %v", direction))
	}
	if debug {
		fmt.Printf("pq len: %v\n", pq.Len())
	}
	return -1
}

func partOne(lines []string) {
	smallestHeatLoss := make([][]map[string]int, len(lines))
	for i, line := range lines {
		if debug {
			fmt.Println(line)
		}
		smallestHeatLoss[i] = make([]map[string]int, len(line))
		for j := range smallestHeatLoss[i] {
			smallestHeatLoss[i][j] = make(map[string]int)
		}
	}

	pq := make(PriorityQueue, 2)

	pq[0] = &Item{
		value:    Path{fromX: 1, fromY: 0, toX: len(lines[0]) - 1, toY: len(lines) - 1, direction: "right", currentHeatLoss: 0, straightDistance: 1},
		priority: 0,
	}

	pq[1] = &Item{
		value:    Path{fromX: 0, fromY: 1, toX: len(lines[0]) - 1, toY: len(lines) - 1, direction: "down", currentHeatLoss: 0, straightDistance: 1},
		priority: 0,
	}

	heap.Init(&pq)

	lowestHeatLoss := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if debug {
			fmt.Printf("fromX: %v, fromY: %v, toX: %v, toY: %v, direction: %v, currentHeatLoss: %v, straightDistance: %v, pq len: %v\n", item.value.fromX, item.value.fromY, item.value.toX, item.value.toY, item.value.direction, item.value.currentHeatLoss, item.value.straightDistance, pq.Len())
		}

		heatLoss := findPathWithLeastHeatLoss(&pq, lines, smallestHeatLoss, item.value.fromX, item.value.fromY, item.value.toX, item.value.toY, item.value.direction, item.value.currentHeatLoss, item.value.straightDistance)
		if debug {
			fmt.Printf("pq len: %v\n", pq.Len())
		}
		if heatLoss > 0 && (lowestHeatLoss == 0 || heatLoss < lowestHeatLoss) {
			lowestHeatLoss = heatLoss
		}
	}
	fmt.Printf("Lowest heat loss: %v\n", lowestHeatLoss)
}

func partTwo(lines []string) {

}

func main() {
	day := 17
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
