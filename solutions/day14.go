package solutions

import (
	"aoc2023/utils"
	"strconv"
	"strings"
)

func Day14Part1(input string) string {
	lines := strings.Split(input, "\n")
	matrix := utils.CreateMatrix(lines)

	totalLoad := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'O' {
				totalLoad += calculateLoad(matrix, i, j)
			}
		}
	}

	return strconv.Itoa(totalLoad)
}

func calculateLoad(matrix [][]rune, i int, j int) int {
	restingIndex := calculateRestingPlace(matrix, i, j)
	return len(matrix) - restingIndex
}

func calculateRestingPlace(matrix [][]rune, i int, j int) int {
	stackCount := 0
	for k := i - 1; k >= 0; k-- {
		if matrix[k][j] == '#' {
			return k + 1 + stackCount
		} else if matrix[k][j] == 'O' {
			stackCount++
		}
	}

	return 0 + stackCount
}

func Day14Part2(input string) string {
	return ""
}
