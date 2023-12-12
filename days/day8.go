package days

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var day8Regex = regexp.MustCompile(`^([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)$`)

func Day8Main() {
	file, err := os.ReadFile("day8.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	lines := strings.Split(string(file), "\n")
	path := lines[0]
	nodeStrs := make(map[string][]string, len(lines))
	for _, nodeStr := range lines[2:] {
		matches := day8Regex.FindStringSubmatch(nodeStr)
		nodeStrs[matches[1]] = matches[2:]
	}
	/*
		i := 0
		for currNode := "AAA"; currNode != "ZZZ"; i++ {
			currNode = nodeStrs[currNode][nextInstruction(path, i)]
		}
	*/
	currNodes2 := make([]string, 0, len(nodeStrs))
	for k := range nodeStrs {
		if k[2] == 'A' {
			currNodes2 = append(currNodes2, k)
		}
	}

	lengths := make([]int, 0, len(currNodes2))
	for _, startNode := range currNodes2 {
		i := 0
		for currNode := startNode; currNode[2] != 'Z'; i++ {
			currNode = nodeStrs[currNode][nextInstruction(path, i)]
		}
		fmt.Printf("%v\n", i)
		lengths = append(lengths, i)
	}
	fmt.Printf("Total %v\n", LCM(lengths[0], lengths[1], lengths[2:]...))
}

func nextInstruction(path string, step int) int {
	if path[step%len(path)] == 'L' {
		return 0
	}
	return 1
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
