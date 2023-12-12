package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6Main() {
	file, err := os.ReadFile("day6.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}

	lines := strings.Split(string(file), "\n")
	timeStrs := numRegex.FindAllString(lines[0], -1)
	distanceStrs := numRegex.FindAllString(lines[1], -1)

	times := make([]uint64, 0, len(timeStrs))
	for _, timeStr := range timeStrs {
		time, _ := strconv.ParseUint(timeStr, 10, 64)
		times = append(times, time)
	}
	distances := make([]uint64, 0, len(distanceStrs))
	for _, distanceStr := range distanceStrs {
		distance, _ := strconv.ParseUint(distanceStr, 10, 64)
		distances = append(distances, distance)
	}

	total := 1
	for i := range times {
		count := 0
		for time := uint64(0); time < times[i]; time++ {
			timeRemaining := times[i] - time
			distance := time * timeRemaining
			if distance > distances[i] {
				count++
			}
		}
		total *= count
	}

	fmt.Printf("Total %v", total)
}
