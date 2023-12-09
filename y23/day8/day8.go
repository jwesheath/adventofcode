package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jwesheath/adventofcode/timer"
)

type Node struct {
	id    string
	left  string
	right string
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

	nodes, partTwoStartingNodes, instructions := parseInput(file)

	partOneAnswer := countStepsToEnd(nodes, nodes["AAA"], instructions)

	var partTwoCounts []int
	for _, node := range partTwoStartingNodes {
		steps := countStepsToEnd(nodes, node, instructions)
		partTwoCounts = append(partTwoCounts, steps)
	}
	partTwoAnswer := lcmMultiple(partTwoCounts)

	fmt.Println("Part One:")
	fmt.Println(partOneAnswer)
	fmt.Println()
	fmt.Println("Part Two:")
	fmt.Println(partTwoAnswer)

}

func parseInput(file *os.File) (map[string]Node, []Node, string) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	nodes := make(map[string]Node)
	var partTwoStartingNodes []Node

	for scanner.Scan() {
		line := scanner.Text()

		splitLine := strings.Split(line, " = ")
		node := splitLine[0]

		nextNodes := splitLine[1]
		nextNodes = strings.ReplaceAll(strings.ReplaceAll(nextNodes, "(", ""), ")", "")

		splitNodes := strings.Split(nextNodes, ", ")

		left := splitNodes[0]
		right := splitNodes[1]

		newNode := Node{node, left, right}
		if []rune(node)[len(node)-1] == 'A' {
			partTwoStartingNodes = append(partTwoStartingNodes, newNode)
		}
		nodes[node] = newNode
	}
	return nodes, partTwoStartingNodes, instructions
}

func countStepsToEnd(nodes map[string]Node, startingNode Node, instructions string) int {
	steps := 0
	node := startingNode

	for {
		if node.id == "ZZZ" || isEndNode(node) {
			break
		}

		steps++

		nextMove := instructions[(steps-1)%len(instructions)]
		if nextMove == 'L' {
			node = nodes[node.left]
		} else {
			node = nodes[node.right]
		}
	}
	return steps
}

func isStartNode(node Node) bool {
	return []rune(node.id)[len(node.id)-1] == 'A'
}

func isEndNode(node Node) bool {
	return []rune(node.id)[len(node.id)-1] == 'Z'
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmMultiple(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return lcm(nums[0], lcmMultiple(nums[1:]))
}
