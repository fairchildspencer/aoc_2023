package main

import (
	"aoc2023/utils"
	"fmt"
)

const day = 1
const part = 1
const input = "inputs/day1.txt"

func main() {
	inputStr := utils.ReadInput(input)

	var result string
	switch day {
	case 1:
		switch part {
		case 1:
			result = day1part1(inputStr)
		case 2:
			result = day1part2(inputStr)
		}
	}

	fmt.Println(result)
	err := utils.CopyToClipboard(result)
	if err != nil {
		fmt.Println(err)
	}
}
