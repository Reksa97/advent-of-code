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

type Pulse struct {
	source string
	target string
	isHigh bool
}

type FlipFlop struct {
	name    string
	targets []string
	isOn    bool
}

type Conjunction struct {
	name    string
	targets []string
	inputs  map[string]bool
}

func parseInput(lines []string) ([]string, map[string]FlipFlop, map[string]Conjunction) {
	broadcaster := make([]string, 0)
	flipFlops := make(map[string]FlipFlop)
	conjunctions := make(map[string]Conjunction)
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		source := split[0]
		targets := strings.Split(split[1], ", ")
		if debug {
			fmt.Println(source, targets)
		}
		if source == "broadcaster" {
			broadcaster = append(broadcaster, targets...)
			continue
		}
		switch source[0] {
		case '%':
			if debug {
				fmt.Println("Flip-flop", source[1:], "to", targets)
			}
			name := source[1:]
			flipFlops[name] = FlipFlop{name, targets, false}
		case '&':
			if debug {
				fmt.Println("Conjunction", source[1:], "to", targets)
			}
			name := source[1:]
			conjunctions[name] = Conjunction{name, targets, make(map[string]bool)}
		}
	}
	for _, flipFlop := range flipFlops {
		for _, target := range flipFlop.targets {
			if _, ok := conjunctions[target]; ok {
				// flip-flop is an input to conjunction
				conjunctions[target].inputs[flipFlop.name] = false
			}
		}
	}
	for _, conjunction := range conjunctions {
		for _, target := range conjunction.targets {
			if _, ok := conjunctions[target]; ok {
				// conjunction is an input to conjunction
				conjunctions[target].inputs[conjunction.name] = false
			}
		}
	}

	if debug {
		fmt.Println("\nBroadcaster", broadcaster)
		fmt.Println("Flip-flops", flipFlops)
		fmt.Println("Conjunctions", conjunctions)
		fmt.Println("Low pulses", totalLowPulses)
		fmt.Println("High pulses", totalHighPulses)
		fmt.Println()
	}
	return broadcaster, flipFlops, conjunctions
}

func pressButton(broadcaster []string, flipFlops map[string]FlipFlop, conjunctions map[string]Conjunction, findCyclesFor map[string]bool) map[string]bool {
	queue := make([]Pulse, 0)
	// First button sends low pulse to broadcaster
	totalLowPulses++
	for _, target := range broadcaster {
		queue = append(queue, Pulse{"broadcaster", target, false})
	}
	if debug {
		fmt.Println("Queue", queue)
	}
	for len(queue) > 0 {
		pulse := queue[0]
		queue = queue[1:]
		if debug {
			fmt.Println("Pulse", pulse)
		}
		if pulse.isHigh {
			if debug {
				fmt.Println("Send -high to", pulse.target)
			}
			totalHighPulses++
			if pulse.target == "dr" {
				if _, ok := findCyclesFor[pulse.source]; ok {
					findCyclesFor[pulse.source] = true
					//fmt.Println("Cycle found for", pulse.source, "send -high to", pulse.target, findCyclesFor)
				}
			}

		} else {
			if debug {
				fmt.Println("Send -low to", pulse.target)
			}
			totalLowPulses++
		}
		if !pulse.isHigh {
			if flipFlop, ok := flipFlops[pulse.target]; ok {
				if debug {
					fmt.Println("Flip-flop", flipFlop)
				}
				flipFlop.isOn = !flipFlop.isOn
				flipFlops[pulse.target] = flipFlop
				for _, target := range flipFlop.targets {
					queue = append(queue, Pulse{flipFlop.name, target, flipFlop.isOn})
				}
				continue
			}
		}

		if conjunction, ok := conjunctions[pulse.target]; ok {
			if debug {
				fmt.Println("Conjunction", conjunction)
			}

			if pulse.source != "broadcaster" {
				conjunction.inputs[pulse.source] = pulse.isHigh
			}

			allInputsHigh := true
			for _, input := range conjunction.inputs {
				if !input {
					allInputsHigh = false
					break
				}
			}

			for _, target := range conjunction.targets {
				queue = append(queue, Pulse{conjunction.name, target, !allInputsHigh})
			}
		}
	}
	if debug {
		fmt.Println()
		fmt.Println("Conjunctions", conjunctions)
		fmt.Println("flip-flops", flipFlops)
		fmt.Println("Low pulses", totalLowPulses)
		fmt.Println("High pulses", totalHighPulses)
		fmt.Println()
	}
	return findCyclesFor
}

