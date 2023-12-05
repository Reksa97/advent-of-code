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

type Map struct {
	destination int
	source      int
	size        int
}

type SourceRange struct {
	min int
	max int
}

func findDestination(maps map[MapName][]Map, mapName MapName, source int) int {
	for _, m := range maps[mapName] {
		if source >= m.source && source < m.source+m.size {
			return m.destination + (source - m.source)
		}
	}
	return source
}

func findDestinations(maps map[MapName][]Map, mapName MapName, sources SourceRange) []SourceRange {
	destinationRanges := []SourceRange{}
	for _, m := range maps[mapName] {
		if sources.min >= m.source && sources.max < m.source+m.size {
			// whole source range goes in single destination range, no need to split source range
			destinationRanges = append(destinationRanges, SourceRange{
				min: m.destination + (sources.min - m.source),
				max: m.destination + (sources.max - m.source),
			})
			return destinationRanges
		} else if sources.min >= m.source && sources.min < m.source+m.size {
			// source range starts in destination range but goes over it
			destinationRanges = append(destinationRanges, SourceRange{
				min: m.destination + (sources.min - m.source),
				max: m.destination + (m.size - 1),
			})
			destinationRanges = append(destinationRanges, findDestinations(maps, mapName, SourceRange{
				min: m.source + m.size,
				max: sources.max,
			})...)
			return destinationRanges
		} else if sources.max >= m.source && sources.max < m.source+m.size {
			// source range ends in destination range but starts before it
			destinationRanges = append(destinationRanges, SourceRange{
				min: m.destination,
				max: m.destination + (sources.max - m.source),
			})
			destinationRanges = append(destinationRanges, findDestinations(maps, mapName, SourceRange{
				min: sources.min,
				max: m.source - 1,
			})...)
			return destinationRanges
		} else if sources.min < m.source && sources.max >= m.source+m.size {
			// source range starts before destination range and ends after it
			destinationRanges = append(destinationRanges, SourceRange{
				min: m.destination,
				max: m.destination + (m.size - 1),
			})
			destinationRanges = append(destinationRanges, findDestinations(maps, mapName, SourceRange{
				min: sources.min,
				max: m.source - 1,
			})...)
			destinationRanges = append(destinationRanges, findDestinations(maps, mapName, SourceRange{
				min: m.source + m.size,
				max: sources.max,
			})...)
			return destinationRanges
		}
	}
	// no destination range found, return original source range
	destinationRanges = append(destinationRanges, sources)
	return destinationRanges
}

func getMaps(lines []string, debug bool) map[MapName][]Map {
	maps := make(map[MapName][]Map)
	mapName := None
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
	return maps
}

func partOne(lines []string, debug bool) {
	seeds := strings.Split(strings.Split(lines[0], "seeds: ")[1], " ")

	maps := getMaps(lines, debug)
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
	seeds := strings.Split(strings.Split(lines[0], "seeds: ")[1], " ")
	minLocation := math.MaxInt
	maps := getMaps(lines, debug)
	for i := 0; i < len(seeds); i += 2 {
		min, _ := strconv.Atoi(seeds[i])
		size, _ := strconv.Atoi(seeds[i+1])
		sourceRange := SourceRange{min: min, max: min + size - 1}
		sourceRanges := []SourceRange{sourceRange}
		for _, mapName := range allMapNames {
			newSourceRanges := []SourceRange{}
			for _, sourceRange := range sourceRanges {
				newSourceRanges = append(newSourceRanges, findDestinations(maps, mapName, sourceRange)...)
			}
			if debug {
				fmt.Println(mapName, newSourceRanges)
			}
			sourceRanges = newSourceRanges
		}
		for _, sourceRange := range sourceRanges {
			if sourceRange.min < minLocation {
				minLocation = sourceRange.min
			}
		}
		if debug {
			fmt.Println(sourceRange)
		}
	}
	fmt.Println("Lowest location number", minLocation)
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
