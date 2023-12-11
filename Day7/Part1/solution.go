package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards []byte
	bet   int
}

var cards = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func isXOfAKind(hand Hand, x int) bool {
	var count map[byte]int = make(map[byte]int)
	for i := 0; i < len(hand.cards); i++ {
		count[hand.cards[i]]++
	}
	for _, c := range count {
		if c == x {
			return true
		}
	}
	return false
}

func isFullHouse(hand Hand) bool {
	return isXOfAKind(hand, 2) && isXOfAKind(hand, 3)
}

func isTwoPair(hand Hand) bool {
	var count map[byte]int = make(map[byte]int)
	for i := 0; i < len(hand.cards); i++ {
		count[hand.cards[i]]++
	}
	pairs := 0
	for _, c := range count {
		if c == 2 {
			pairs++
		}
	}
	return pairs == 2
}

func compareHands(a, b Hand) bool {
	if isXOfAKind(a, 5) && !isXOfAKind(b, 5) {
		return true
	} else if !isXOfAKind(a, 5) && isXOfAKind(b, 5) {
		return false
	} else if isXOfAKind(a, 4) && !isXOfAKind(b, 4) {
		return true
	} else if !isXOfAKind(a, 4) && isXOfAKind(b, 4) {
		return false
	} else if isFullHouse(a) && !isFullHouse(b) {
		return true
	} else if !isFullHouse(a) && isFullHouse(b) {
		return false
	} else if isXOfAKind(a, 3) && !isXOfAKind(b, 3) {
		return true
	} else if !isXOfAKind(a, 3) && isXOfAKind(b, 3) {
		return false
	} else if isTwoPair(a) && !isTwoPair(b) {
		return true
	} else if !isTwoPair(a) && isTwoPair(b) {
		return false
	} else if isXOfAKind(a, 2) && !isXOfAKind(b, 2) {
		return true
	} else if !isXOfAKind(a, 2) && isXOfAKind(b, 2) {
		return false
	} else {
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] > b.cards[i] {
				return true
			} else if a.cards[i] < b.cards[i] {
				return false
			}
		}
	}
	return false
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	hands := make([]Hand, len(lines))

	for i, line := range lines {
		cardsString := strings.Split(line, " ")[0]
		bet, _ := strconv.Atoi(strings.Split(line, " ")[1])

		hands[i].bet = bet

		for _, card := range cardsString {
			for j, c := range cards {
				if card == rune(c) {
					hands[i].cards = append(hands[i].cards, byte(j))
				}
			}
		}
	}

	// sort hands
	sort.SliceStable(hands, func(i, j int) bool {
		return !compareHands(hands[i], hands[j])
	})

	totalWinnings := 0
	for i := 0; i < len(hands); i++ {
		totalWinnings += hands[i].bet * (i + 1)
	}

	fmt.Println(totalWinnings)
}
