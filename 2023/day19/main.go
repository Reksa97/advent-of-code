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
	variable   string
	operator   string
	value      int
	goTo       string
	alwaysGoTo bool
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
			if rule.alwaysGoTo {
				currentWorkflowName = rule.goTo
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
					currentWorkflowName = rule.goTo
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
						fmt.Printf("New workflow: %v\n", rule.goTo)
					}
					foundTrue = true
					currentWorkflowName = rule.goTo
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
	rule.goTo = parts[1]

	return rule
}

func partOne(lines []string) {
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
				workflow = append(workflow, Rule{goTo: currentRule, alwaysGoTo: true})
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

func partTwo(lines []string) {

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
