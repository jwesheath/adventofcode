package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jwesheath/adventofcode/timer"
)

func main() {
	// time it
	defer timer.Timer()()

	// read input
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}
	defer file.Close()

	// get answers
	scanner := bufio.NewScanner(file)
	var numWinnings []int
	partOneAnswer := 0

	for scanner.Scan() {
		line := scanner.Text()
		card, winningNumbers := parseInput(line)
		numWinning, score := scoreCard(card, winningNumbers)
		partOneAnswer += score
		numWinnings = append(numWinnings, numWinning)
	}

	partTwoAnswer := countCardCopies(numWinnings)

	// print answers
	fmt.Println("Part One:")
	fmt.Println(partOneAnswer)
	fmt.Println()

	fmt.Println("Part Two:")
	fmt.Println(partTwoAnswer)
}

func scoreCard(card []int, winningNumbers map[int]bool) (int, int) {
	numWinning := 0
	for _, number := range card {
		if winningNumbers[number] {
			numWinning++
		}
	}
	if numWinning > 0 {
		score := int(math.Pow(2, float64(numWinning-1)))
		return numWinning, score
	}
	return 0, 0
}

func countCardCopies(numWinnings []int) int {

	cardCounts := make(map[int]int, len(numWinnings))
	for i := range numWinnings {
		cardCounts[i] = 1
	}
	for i, numWinning := range numWinnings {
		copies := cardCounts[i]
		for j := 1; j <= numWinning; j++ {
			cardCounts[i+j] += copies
		}
	}
	numCopies := 0
	for _, value := range cardCounts {
		numCopies += value
	}
	return numCopies
}

func parseInput(line string) ([]int, map[int]bool) {

	split := strings.Split(strings.Split(line, ": ")[1], " | ")

	var card []int
	for _, cardString := range strings.Fields(split[0]) {
		cardNumber, err := strconv.Atoi(cardString)
		if err != nil {
			panic("couldn't parse number from string")
		}
		card = append(card, cardNumber)
	}

	winningNumbers := make(map[int]bool)
	for _, winningString := range strings.Fields(split[1]) {
		winningNumber, err := strconv.Atoi(winningString)
		if err != nil {
			panic("couldn't parse number from string")
		}
		winningNumbers[winningNumber] = true
	}

	return card, winningNumbers
}
