package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jwesheath/adventofcode/safesum"
	"github.com/jwesheath/adventofcode/timer"
)

type Coord struct {
	i, j int
}

func (coord *Coord) getAdjacents() []Coord {
	return []Coord{
		{coord.i - 1, coord.j - 1},
		{coord.i - 1, coord.j},
		{coord.i - 1, coord.j + 1},
		{coord.i, coord.j + 1},
		{coord.i + 1, coord.j + 1},
		{coord.i + 1, coord.j},
		{coord.i + 1, coord.j - 1},
		{coord.i, coord.j - 1},
	}
}

type PartNumber struct {
	number int
	coords []Coord
}

func (partNumber *PartNumber) isAdjacentToAny(coordSet map[Coord]bool) bool {
	for _, coord := range partNumber.coords {
		for _, adjacent := range coord.getAdjacents() {
			if coordSet[adjacent] {
				return true
			}
		}
	}
	return false
}

func main() {
	// time it
	defer timer.Timer()()

	// input
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("couldn't read input")
	}
	defer file.Close()

	// parse input
	partNumbers, symbolCoordSet, starCoords, coordToPartNumber := parseInput(file)

	// sum
	partOneAnswer := safesum.NewSafeSum()
	partTwoAnswer := safesum.NewSafeSum()

	var wg sync.WaitGroup
	for _, partNumber := range partNumbers {
		wg.Add(1)
		go func(partNumber PartNumber) {
			if partNumber.isAdjacentToAny(symbolCoordSet) {
				partOneAnswer.Add(partNumber.number)
			}
			wg.Done()
		}(partNumber)
	}

	for _, starCoord := range starCoords {
		wg.Add(1)

		go func(starCoord Coord) {
			adjacentParts := make(map[int]bool)
			for _, coord := range starCoord.getAdjacents() {
				number, exists := coordToPartNumber[coord]
				if exists {
					adjacentParts[number] = true
				}
			}

			if len(adjacentParts) == 2 {
				gearRatio := 1
				for k := range adjacentParts {
					gearRatio *= k
				}
				partTwoAnswer.Add(gearRatio)
			}
			wg.Done()
		}(starCoord)
	}
	wg.Wait()

	// answer
	fmt.Println("Part 1:")
	partOneAnswer.Print()
	fmt.Println("Part 2:")
	partTwoAnswer.Print()
	fmt.Println("")

}

func parseInput(file *os.File) ([]PartNumber, map[Coord]bool, []Coord, map[Coord]int) {
	scanner := bufio.NewScanner(file)

	var partNumbers []PartNumber
	symbolCoordSet := make(map[Coord]bool)
	var starCoords []Coord
	coordToPartNumber := make(map[Coord]int)

	i := -1
	for scanner.Scan() {
		i++

		var num []byte
		buff := bytes.NewBuffer(num)
		var coords []Coord

		for j, char := range scanner.Text() {

			if isDigit(char) {
				buff.WriteRune(char)
				coords = append(coords, Coord{i, j})
				continue
			}

			if buff.Len() > 0 {
				number := flushBuffToNum(buff)
				partNumbers = append(partNumbers, PartNumber{number, coords})
				for _, coord := range coords {
					coordToPartNumber[coord] = number
				}
				coords = nil
			}

			if isSymbol(char) {
				symbolCoordSet[Coord{i, j}] = true
			}

			if isStar(char) {
				starCoords = append(starCoords, Coord{i, j})
			}
		}

		if buff.Len() > 0 {
			number := flushBuffToNum(buff)
			partNumbers = append(partNumbers, PartNumber{number, coords})
			for _, coord := range coords {
				coordToPartNumber[coord] = number
			}
		}
	}
	return partNumbers, symbolCoordSet, starCoords, coordToPartNumber
}

func flushBuffToNum(buff *bytes.Buffer) int {
	number, err := strconv.Atoi(buff.String())
	if err != nil {
		panic("stuff in buffer doesn't parse to int")
	}
	buff.Reset()
	return number
}

func isEmpty(char rune) bool {
	return char == '.'
}

func isDigit(char rune) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func isSymbol(char rune) bool {
	return !isEmpty(char) && !isDigit(char)
}

func isStar(char rune) bool {
	return char == '*'
}
