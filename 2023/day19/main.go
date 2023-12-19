package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

type Rating struct {
	x int
	m int
	a int
	s int
}

type Rule struct {
	variable              string
	operator              string
	value                 int
	destination           string
	alwaysGoToDestination bool
}

type RatingRange struct {
	minX int
	maxX int
	minM int
	maxM int
	minA int
	maxA int
	minS int
	maxS int
}

func isAccepted(rating Rating, workflows map[string][]Rule) bool {
	currentWorkflowName := "in"
	for true {
		if debug {
			fmt.Printf("Current workflow: %v\n", currentWorkflowName)
		}
		if currentWorkflowName == "A" {
			if debug {
				fmt.Println("Accepted!")
			}
			return true
		}
		if currentWorkflowName == "R" {
			if debug {
				fmt.Println("Rejected!")
			}
			return false
		}
		for _, rule := range workflows[currentWorkflowName] {
			if debug {
				fmt.Printf("Rule: %v\n", rule)
			}
			if rule.alwaysGoToDestination {
				currentWorkflowName = rule.destination
				break
			}
			value := 0
			switch rule.variable {
			case "x":
				value = rating.x
			case "m":
				value = rating.m
			case "a":
				value = rating.a
			case "s":
				value = rating.s
			}

			foundTrue := false
			switch rule.operator {
			case "<":
				if debug {
					fmt.Printf("Value: %v < rule.value: %v = %v\n", value, rule.value, value < rule.value)
				}
				if value < rule.value {
					currentWorkflowName = rule.destination
					foundTrue = true
					if debug {
						fmt.Printf("New workflow: %v\n", currentWorkflowName)
					}
					break
				}
			case ">":
				if debug {
					fmt.Printf("Value: %v > rule.value: %v = %v\n", value, rule.value, value > rule.value)
				}
				if value > rule.value {
					if debug {
						fmt.Printf("New workflow: %v\n", rule.destination)
					}
					foundTrue = true
					currentWorkflowName = rule.destination
					break
				}
			default:
				panic(fmt.Sprintf("Unknown operator: %v", rule.operator))
			}
			if foundTrue {
				break
			}
		}
	}
	panic("Not implemented")
}

func parseRule(ruleString string) Rule {
	parts := strings.Split(ruleString, ":")
	rule := Rule{}

	rule.variable = parts[0][0:1]
	rule.operator = parts[0][1:2]
	rule.value, _ = strconv.Atoi(parts[0][2:])
	rule.destination = parts[1]

	return rule
}

func parseWorkflows(lines []string) (map[string][]Rule, int) {
	workflows := make(map[string][]Rule)
	ratingStartIndex := 0
	for i, workflowLine := range lines {
		if workflowLine == "" {
			ratingStartIndex = i + 1
			break
		}
		workflowName := ""
		i := 0
		for true {
			if workflowLine[i] == '{' {
				i++
				break
			}
			workflowName += string(workflowLine[i])
			i++
		}

		workflows[workflowName] = make([]Rule, 0)

		currentRule := ""
		for true {
			if workflowLine[i] == '}' {
				workflow := workflows[workflowName]
				workflow = append(workflow, Rule{destination: currentRule, alwaysGoToDestination: true})
				workflows[workflowName] = workflow
				break
			}
			if workflowLine[i] == ',' {
				workflow := workflows[workflowName]
				workflow = append(workflow, parseRule(currentRule))
				workflows[workflowName] = workflow
				i++
				currentRule = ""
			}
			currentRule += string(workflowLine[i])
			i++
		}
	}
	return workflows, ratingStartIndex
}

