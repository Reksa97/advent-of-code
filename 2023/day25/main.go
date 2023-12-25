package main

import (
	"aoc2023/common"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/awalterschulze/gographviz"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func createGraph(lines []string) map[string]map[string]bool {
	edges := make(map[string]map[string]bool)

	for _, line := range lines {
		split := strings.Split(line, ": ")
		name := split[0]
		if _, ok := edges[name]; !ok {
			edges[name] = make(map[string]bool)
		}
		connections := strings.Split(split[1], " ")
		for _, connection := range connections {
			edges[name][connection] = true
			if _, ok := edges[connection]; !ok {
				edges[connection] = make(map[string]bool)
			}
			edges[connection][name] = true
		}
	}

	return edges
}

type Edge struct {
	firstNode  string
	secondNode string
}

func getAllEdges(edges map[string]map[string]bool) []Edge {
	var result []Edge
	alreadyAdded := make(map[string]bool)
	for firstNode := range edges {
		for secondNode := range edges[firstNode] {
			if !alreadyAdded[firstNode+secondNode] && !alreadyAdded[secondNode+firstNode] {
				result = append(result, Edge{firstNode: firstNode, secondNode: secondNode})
				alreadyAdded[firstNode+secondNode] = true
			}
		}
	}
	return result
}

func findConnected(edges map[string]map[string]bool, startNode string) int {
	visited := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, startNode)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if visited[currentNode] {
			continue
		}
		visited[currentNode] = true
		for node, isConnected := range edges[currentNode] {
			if isConnected {
				queue = append(queue, node)
			}
		}
	}
	return len(visited)
}

func printGraph(edges map[string]map[string]bool, edgesInArray []Edge) {
	g := gographviz.NewGraph()
	if err := g.SetName("G"); err != nil {
		panic(err)
	}
	/* if err := g.SetDir(true); err != nil {
		panic(err)
	} */

	for node := range edges {
		g.AddNode("G", node, nil)
	}

	for _, edge := range edgesInArray {
		g.AddEdge(edge.firstNode, edge.secondNode, false, nil)
	}

	s := g.String()
	if debug {
		fmt.Println(s)
	}

	// Create the output file
	outFile := "day25/graph.dot"
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Write to the output file
	_, err = f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the output image
	cmd := exec.Command("dot", "-Tsvg", outFile, "-o", "day25/graph.svg", "-Kneato")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func partOne(lines []string) {
	edges := createGraph(lines)

	nodesCount := len(edges)

	edgesInArray := getAllEdges(edges)
	if debug {
		fmt.Println(edgesInArray)
	}

	printGraph(edges, edgesInArray)

	if !debug {
		fmt.Println("Warning: running this with real input will take almost forever, check the graph.svg file instead")
		edgesInArray = []Edge{
			{firstNode: "gtj", secondNode: "tzj"},
			{firstNode: "jzv", secondNode: "qvq"},
			{firstNode: "bbp", secondNode: "dvr"},
		}
	}

	for i, firstEdge := range edgesInArray {
		for j, secondEdge := range edgesInArray {
			if i == j || j < i {
				continue
			}
			for k, thirdEdge := range edgesInArray {
				if i == k || j == k {
					continue
				}
				if k < j {
					continue
				}
				// remove edges
				edges[firstEdge.firstNode][firstEdge.secondNode] = false
				edges[firstEdge.secondNode][firstEdge.firstNode] = false
				edges[secondEdge.firstNode][secondEdge.secondNode] = false
				edges[secondEdge.secondNode][secondEdge.firstNode] = false
				edges[thirdEdge.firstNode][thirdEdge.secondNode] = false
				edges[thirdEdge.secondNode][thirdEdge.firstNode] = false

				// check if graph is connected
				connectedCount := findConnected(edges, firstEdge.firstNode)
				if connectedCount != nodesCount {
					groupOneSize := connectedCount
					groupTwoSize := nodesCount - connectedCount
					fmt.Printf("Group 1 size: %v, Group 2 size: %v\n", groupOneSize, groupTwoSize)
					fmt.Printf("Product: %v\n", groupOneSize*groupTwoSize)
					return
				}

				// reset back to original
				edges[firstEdge.firstNode][firstEdge.secondNode] = true
				edges[firstEdge.secondNode][firstEdge.firstNode] = true
				edges[secondEdge.firstNode][secondEdge.secondNode] = true
				edges[secondEdge.secondNode][secondEdge.firstNode] = true
				edges[thirdEdge.firstNode][thirdEdge.secondNode] = true
				edges[thirdEdge.secondNode][thirdEdge.firstNode] = true
			}
		}
	}
	panic("No solution found")
}

func partTwo(lines []string) {

}

func main() {
	day := 25
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
