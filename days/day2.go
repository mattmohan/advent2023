package days

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var gameExp = regexp.MustCompile(`^Game (\d+):`)
var roundExp = regexp.MustCompile(`[^;]+`)
var playExp = regexp.MustCompile(`(\d+) (red|green|blue)`)

func Day2Main() {
	file, err := os.ReadFile("day2.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	lines := strings.Split(string(file), "\n")
	total := 0
	sumPower := 0

	rgbLimits := []int{12, 13, 14}
	mapColourNamesToIdx := map[string]int{"red": 0, "green": 1, "blue": 2}

	for i, line := range lines {
		gameStrs := gameExp.FindStringSubmatch(lines[i])
		game, err := strconv.Atoi(gameStrs[1])
		if err != nil {
			panic(fmt.Errorf("Failed to parse game from: %v, %w", lines[i], err))
		}

		roundsString := line[len(gameStrs[0])+1:]

		rounds := roundExp.FindAllStringSubmatch(roundsString, -1)

		valid := true
		rgbMaximums := []int{0, 0, 0}
		for _, round := range rounds {
			plays := playExp.FindAllStringSubmatch(round[0], -1)
			for _, play := range plays {
				count, _ := strconv.Atoi(play[1])
				idx := mapColourNamesToIdx[play[2]]
				if count > rgbLimits[idx] {
					valid = false
				}
				rgbMaximums[idx] = IntMin(rgbMaximums[idx], count)
			}
		}
		if valid {
			total += game
		}
		sumPower += rgbMaximums[0] * rgbMaximums[1] * rgbMaximums[2]
	}
	fmt.Printf("Total: %d, SumPower: %d", total, sumPower)
}

func IntMin(x, y int) int {
	if x > y {
		return x
	}
	return y
}
