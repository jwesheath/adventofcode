package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jwesheath/adventofcode/timer"
)

func main() {

	// time it
	defer timer.Timer()()

	// parse input
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringNums := strings.Fields(scanner.Text())
		var line []int
		for _, stringNum := range stringNums {
			num, err := strconv.Atoi(stringNum)
			if err != nil {
				panic("bad input")
			}
			line = append(line, num)
		}
		lines = append(lines, line)
	}

	// answer
	partOneAnswer := 0
	partTwoAnswer := 0
	for _, line := range lines {
		prevValue, nextValue := solveLine(line)
		partOneAnswer += nextValue
		partTwoAnswer += prevValue
	}

	// print
	fmt.Println("Part One:")
	fmt.Println(partOneAnswer)
	fmt.Println()
	fmt.Println("Part Two:")
	fmt.Println(partTwoAnswer)
}

func solveLine(line []int) (int, int) {
	pyramid := createPyramid(line)
	prevValue := getPrevLineValues(pyramid)
	nextValue := getNextLineValues(pyramid)
	return prevValue, nextValue
}

func createPyramid(line []int) [][]int {
	var lines [][]int
	lines = append(lines, line)

	for {
		lastLine := lines[len(lines)-1]
		if allZeros(lastLine) {
			break
		}
		var newLine []int
		for i := 1; i < len(lastLine); i++ {
			newLine = append(newLine, lastLine[i]-lastLine[i-1])
		}
		lines = append(lines, newLine)
	}

	return lines
}

func getNextLineValues(pyramid [][]int) int {
	sum := 0
	for _, line := range pyramid {
		sum += line[len(line)-1]
	}
	return sum
}

func getPrevLineValues(pyramid [][]int) int {
	val := 0
	for i := len(pyramid) - 2; i >= 0; i-- {
		val = pyramid[i][0] - val
	}
	return val
}

func allZeros(line []int) bool {
	ret := true
	for _, num := range line {
		if num != 0 {
			ret = false
		}
	}
	return ret
}
