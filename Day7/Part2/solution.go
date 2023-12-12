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

var cards = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

// var cards = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func isXOfAKind(count map[byte]byte, x byte) bool {
	for _, c := range count {
		if c == x {
			return true
		}
	}
	return false
}

func isFullHouse(count map[byte]byte) bool {
	return isXOfAKind(count, 2) && isXOfAKind(count, 3)
}

func isTwoPair(count map[byte]byte) bool {
	var pairs byte = 0
	for _, c := range count {
		if c == 2 {
			pairs++
		}
	}
	return pairs == 2
}

func replaceJoker(count *map[byte]byte) {
	numberOfJokers := (*count)[0]
	if numberOfJokers == 5 {
		// all cards are jokers
		(*count)[byte(len(cards)-1)] = 5
		delete(*count, 0)
		return
	}
	// find the highest count card other than the joker
	highestCount := byte(0)
	highestCountCard := byte(0)
	for k, c := range *count {
		if k != 0 && (c > highestCount || (c == highestCount && k > highestCountCard)) {
			highestCount = c
			highestCountCard = k
		}
	}
	(*count)[highestCountCard] += numberOfJokers
	delete(*count, 0)
}

func compareHands(a, b Hand) bool {
	var countA map[byte]byte = make(map[byte]byte)
	var countB map[byte]byte = make(map[byte]byte)
	for i := 0; i < len(a.cards); i++ {
		countA[a.cards[i]]++
		countB[b.cards[i]]++
	}
	replaceJoker(&countA)
	replaceJoker(&countB)
	if isXOfAKind(countA, 5) && !isXOfAKind(countB, 5) {
		return true
	} else if !isXOfAKind(countA, 5) && isXOfAKind(countB, 5) {
		return false
	} else if isXOfAKind(countA, 4) && !isXOfAKind(countB, 4) {
		return true
	} else if !isXOfAKind(countA, 4) && isXOfAKind(countB, 4) {
		return false
	} else if isFullHouse(countA) && !isFullHouse(countB) {
		return true
	} else if !isFullHouse(countA) && isFullHouse(countB) {
		return false
	} else if isXOfAKind(countA, 3) && !isXOfAKind(countB, 3) {
		return true
	} else if !isXOfAKind(countA, 3) && isXOfAKind(countB, 3) {
		return false
	} else if isTwoPair(countA) && !isTwoPair(countB) {
		return true
	} else if !isTwoPair(countA) && isTwoPair(countB) {
		return false
	} else if isXOfAKind(countA, 2) && !isXOfAKind(countB, 2) {
		return true
	} else if !isXOfAKind(countA, 2) && isXOfAKind(countB, 2) {
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

	for _, hand := range hands {
		fmt.Println(hand)
	}

	totalWinnings := 0
	for i := 0; i < len(hands); i++ {
		totalWinnings += hands[i].bet * (i + 1)
	}

	fmt.Println(totalWinnings)
}
