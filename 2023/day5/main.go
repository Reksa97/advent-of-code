package main

import (
	"aoc2023/common"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MapName string

const (
	SeedToSoil            MapName = "seed-to-soil"
	SoilToFertilizer      MapName = "soil-to-fertilizer"
	FertilizerToWater     MapName = "fertilizer-to-water"
	WaterToLight          MapName = "water-to-light"
	LightToTemperature    MapName = "light-to-temperature"
	TemperatureToHumidity MapName = "temperature-to-humidity"
	HumidityToLocation    MapName = "humidity-to-location"
	None                  MapName = ""
)

var allMapNames = []MapName{
	SeedToSoil,
	SoilToFertilizer,
	FertilizerToWater,
	WaterToLight,
	LightToTemperature,
	TemperatureToHumidity,
	HumidityToLocation,
}

func toMapName(s string) MapName {
	switch s {
	case "seed-to-soil":
		return SeedToSoil
	case "soil-to-fertilizer":
		return SoilToFertilizer
	case "fertilizer-to-water":
		return FertilizerToWater
	case "water-to-light":
		return WaterToLight
	case "light-to-temperature":
		return LightToTemperature
	case "temperature-to-humidity":
		return TemperatureToHumidity
	case "humidity-to-location":
		return HumidityToLocation
	}
	return None
}

/* type MapInt int

const (
	SeedToSoil MapInt = iota
	SoilToFertilizer
	FertilizerToWater
	WaterToLight
	LightToTemperature
	TemperatureToHumidity
	HumidityToLocation
) */

type Map struct {
	destination int
	source      int
	size        int
}

func findDestination(maps map[MapName][]Map, mapName MapName, source int) int {
	for _, m := range maps[mapName] {
		if source >= m.source && source < m.source+m.size {
			return m.destination + (source - m.source)
		}
	}
	return source
}

func partOne(lines []string, debug bool) {
	seeds := strings.Split(strings.Split(lines[0], "seeds: ")[1], " ")

	mapName := None

	maps := make(map[MapName][]Map)

	for _, line := range lines[2:] {
		if mapName == "" && strings.HasSuffix(line, " map:") {
			mapName = toMapName(strings.Split(line, " map:")[0])
			if debug {
				fmt.Println(mapName)
			}
			continue
		}
		if line == "" {
			mapName = ""
			continue
		}

		sliceSplit := strings.Split(line, " ")
		destination, _ := strconv.Atoi(sliceSplit[0])
		source, _ := strconv.Atoi(sliceSplit[1])
		size, _ := strconv.Atoi(sliceSplit[2])

		maps[mapName] = append(maps[mapName], Map{
			destination: destination,
			source:      source,
			size:        size,
		})

	}

	if debug {
		fmt.Println(maps)
	}

	minLocation := math.MaxInt
	for _, seedString := range seeds {
		source, _ := strconv.Atoi(seedString)
		if debug {
			fmt.Println("seed", source)
		}
		for _, mapName := range allMapNames {
			source = findDestination(maps, mapName, source)
			if debug {
				fmt.Println(mapName, source)
			}
		}
		if source < minLocation {
			minLocation = source
		}
		if debug {
			fmt.Println("location", source)
			fmt.Println("")
		}
	}
	fmt.Println("Lowest location number", minLocation)
}

func partTwo(lines []string, debug bool) {

}

func main() {
	day := 5
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	part := common.GetPart(os.Args)
	debug := common.IsTestMode(os.Args)
	fmt.Printf("Part: %v\n", part)

	if part == 1 {
		partOne(lines, debug)
	} else {
		partTwo(lines, debug)
	}
}
