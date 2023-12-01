package main

import (
	"aoc2023/solutions"
	"aoc2023/utils"
	"fmt"
)

const day = 1
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
	}

	fmt.Println(result)
	err := utils.CopyToClipboard(result)
	if err != nil {
		fmt.Println(err)
	}
}
