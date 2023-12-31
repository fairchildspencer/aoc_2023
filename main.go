package main

import (
	"aoc2023/solutions"
	"aoc2023/utils"
	"fmt"
)

const day = 15
const part = 2

var input = fmt.Sprintf("inputs/day%d.txt", day)

func main() {
	inputStr := utils.ReadInput(input)

	var result string
	switch day {
	case 1:
		switch part {
		case 1:
			result = solutions.Day1Part1(inputStr)
		case 2:
			result = solutions.Day1Part2(inputStr)
		}
	case 2:
		switch part {
		case 1:
			result = solutions.Day2Part1(inputStr)
		case 2:
			result = solutions.Day2Part2(inputStr)
		}
	case 3:
		switch part {
		case 1:
			result = solutions.Day3Part1(inputStr)
		case 2:
			result = solutions.Day3Part2(inputStr)
		}
	case 4:
		switch part {
		case 1:
			result = solutions.Day4Part1(inputStr)
		case 2:
			result = solutions.Day4Part2(inputStr)
		}
	case 14:
		switch part {
		case 1:
			result = solutions.Day14Part1(inputStr)
		case 2:
			result = solutions.Day14Part2(inputStr)
		}
	case 15:
		switch part {
		case 1:
			result = solutions.Day15Part1(inputStr)
		case 2:
			result = solutions.Day15Part2(inputStr)
		}
	}

	fmt.Println(result)
	err := utils.CopyToClipboard(result)
	if err != nil {
		fmt.Println(err)
	}
}
