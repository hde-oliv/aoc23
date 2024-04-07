package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Combination int

const (
	Empty Combination = iota
	FiveKind
	FourKind
	FullHouse
	ThreeKind
	TwoPair
	OnePair
	HighCard
)

var values = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
}

var values2 = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"T": 4,
	"9": 5,
	"8": 6,
	"7": 7,
	"6": 8,
	"5": 9,
	"4": 10,
	"3": 11,
	"2": 12,
	"J": 13,
}

type Hand struct {
	Cards []string
	Bid   int
	Comb  Combination
}

func count(slice []string, str string) int {
	count := 0
	for _, s := range slice {
		if s == str {
			count++
		}
	}
	return count
}

func parseInputLine(line string) *Hand {
	s := strings.Split(line, " ")

	cards := strings.Split(s[0], "")
	bid, _ := strconv.Atoi(s[1])

	return &Hand{Cards: cards, Bid: bid}
}

func getCombination(cards []string) Combination {
	t := make([]string, len(cards))
	copy(t, cards)

	t = slices.Compact(t) // this modifies the slice

	if len(t) == 5 {
		return HighCard
	}

	if len(t) == 4 {
		return OnePair
	}

	if len(t) == 3 {
		if (count(cards, t[0]) == 3) ||
			(count(cards, t[1]) == 3) ||
			(count(cards, t[2]) == 3) {
			return ThreeKind
		}
		return TwoPair
	}

	if len(t) == 2 {
		if (count(cards, t[0]) == 4) ||
			(count(cards, t[1]) == 4) {
			return FourKind
		}
		return FullHouse
	}

	if len(t) == 1 {
		return FiveKind
	}

	return Empty
}

func categorizeHands(hands []Hand) []Hand {
	n := make([]Hand, 0)

	for _, v := range hands {
		tmp := make([]string, len(v.Cards))
		copy(tmp, v.Cards)

		slices.SortFunc(tmp, func(a, b string) int {
			return cmp.Compare[int](values[a], values[b])
		})

		v.Comb = getCombination(tmp)

		n = append(n, v)
	}

	return n
}

func categorizeHands2(hands []Hand) []Hand {
	n := make([]Hand, 0)

	for _, v := range hands {
		tmp := make([]string, len(v.Cards))
		copy(tmp, v.Cards)

		slices.SortFunc(tmp, func(a, b string) int {
			return cmp.Compare[int](values2[a], values2[b])
		})

		v.Comb = getCombination2(tmp)

		n = append(n, v)
	}

	return n
}

func getCombination2(cards []string) Combination {
	t := make([]string, len(cards))
	copy(t, cards)

	t = slices.Compact(t) // this modifies the slice

	if len(t) == 5 {
		if slices.Contains(t, "J") {
			return OnePair
		}

		return HighCard
	}

	if len(t) == 4 {
		if slices.Contains(t, "J") {
			return ThreeKind
		}

		return OnePair
	}

	if len(t) == 3 {
		t0 := count(cards, t[0])
		t1 := count(cards, t[1])
		t2 := count(cards, t[2])

		if t0 == 3 || t1 == 3 || t2 == 3 {
			if slices.Contains(t, "J") {
				return FourKind
			}

			return ThreeKind
		}

		if slices.Contains(t, "J") {
			if count(cards, "J") == 1 {
				return FullHouse
			}

			return FourKind
		}

		return TwoPair
	}

	if len(t) == 2 {
		t0 := count(cards, t[0])
		t1 := count(cards, t[1])

		if slices.Contains(t, "J") {
			return FiveKind
		}

		if t0 == 4 || t1 == 4 {
			return FourKind
		}

		return FullHouse
	}

	if len(t) == 1 {
		return FiveKind
	}

	return Empty
}

func sortHands(hands []Hand) []Hand {
	t := make([]Hand, len(hands))
	copy(t, hands)

	slices.SortFunc(t, func(a, b Hand) int {
		if n := cmp.Compare[int](int(a.Comb), int(b.Comb)); n != 0 {
			return n
		}

		for i := 0; i < 5; i++ {
			if n := cmp.Compare[int](values[a.Cards[i]], values[b.Cards[i]]); n != 0 {
				return n
			}
		}
		return 0
	})

	return t
}

func sortHands2(hands []Hand) []Hand {
	t := make([]Hand, len(hands))
	copy(t, hands)

	slices.SortFunc(t, func(a, b Hand) int {
		if n := cmp.Compare[int](int(a.Comb), int(b.Comb)); n != 0 {
			return n
		}

		for i := 0; i < 5; i++ {
			if n := cmp.Compare[int](values2[a.Cards[i]], values2[b.Cards[i]]); n != 0 {
				return n
			}
		}
		return 0
	})

	return t
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hands := make([]Hand, 0)
	hands2 := make([]Hand, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		hand := parseInputLine(str)
		hands = append(hands, *hand)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	copy(hands2, hands)

	// Part One
	hands = categorizeHands(hands)
	hands = sortHands(hands)

	var result uint64
	for i, v := range hands {
		m := 1000 - i

		result += uint64(v.Bid * m)
	}

	fmt.Println(result)

	// Part Two
	result = 0
	hands2 = categorizeHands2(hands)
	hands2 = sortHands2(hands2)

	for i, v := range hands2 {
		m := 1000 - i

		result += uint64(v.Bid * m)
	}

	fmt.Println(result)
}
