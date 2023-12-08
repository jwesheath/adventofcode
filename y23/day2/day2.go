package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/jwesheath/adventofcode/safesum"
	"github.com/jwesheath/adventofcode/timer"
)

type CubeCount struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	rounds []CubeCount
}

var bag = CubeCount{
	red:   12,
	green: 13,
	blue:  14,
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

	// init
	partOneAnswer := safesum.NewSafeSum()
	partTwoAnswer := safesum.NewSafeSum()
	scanner := bufio.NewScanner(file)

	// scan and sum
	var wg sync.WaitGroup
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			game := processLine(line)
			if isPossible(game) {
				partOneAnswer.Add(game.id)
			}
			partTwoAnswer.Add(minimumPossible(game))
			wg.Done()
		}(line)
	}
	wg.Wait()

	// answer
	fmt.Println(("Part One: "))
	partOneAnswer.Print()
	fmt.Println(("Part Two: "))
	partTwoAnswer.Print()
}

func isPossible(game Game) bool {
	for _, round := range game.rounds {
		if round.red > bag.red ||
			round.green > bag.green ||
			round.blue > bag.blue {
			return false
		}
	}
	return true
}

func minimumPossible(game Game) int {
	count := CubeCount{}
	for _, round := range game.rounds {
		count.red = max(count.red, round.red)
		count.green = max(count.green, round.green)
		count.blue = max(count.blue, round.blue)
	}
	return count.red * count.green * count.blue
}

// disgusting, totally unprincipled input parsing
func processLine(line string) Game {
	pieces := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Fields(pieces[0])[1])
	roundStrings := strings.Split(pieces[1], "; ")

	var rounds []CubeCount
	for _, roundString := range roundStrings {
		roundMap := make(map[string]int)
		countAndColors := strings.Split(roundString, ", ")
		for _, countAndColor := range countAndColors {
			countAndColorSplit := strings.Fields(countAndColor)
			count, _ := strconv.Atoi(countAndColorSplit[0])
			color := countAndColorSplit[1]
			roundMap[color] = count
		}
		newRound := CubeCount{
			red:   roundMap["red"],
			green: roundMap["green"],
			blue:  roundMap["blue"],
		}
		rounds = append(rounds, newRound)
	}
	return Game{gameId, rounds}
}
