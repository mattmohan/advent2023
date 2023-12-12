package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9Main() {
	file, err := os.ReadFile("day9.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	total := int64(0)
	total2 := int64(0)

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		parsed := strings.Split(line, " ")
		curr := make([]int64, 0, len(parsed))

		for _, str := range parsed {
			num, _ := strconv.ParseInt(str, 10, 64)
			curr = append(curr, num)
		}
		total += day9Part1(curr)
		total2 += day9Part2(curr)
	}
	fmt.Printf("Total %v\n", total)
	fmt.Printf("Total2 %v\n", total2)

}

func day9Part1(curr []int64) int64 {
	lastIdx := len(curr) - 1
	next := make([]int64, lastIdx)
	allZero := true
	for i := range curr[:lastIdx] {
		next[i] = curr[i+1] - curr[i]
		if next[i] != 0 {
			allZero = false
		}
	}

	if allZero {
		return curr[lastIdx]
	}

	return curr[lastIdx] + day9Part1(next)
}

func day9Part2(curr []int64) int64 {
	lastIdx := len(curr) - 1
	next := make([]int64, lastIdx)
	allZero := true
	for i := range curr[:lastIdx] {
		next[i] = curr[i+1] - curr[i]
		if next[i] != 0 {
			allZero = false
		}
	}

	if allZero {
		return curr[0]
	}

	return curr[0] - day9Part2(next)
}
