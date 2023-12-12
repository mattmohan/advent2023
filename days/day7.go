package days

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var day7Regex = regexp.MustCompile(`^([AKQJT1-9]{5}) (\d+)$`)

type Hand struct {
	cards  []uint8
	bid    uint16
	class  uint8
	class2 uint8
}

const (
	FIVE = iota
	FOUR
	FULL
	THREE
	TWO_PAIR
	PAIR
	HIGH
)

func Day7Main() {
	file, err := os.ReadFile("day7.txt")
	if err != nil {
		panic(fmt.Errorf("Got error: %w", err))
	}
	var total uint64 = 0
	lines := strings.Split(string(file), "\n")
	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		parts := day7Regex.FindStringSubmatch(line)[1:]
		cards := convertCards(parts[0])
		class := classifyCards(cards)
		class2 := classifyCards2(cards)
		bid, _ := strconv.ParseUint(parts[1], 10, 16)
		hands = append(hands, Hand{cards: cards, bid: uint16(bid), class: class, class2: class2})
	}
	sort.Slice(hands, func(i, j int) bool {
		cI := hands[i].class
		cJ := hands[j].class
		if cI == cJ {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] != hands[j].cards[k] {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
		}
		return cI > cJ
	})
	for i, hand := range hands {
		total += uint64(i+1) * uint64(hand.bid)
	}

	fmt.Printf("Total1 %v\n", total)
	total = 0
	sort.Slice(hands, func(i, j int) bool {
		cI := hands[i].class2
		cJ := hands[j].class2
		if cI == cJ {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] != hands[j].cards[k] {
					if hands[i].cards[k] == 9 {
						return true
					}
					if hands[j].cards[k] == 9 {
						return false
					}
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
		}
		return cI > cJ
	})
	for i, hand := range hands {
		total += uint64(i+1) * uint64(hand.bid)
	}
	fmt.Printf("Total2 %v\n", total)
}

func classifyCards(cards []uint8) uint8 {
	m := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 5; i++ {
		m[cards[i]]++
	}
	sort.Slice(m, func(i, j int) bool { return m[i] > m[j] })
	if m[0] == 5 {
		return FIVE
	} else if m[0] == 4 {
		return FOUR
	} else if m[0] == 3 && m[1] == 2 {
		return FULL
	} else if m[0] == 3 {
		return THREE
	} else if m[0] == 2 && m[1] == 2 {
		return TWO_PAIR
	} else if m[0] == 2 {
		return PAIR
	}

	return HIGH
}

func classifyCards2(cards []uint8) uint8 {
	m := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, card := range cards {
		m[card]++
	}

	// Pull the jokers out
	jokers := m[9]
	// swap the last value and the joker then drop the last element(now the joker)
	m[9] = m[12]
	m = m[:12]
	sort.Slice(m, func(i, j int) bool {
		return m[i] > m[j]
	})
	if (m[0] + jokers) == 5 {
		return FIVE
	} else if (m[0] + jokers) == 4 {
		return FOUR
	} else if ((m[0] + jokers) == 3) && (m[1] == 2) {
		return FULL
	} else if (m[0] + jokers) == 3 {
		return THREE
	} else if m[0] == 2 && m[1] == 2 {
		return TWO_PAIR
	} else if m[0]+jokers == 2 {
		return PAIR
	}

	return HIGH
}
func convertCards(hand string) []uint8 {
	cards := make([]uint8, 0, 5)
	for i := 0; i < 5; i++ {
		switch hand[i] {
		case 'A':
			cards = append(cards, 12)
		case 'K':
			cards = append(cards, 11)
		case 'Q':
			cards = append(cards, 10)
		case 'J':
			cards = append(cards, 9)
		case 'T':
			cards = append(cards, 8)
		default:
			cards = append(cards, hand[i]-'2')
		}
	}
	return cards
}
