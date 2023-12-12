package days

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var numsRegex = regexp.MustCompile(`(\d+)`)

type Range struct {
	sourceStart uint64
	destStart   uint64
	length      uint64
	sourceEnd   uint64
}

type Part2Range struct {
	start uint64
	end   uint64
}

func Day5Main() {
	file, err := os.ReadFile("day5.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}

	rawBlocks := strings.Split(string(file), "\n\n")
	blocks := rawBlocks[1:]

	seedStrs := numsRegex.FindAllString(rawBlocks[0], -1)
	seeds := make([]uint64, 0, len(seedStrs))
	part2Seeds := make([]Part2Range, 0, len(seeds)/2)

	for _, seedStr := range seedStrs {
		num, err := strconv.ParseUint(seedStr, 10, 64)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, num)
	}
	for i := 0; i < len(seeds); i += 2 {
		part2Seeds = append(part2Seeds, Part2Range{start: seeds[i], end: seeds[i] + seeds[i+1]})
	}

	rounds := make([][]Range, 0, len(blocks)-1)
	for i := range blocks {
		// Skip the first two lines of the block to remove headers+whitespace
		block := strings.Split(blocks[i], "\n")[1:]

		rounds = append(rounds, make([]Range, 0, len(block)))
		for _, curRange := range block {
			partStrs := numsRegex.FindAllString(curRange, -1)
			parts := []uint64{0, 0, 0}

			for k := 0; k < 3; k++ {
				curNum, err := strconv.ParseUint(partStrs[k], 10, 64)
				if err != nil {
					panic(err)
				}
				parts[k] = curNum
			}
			rounds[i] = append(rounds[i], Range{sourceStart: parts[1], destStart: parts[0], length: parts[2], sourceEnd: parts[1] + parts[2]})
		}
	}
	var lowest uint64 = math.MaxUint64
	var lowest2 uint64 = math.MaxUint64

	for _, seed := range seeds {
		var next uint64 = seed

		for _, rnge := range rounds {
			next = findRange(next, &rnge)
		}

		if next < lowest {
			lowest = next
		}
	}

	next := part2Seeds

	for c, round := range rounds {
		prev := next
		next = findPart2Ranges(next, &round)
		fmt.Printf("Debug %v: %v => %v\n", c, prev, next)
	}
	for _, candidate := range next {
		if candidate.start < lowest2 {
			lowest2 = candidate.start
		}
	}

	fmt.Printf("==========\nPart1: %v\nPart2: %v\n", lowest, lowest2)
}

func findRange(seed uint64, ranges *[]Range) uint64 {
	for _, rnge := range *ranges {
		if seed >= rnge.sourceStart && seed <= rnge.sourceEnd {
			return (seed - rnge.sourceStart) + rnge.destStart
		}
	}
	return seed
}
func findPart2Ranges(in []Part2Range, ranges *[]Range) []Part2Range {
	out := make([]Part2Range, 0, 1024)
	for _, inRange := range in {
		covered := make([]Part2Range, 0, 1024)
		for _, rnge := range *ranges {
			start := Uint64Max(rnge.sourceStart, inRange.start)
			end := Uint64Min(rnge.sourceEnd, inRange.end)
			if start < end {
				fmt.Printf("Int: %v => %v, %v => %v,%v\n", inRange, start, end, findRange(start, ranges), findRange(end, ranges))
				out = append(out, Part2Range{start: findRange(start, ranges), end: findRange(end, ranges)})
				covered = append(covered, Part2Range{start: start, end: end})
			}
		}
		fmt.Printf("Out: %v\n", out)

		if len(covered) < 1 {
			out = append(out, inRange)
		} else {
			sort.Slice(covered, func(i, j int) bool { return covered[i].start < covered[j].start })
			var first = Part2Range{start: math.MaxInt64, end: math.MaxInt64}
			var last = Part2Range{start: math.MaxInt64, end: math.MaxInt64}
			for _, rng := range covered {
				if rng.start < first.start {
					first = rng
				}
				if rng.end > last.end {
					last = rng
				}
			}
			if inRange.start < first.start {
				out = append(out, Part2Range{start: findRange(inRange.start, ranges), end: findRange(first.start-1, ranges)})
				fmt.Printf("Out - prefix: %v\n", out)
			}
			if inRange.end > last.end {
				out = append(out, Part2Range{start: findRange(last.end+1, ranges), end: findRange(inRange.end, ranges)})
				fmt.Printf("Out - suffix: %v,%v,%v\n", inRange, last, out)
			}
			for i := 1; i < len(covered); i++ {
				start := covered[i-1].end + 1
				end := covered[i].start - 1
				if start < end {
					newPart := Part2Range{start: findRange(start, ranges), end: findRange(end, ranges)}
					out = append(out, newPart)
					fmt.Printf("Out - infix: %v, %v,%v => %v\n", newPart, start, end, out)
				}
			}
		}
	}
	return out
}
func Uint64Min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func Uint64Max(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}
