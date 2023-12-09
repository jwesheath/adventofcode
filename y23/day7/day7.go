package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jwesheath/adventofcode/timer"
)

// hand strenghts
// 7 - five of a kind
// 6 - four of a kind
// 5 - full house
// 4 - three of a kind
// 3 - two pair
// 2 - one pair
// 1 - high card

var wildCardUpgradeMap = map[int]int{
	// four of a kind -> five of a kind
	6: 7,
	// three of a kind -> four of a kind
	4: 6,
	// two pair -> full house
	3: 5,
	// one pair -> three of a kind
	2: 4,
	// high card -> one pair
	1: 2,
}

type Hand struct {
	cards    string
	handType int
	bid      int
}

func main() {
	// time it
	defer timer.Timer()()

	// parse input part one
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}
	defer file.Close()
	handsPartOne := parseInput(file)

	// parse input part two
	filePartTwo, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}
	defer file.Close()
	handsPartTwo := parseInputPartTwo(filePartTwo)

	// rank
	slices.SortFunc(handsPartOne, sortFunc)
	slices.SortFunc(handsPartTwo, sortFuncPartTwo)

	// answer
	partOneAnswer := 0
	for i, hand := range handsPartOne {
		partOneAnswer += (i + 1) * hand.bid
	}

	partTwoAnswer := 0
	for i, hand := range handsPartTwo {
		partTwoAnswer += (i + 1) * hand.bid
	}

	// print
	fmt.Println("Part One:")
	fmt.Println(partOneAnswer)
	fmt.Println()
	fmt.Println("Part Two:")
	fmt.Println(partTwoAnswer)

}

// part one
func parseInput(file *os.File) []Hand {
	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		cardsString, bidString := splitLine[0], splitLine[1]
		handType := determineHandType(cardsString)
		bid, err := strconv.Atoi(bidString)
		if err != nil {
			panic("bad input")
		}
		hands = append(hands, Hand{cardsString, handType, bid})
	}
	return hands
}

func determineHandType(cardsString string) int {

	cardCounts := make(map[rune]int)
	for _, card := range cardsString {
		cardCounts[card] += 1
	}

	countCounts := make(map[int]int)
	for _, count := range cardCounts {
		countCounts[count] += 1
	}

	switch {
	case countCounts[5] == 1:
		return 7
	case countCounts[4] == 1:
		return 6
	case countCounts[3] == 1 && countCounts[2] == 1:
		return 5
	case countCounts[3] == 1:
		return 4
	case countCounts[2] == 2:
		return 3
	case countCounts[2] == 1:
		return 2
	default:
		return 1
	}
}

func sortFunc(a Hand, b Hand) int {
	switch {
	case a.handType > b.handType:
		return 1
	case a.handType < b.handType:
		return -1
	default:
		return determineStrongerFirstCard(a, b)
	}
}

func determineStrongerFirstCard(a Hand, b Hand) int {
	cardValues := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	aCards, bCards := []rune(a.cards), []rune(b.cards)
	for i := range a.cards {
		aCard, bCard := aCards[i], bCards[i]
		switch {
		case cardValues[aCard] > cardValues[bCard]:
			return 1
		case cardValues[aCard] < cardValues[bCard]:
			return -1
		default:
			continue
		}
	}
	return 0
}

// part two
func parseInputPartTwo(file *os.File) []Hand {
	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		cardsString, bidString := splitLine[0], splitLine[1]
		handType := determineHandTypePartTwo(cardsString)
		bid, err := strconv.Atoi(bidString)
		if err != nil {
			panic("bad input")
		}
		hands = append(hands, Hand{cardsString, handType, bid})
	}
	return hands
}

func determineHandTypePartTwo(cardsString string) int {

	wildCount := strings.Count(cardsString, "J")
	cardsString = strings.ReplaceAll(cardsString, "J", "")

	cardCounts := make(map[rune]int)
	for _, card := range cardsString {
		cardCounts[card] += 1
	}

	countCounts := make(map[int]int)
	for _, count := range cardCounts {
		countCounts[count] += 1
	}

	handType := 1
	switch {
	case countCounts[5] == 1:
		handType = 7
	case countCounts[4] == 1:
		handType = 6
	case countCounts[3] == 1 && countCounts[2] == 1:
		handType = 5
	case countCounts[3] == 1:
		handType = 4
	case countCounts[2] == 2:
		handType = 3
	case countCounts[2] == 1:
		handType = 2
	default:
		handType = 1
	}

	// special case for all wilds
	if cardsString == "" && wildCount == 5 {
		handType = 7
		wildCount = 0
	}

	for i := 0; i < wildCount; i++ {
		handType = wildCardUpgradeMap[handType]
	}

	return handType
}

func sortFuncPartTwo(a Hand, b Hand) int {
	switch {
	case a.handType > b.handType:
		return 1
	case a.handType < b.handType:
		return -1
	default:
		return determineStrongerFirstCardPartTwo(a, b)
	}
}

func determineStrongerFirstCardPartTwo(a Hand, b Hand) int {
	cardValues := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	aCards, bCards := []rune(a.cards), []rune(b.cards)
	for i := range a.cards {
		aCard, bCard := aCards[i], bCards[i]
		switch {
		case cardValues[aCard] > cardValues[bCard]:
			return 1
		case cardValues[aCard] < cardValues[bCard]:
			return -1
		default:
			continue
		}
	}
	return 0
}
