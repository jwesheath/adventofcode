package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/jwesheath/adventofcode/timer"
)

type Range struct {
	destination int
	source      int
	length      int
}

type Seeds []int
type Almanac map[string][]Range

func main() {
	// time it
	defer timer.Timer()()

	// read input
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}

	// parse input
	seeds, almanac := parseInput(file)
	locations := make([]int, len(seeds))

	// part one
	for i, seed := range seeds {
		locations[i] = processSeed(seed, almanac)
	}

	partOneAnswer := slices.Min(locations)

	// part two
	// TODO: find optimal fast solution
	var partTwoLocations []int

	var seedStarts []int
	var seedRanges []int
	for i, seed := range seeds {
		if i%2 == 0 {
			seedStarts = append(seedStarts, seed)
		} else {
			seedRanges = append(seedRanges, seed)
		}
	}

	for i, seedStart := range seedStarts {
		for j := 0; j < seedRanges[i]; j++ {
			seed := seedStart + j
			partTwoLocations = append(partTwoLocations, processSeed(seed, almanac))
		}
	}

	partTwoAnswer := slices.Min(partTwoLocations)

	// print
	fmt.Println("Part One:")
	fmt.Println(partOneAnswer)
	fmt.Println()
	fmt.Println("Part Two:")
	fmt.Println(partTwoAnswer)

}

func mapSeed(seed int, almanacMap []Range) int {
	for _, r := range almanacMap {
		if seed >= r.source && seed <= r.source+r.length-1 {
			return seed + r.destination - r.source
		}
	}
	return seed
}

func processSeed(seed int, almanac Almanac) int {
	seed = mapSeed(seed, almanac["seedToSoil"])
	seed = mapSeed(seed, almanac["soilToFertilizer"])
	seed = mapSeed(seed, almanac["fertilizerToWater"])
	seed = mapSeed(seed, almanac["waterToLight"])
	seed = mapSeed(seed, almanac["lightToTemperature"])
	seed = mapSeed(seed, almanac["temperatureToHumidity"])
	seed = mapSeed(seed, almanac["humidityToLocation"])
	return seed
}

func parseInput(file *os.File) (Seeds, Almanac) {
	isAlpha := regexp.MustCompile(`^[a-z]`)
	isNumeric := regexp.MustCompile(`^[0-9]`)

	almanac := make(Almanac)

	scanner := bufio.NewScanner(file)

	// seeds
	scanner.Scan()
	firstLine := scanner.Text()
	seedStrings := strings.Fields(strings.Split(firstLine, ": ")[1])
	seeds := make([]int, len(seedStrings))
	for i := 0; i < len(seedStrings); i++ {
		intValue, _ := strconv.Atoi(seedStrings[i])
		seeds[i] = intValue
	}

	// almanac
	var mapName string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if isAlpha.MatchString(line) {
			line = strings.Split(line, " ")[0]
			lineRunes := []rune(line)
			for i := 0; i < 2; i++ {
				dashIdx := strings.Index(string(lineRunes), "-")
				lineRunes[dashIdx+1] = unicode.ToUpper(rune(lineRunes[dashIdx+1]))
				lineRunes = append(lineRunes[:dashIdx], lineRunes[dashIdx+1:]...)
			}
			mapName = string(lineRunes)
		}

		if isNumeric.MatchString(line) {
			rangeStrings := strings.Fields(line)
			rangeInts := make([]int, len(rangeStrings))
			for i := 0; i < len(rangeStrings); i++ {
				rangeInt, _ := strconv.Atoi(rangeStrings[i])
				rangeInts[i] = rangeInt
			}
			newRange := Range{rangeInts[0], rangeInts[1], rangeInts[2]}
			almanac[mapName] = append(almanac[mapName], newRange)
		}

	}

	return seeds, almanac
}
