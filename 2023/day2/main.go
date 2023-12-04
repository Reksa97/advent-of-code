package main

import (
	"aoc2023/common"
	"fmt"
	"os"
)

func main() {
	day := 2
	games, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	part := common.GetPart(os.Args)
	fmt.Printf("Part: %v\n", part)
	for _, game := range games {
		fmt.Println(game)
	}
}
