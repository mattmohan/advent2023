package main

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
	C byte
}
type Number struct {
	Number int
	Start  int
	End    int
}

func Day3Main() {
	file, err := os.ReadFile("day3.txt")
	if err != nil {
		panic(fmt.Errorf("Got error reading input: %w", err))
	}

	row := 0
	col := 0

	symbols := make([]Point, 0, len(file))
	// Store an array of "Numbers" for each row
	numbers := make([][]Number, 0, 140)
	numbers = append(numbers, make([]Number, 70))
	current := 0
	currentStart := -1

	// A legit bytewise parser!
	for _, char := range file {
		oldCol := col
		col++
		if char >= '0' && char <= '9' {
			current = current*10 + int(char-'0')
			if currentStart < 0 {
				currentStart = oldCol
			}
			continue
		}
		if current > 0 {
			numbers[row][currentStart] = Number{Number: current, Start: currentStart, End: oldCol - 1}
			currentStart = -1
			current = 0
		}
		if char == '\n' {
			col = 0
			row++
			numbers = append(numbers, make([]Number, 70))
			continue
		} else if char != '.' {
			symbols = append(symbols, Point{oldCol, row, char})
		}
	}

	total := 0
	gearTotal := 0
	// Map used to track which "Number"s we've already considered for part1
	considered := make(map[Number]bool)
	for _, symbol := range symbols {
		nums := make([]int, 0, 8)
		for i := symbol.Y - 1; i <= symbol.Y+1; i++ {
			for _, num := range numbers[i] {
				if (AbsInt(symbol.X-num.Start) <= 1) || (AbsInt(symbol.X-num.End) <= 1) {
					if !considered[num] {
						total += num.Number
						considered[num] = true
					}
					nums = append(nums, num.Number)
				}
			}
		}

		// Handle part2
		if symbol.C == '*' && len(nums) == 2 {
			gearTotal += nums[0] * nums[1]
		}
	}

	fmt.Printf("Total: %v, GearTotal: %v", total, gearTotal)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
