package days

import (
	"fmt"
	"os"
	"strings"
)

func Day11Main() {
	file, err := os.ReadFile("day11.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	lines := strings.Split(string(file), "\n")

	total := uint64(0)
	total2 := uint64(0)

	// Track all galaxies
	galaxies := make([]GridCoord, 0, len(file))
	// find empty rows/cols
	cols := make([]bool, len(lines[0]))
	rows := make([]bool, len(lines))
	for row, line := range lines {
		for col, c := range line {
			if c == '#' {
				rows[row] = true
				cols[col] = true
				galaxies = append(galaxies, GridCoord{row: row, col: col})
			}
		}
	}

	for i, galaxy := range galaxies {
		for _, b := range galaxies[i:] {
			total += findDist(galaxy, b, rows, cols, 1)
			total2 += findDist(galaxy, b, rows, cols, 999999)

		}
	}

	fmt.Printf("Total: %v\n", total)
	fmt.Printf("Total2: %v\n", total2)

}

func findDist(a GridCoord, b GridCoord, rows []bool, cols []bool, extraGapSpaces int) uint64 {
	if a == b {
		return 0
	}
	count := uint64(0)

	xStep := 1
	if (a.col - b.col) > 0 {
		xStep = -1
	}
	yStep := 1
	if (a.row - b.row) > 0 {
		yStep = -1
	}
	for x := a.col; x != b.col; {
		count++
		x += xStep
		if !cols[x] {
			count += uint64(extraGapSpaces)
		}
	}
	for y := a.row; y != b.row; {
		count++
		y += yStep
		if !rows[y] {
			count += uint64(extraGapSpaces)
		}
	}
	return count
}
