package days

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type GridCoord struct {
	row int
	col int
}
type Node struct {
	distance int
}

func Day10Main() {
	file, err := os.ReadFile("day10.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	lines := strings.Split(string(file), "\n")
	startRow := 0
	startCol := 0
	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				startRow = i
				startCol = j
			}
		}
	}
	start := GridCoord{startRow, startCol}

	nodes := make(map[GridCoord]*Node, len(file))
	nodes[start] = &Node{distance: 0}

	// From the starting point find neighbors that link back to us
	candidates := findConnections(lines, start)
	queue := make([]GridCoord, 0, 4)
	for _, candidate := range candidates {
		if slices.Contains(findConnections(lines, candidate), start) {
			queue = append(queue, candidate)
		}
	}

	// Iterate over neighbors adding the next gen of neighbors to the queue
	for len(queue) > 0 {
		nextNode := queue[0]
		queue = queue[1:]
		nodes[nextNode] = &Node{distance: math.MaxInt}

		next := findConnections(lines, nextNode)
		fmt.Printf("queue: %v => %#v\n%v\n", nextNode, queue, next)
		for _, n := range next {
			node, ok := nodes[n]
			if !ok {
				queue = append(queue, n)
			} else {
				if node.distance+1 < nodes[nextNode].distance {
					nodes[nextNode].distance = nodes[n].distance + 1
				}
			}
		}
	}

	max := 0
	for _, v := range nodes {
		if v.distance > max {
			max = v.distance
		}
	}
	fmt.Printf("Total: %v\n", max)
}

func findConnections(grid []string, pos GridCoord) []GridCoord {
	switch grid[pos.row][pos.col] {
	case 'S':
		candidates := []GridCoord{
			{pos.row - 1, pos.col},
			{pos.row + 1, pos.col},
			{pos.row, pos.col - 1},
			{pos.row, pos.col + 1},
		}
		return candidates
	case '|':
		return []GridCoord{
			{pos.row - 1, pos.col},
			{pos.row + 1, pos.col},
		}
	case '-':
		return []GridCoord{
			{pos.row, pos.col - 1},
			{pos.row, pos.col + 1},
		}
	case 'F':
		return []GridCoord{
			{pos.row + 1, pos.col},
			{pos.row, pos.col + 1},
		}

	case 'J':
		return []GridCoord{
			{pos.row - 1, pos.col},
			{pos.row, pos.col - 1},
		}
	case '7':
		return []GridCoord{
			{pos.row + 1, pos.col},
			{pos.row, pos.col - 1},
		}
	case 'L':
		return []GridCoord{
			{pos.row - 1, pos.col},
			{pos.row, pos.col + 1},
		}
	}
	return []GridCoord{}
}
