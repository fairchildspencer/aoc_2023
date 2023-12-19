package solutions

import (
	"strconv"
	"strings"
)

func Day4Part1(input string) string {
	cards := strings.Split(input, "\n")
	sum := 0
	for _, card := range cards {
		text := strings.Split(card, ":")
		if len(text) != 2 {
			panic("Invalid card")
		}

		numSets := strings.Split(text[1], "|")
		if len(numSets) != 2 {
			panic("Invalid card")
		}

		winningNumbers := strings.Split(numSets[0], " ")
		myNumbers := strings.Split(numSets[1], " ")

		winCount := 0
		for _, myNumber := range myNumbers {
			if myNumber == "" {
				continue
			}
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					if winCount == 0 {
						winCount++
					} else {
						winCount *= 2
					}
				}
			}
		}
		sum += winCount
	}

	return strconv.Itoa(sum)
}

func Day4Part2(input string) string {
	cards := strings.Split(input, "\n")
	totalCards := make(map[int]int)
	for i := range cards {
		totalCards[i] = 1
	}

	for i, card := range cards {
		text := strings.Split(card, ":")
		if len(text) != 2 {
			panic("Invalid card")
		}

		numSets := strings.Split(text[1], "|")
		if len(numSets) != 2 {
			panic("Invalid card")
		}

		winningNumbers := strings.Split(numSets[0], " ")
		myNumbers := strings.Split(numSets[1], " ")
		winCount := 0
		for _, myNumber := range myNumbers {
			if myNumber == "" {
				continue
			}
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					winCount++
				}
			}
		}

		for j := 1; j <= winCount; j++ {
			totalCards[i+j] += totalCards[i]
		}
	}

	sum := 0
	for _, card := range totalCards {
		sum += card
	}

	return strconv.Itoa(sum)
}