func partOne(lines []string) {
	workflows, ratingStartIndex := parseWorkflows(lines)

	sumOfAcceptedRatings := 0
	for _, line := range lines[ratingStartIndex:] {
		values := strings.Split(line[1:len(line)-1], ",")
		rating := Rating{}
		for _, value := range values {
			parts := strings.Split(value, "=")
			switch parts[0] {
			case "x":
				rating.x, _ = strconv.Atoi(parts[1])
			case "m":
				rating.m, _ = strconv.Atoi(parts[1])
			case "a":
				rating.a, _ = strconv.Atoi(parts[1])
			case "s":
				rating.s, _ = strconv.Atoi(parts[1])
			}
		}
		accepted := isAccepted(rating, workflows)
		ratingSum := rating.x + rating.m + rating.a + rating.s
		if accepted {
			sumOfAcceptedRatings += ratingSum
		}
		if debug {
			fmt.Println(rating, accepted, ratingSum)
		}
	}
	fmt.Printf("Sum of accepted ratings: %v\n", sumOfAcceptedRatings)
}

func getCountOfAcceptedRatings(ratingRange RatingRange, workflows map[string][]Rule, workflowName string) int {
	if workflowName == "R" {
		return 0
	}
	if workflowName == "A" {
		count := (ratingRange.maxX - ratingRange.minX + 1) * (ratingRange.maxM - ratingRange.minM + 1) * (ratingRange.maxA - ratingRange.minA + 1) * (ratingRange.maxS - ratingRange.minS + 1)
		return count
	}
	if debug {
		fmt.Printf("Current workflow: %v range:%v\n", workflowName, ratingRange)
	}
	workflow := workflows[workflowName]
	for _, rule := range workflow {
		if rule.alwaysGoToDestination {
			return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
		}

		switch rule.variable {
		case "x":
			if rule.operator == "<" {
				if ratingRange.maxX < rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.minX < rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.maxX = rule.value - 1
					ratingRange.minX = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else if rule.operator == ">" {
				if ratingRange.minX > rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.maxX > rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.minX = rule.value + 1
					ratingRange.maxX = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else {
				panic(fmt.Sprintf("Unknown operator: %v", rule.operator))
			}
		case "m":
			if rule.operator == "<" {
				if ratingRange.maxM < rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.minM < rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.maxM = rule.value - 1
					ratingRange.minM = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else if rule.operator == ">" {
				if ratingRange.minM > rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.maxM > rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.minM = rule.value + 1
					ratingRange.maxM = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else {
				panic(fmt.Sprintf("Unknown operator: %v", rule.operator))
			}
		case "a":
			if rule.operator == "<" {
				if ratingRange.maxA < rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.minA < rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.maxA = rule.value - 1
					ratingRange.minA = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else if rule.operator == ">" {
				if ratingRange.minA > rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.maxA > rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.minA = rule.value + 1
					ratingRange.maxA = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else {
				panic(fmt.Sprintf("Unknown operator: %v", rule.operator))
			}
		case "s":
			if rule.operator == "<" {
				if ratingRange.maxS < rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.minS < rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.maxS = rule.value - 1
					ratingRange.minS = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else if rule.operator == ">" {
				if ratingRange.minS > rule.value {
					// whole range goes to destination
					return getCountOfAcceptedRatings(ratingRange, workflows, rule.destination)
				} else if ratingRange.maxS > rule.value {
					// only a part of the range goes to destination
					newRange := ratingRange
					newRange.minS = rule.value + 1
					ratingRange.maxS = rule.value
					return getCountOfAcceptedRatings(newRange, workflows, rule.destination) + getCountOfAcceptedRatings(ratingRange, workflows, workflowName)
				} else {
					// try next rule
				}
			} else {
				panic(fmt.Sprintf("Unknown operator: %v", rule.operator))
			}
		}
	}
	panic(fmt.Sprintf("No rule found for workflow: %v", workflowName))
}

func partTwo(lines []string) {
	workflows, _ := parseWorkflows(lines)
	initialRange := RatingRange{
		minX: 1,
		maxX: 4000,
		minM: 1,
		maxM: 4000,
		minA: 1,
		maxA: 4000,
		minS: 1,
		maxS: 4000,
	}

	countOfAcceptedRatings := getCountOfAcceptedRatings(initialRange, workflows, "in")
	fmt.Printf("Count of accepted ratings: %v\n", countOfAcceptedRatings)
}

func main() {
	day := 19
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
