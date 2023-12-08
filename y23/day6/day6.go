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

	// read and parse input
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	timeStrings := strings.Fields(strings.Split(lines[0], ": ")[1])
	distanceStrings := strings.Fields(strings.Split(lines[1], ": ")[1])

	times, distances := make([]int, len(timeStrings)), make([]int, len(timeStrings))
	for i := range timeStrings {
		time, _ := strconv.Atoi(timeStrings[i])
		distance, _ := strconv.Atoi(distanceStrings[i])
		times[i] = time
		distances[i] = distance
	}

	bigTime, _ := strconv.Atoi(strings.Join(timeStrings, ""))
	bigDistance, _ := strconv.Atoi(strings.Join(distanceStrings, ""))

	// answer
	partOneAnswer := 1
	for i, time := range times {
		distance := distances[i]
		numBetter := solveQuadratic(float64(time), float64(distance))
		partOneAnswer *= numBetter
	}

	partTwoAnswer := solveQuadratic(float64(bigTime), float64(bigDistance))

	fmt.Println("Part One:")
	fmt.Println(partOneAnswer)
	fmt.Println()
	fmt.Println("Part Two:")
	fmt.Println(partTwoAnswer)
}

func solveQuadratic(time float64, goal float64) int {
	// x * (t - x) = g
	// -x^2 + tx + -g = 0
	low := ((-1 * time) + math.Sqrt(math.Pow(time, 2)-4*goal)) / -2
	high := ((-1 * time) - math.Sqrt(math.Pow(time, 2)-4*goal)) / -2
	return int(high) - int(low)

}