var totalLowPulses = 0
var totalHighPulses = 0

func partOne(lines []string) {
	broadcaster, flipFlops, conjunctions := parseInput(lines)
	for i := 0; i < 1000; i++ {
		pressButton(broadcaster, flipFlops, conjunctions, make(map[string]bool))
	}
	answer := totalHighPulses * totalLowPulses
	fmt.Printf("Total low pulses: %v\n", totalLowPulses)
	fmt.Printf("Total high pulses: %v\n", totalHighPulses)

	fmt.Printf("Answer: %v\n", answer)
}

type Cycle struct {
	previousIndex int
	length        int
}

func partTwo(lines []string) {
	broadcaster, flipFlops, conjunctions := parseInput(lines)

	levels := 4
	previousState := make([][]int, levels)
	rxLevels := make([][]string, levels)
	rxLevels[0] = []string{"rx"}

	for level := 1; level < levels; level++ {
		rxLevels[level] = make([]string, 0)
		previousState[level] = make([]int, 0)
		for _, conjunction := range conjunctions {
			for _, target := range conjunction.targets {
				for _, levelTarget := range rxLevels[level-1] {
					if target == levelTarget {
						rxLevels[level] = append(rxLevels[level], conjunction.name)
						previousState[level] = append(previousState[level], 0)
					}
				}
			}
		}
		for _, flipFlop := range flipFlops {
			for _, target := range flipFlop.targets {
				for _, levelTarget := range rxLevels[level-1] {
					if target == levelTarget {
						fmt.Printf("Flip-flop %v found on level %v\n", flipFlop.name, level)
						panic("Wrong assumption: Flip-flop found on level")
					}
				}
			}
		}
	}

	if len(rxLevels[1]) != 1 || len(rxLevels[2]) != 4 || len(rxLevels[3]) != 4 {
		fmt.Println(rxLevels)
		panic("Wrong assumptions about conjunctions leading to rx")
	}

	findCyclesFor := make(map[string]bool)
	cycles := make(map[string]Cycle)
	for _, conjunction := range rxLevels[2] {
		findCyclesFor[conjunction] = false
		cycles[conjunction] = Cycle{0, 0}
	}

	fmt.Println(rxLevels, findCyclesFor)

	firstAllFoundIndex := 0
	for i := 1; true; i++ {
		ccFindCyclesFor := make(map[string]bool)
		for conjunction := range findCyclesFor {
			ccFindCyclesFor[conjunction] = false
		}
		foundCyclesFor := pressButton(broadcaster, flipFlops, conjunctions, ccFindCyclesFor)
		for conjunction, found := range foundCyclesFor {
			if found {
				if cycles[conjunction].length == 0 {
					fmt.Println("Cycle found for", conjunction, "at", i, "length", i)
					cycles[conjunction] = Cycle{i, i}
				} else if cycles[conjunction].previousIndex != i-cycles[conjunction].length {
					fmt.Println("Cycle changed for", conjunction, "at", i, "length", i-cycles[conjunction].previousIndex)
					cycles[conjunction] = Cycle{i, i - cycles[conjunction].previousIndex}
					panic("Cycle changed")
				} else {
					cycles[conjunction] = Cycle{i, cycles[conjunction].length}
				}
			}
		}
		allFound := true

		for _, cycle := range cycles {
			if cycle.length == 0 {
				allFound = false
				break
			}
		}

		if allFound && firstAllFoundIndex == 0 {
			firstAllFoundIndex = i
			fmt.Println("All cycles found at", i)
		}

		if firstAllFoundIndex > 0 && i > firstAllFoundIndex*4 {
			fmt.Println("Just made sure the loops stay as they were... at", i)
			break
		}
	}
	cyclesAsIntArray := make([]int, 0)
	for _, cycle := range cycles {
		cyclesAsIntArray = append(cyclesAsIntArray, cycle.length)
	}
	answer := common.LcmArray(cyclesAsIntArray)

	fmt.Printf("Answer: %v\n", answer)
}

func main() {
	day := 20
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
