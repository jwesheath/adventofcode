package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/jwesheath/adventofcode/safesum"
	"github.com/jwesheath/adventofcode/timer"
)

const (
	digits = "0123456789"
)

var numbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	fmt.Println("Part One:")
	getAnswer('a')
	fmt.Println("")
	fmt.Println("Part Two:")
	getAnswer('b')
}

func getAnswer(part rune) {
	// time it
	defer timer.Timer()()

	// read input
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("couldn't read input")
	}
	defer file.Close()

	// init
	scanner := bufio.NewScanner(file)
	sum := safesum.NewSafeSum()

	includeWords := false
	if part == 'b' {
		includeWords = true
	}

	// scan and sum
	var wg sync.WaitGroup
	for scanner.Scan() {
		text := scanner.Text()
		wg.Add(1)
		go func(input string) {
			result := extractCalibrationValue(input, includeWords)
			sum.Add(result)
			wg.Done()
		}(text)
	}
	wg.Wait()

	// answer
	sum.Print()
}

func extractCalibrationValue(input string, includeWords bool) int {

	firstIdx, lastIdx := len(input), -1
	firstInt, lastInt := -1, -1

	// part A - ("0", "1", "2", ...) indices
	firstIdx = strings.IndexAny(input, digits)
	lastIdx = strings.LastIndexAny(input, digits)
	firstInt = int(input[firstIdx] - '0')
	lastInt = int(input[lastIdx] - '0')

	// partB - ("zero", "one", "two", ...)
	if includeWords {
		for numberString, number := range numbers {

			var idx int
			idx = strings.Index(input, numberString)
			if idx != -1 && idx < firstIdx {
				firstIdx = idx
				firstInt = number
			}

			idx = strings.LastIndex(input, numberString)
			if idx != -1 && idx > lastIdx {
				lastIdx = idx
				lastInt = number
			}
		}
	}

	if firstInt == -1 || lastInt == -1 {
		panic("your assumptions are wrong")
	}

	return firstInt*10 + lastInt
}
