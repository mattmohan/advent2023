package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var numRegex = regexp.MustCompile(`(\d+)`)

func Day4Main() {
	file, err := os.ReadFile("day4.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	total := 0
	lines := strings.Split(string(file), "\n")
	cardCount := 0
	cardCounts := make([]int, len(lines))

	// We start out with one of each card
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for i, line := range lines {
		// Add the count of the current card to the tally
		cardCount += cardCounts[i]

		// Find and split on the ':' and '|'
		firstSplit := strings.Index(line, ":")
		rest := line[firstSplit+1:]
		split := strings.IndexByte(rest, '|')

		// I'm too tired to parse the numbers myself ¯\_(ツ)_/¯
		winning := numRegex.FindAllStringSubmatch(rest[:split-1], -1)
		myNums := numRegex.FindAllStringSubmatch(rest[split+1:], -1)

		matches := 0
		// I considered maps, but the set is small enough that a loop is probably faster
		for _, win := range winning {
			for _, myNum := range myNums {
				if strings.Compare(win[1], myNum[1]) == 0 {
					matches += 1
				}
			}
		}
		if matches > 0 {
			// Bitshift to convert matches to points
			total += 1 << (matches - 1)

			// Handle winning copies of other cards for part2
			for j := i + 1; j <= i+matches && j < len(cardCounts); j++ {
				cardCounts[j] += cardCounts[i]
			}
		}
	}
	fmt.Printf("%v\n%v\n", total, cardCount)
}
