package solutions

import (
	"aoc2023/utils"
	"strconv"
	"strings"
	"unicode"
)

func Day3Part1(input string) string {
	lines := strings.Split(input, "\n")
	matrix := utils.CreateMatrix(lines)

	sum := 0
	for i, row := range matrix {
		currNumber := ""
		isPartNumber := false
		for j, char := range row {
			if unicode.IsDigit(char) {
				currNumber += string(char)
				if !isPartNumber {
					isPartNumber = checkIsPartNumber(i, j, matrix)
				}
			} else if isPartNumber && currNumber != "" {
				intVal, _ := strconv.Atoi(currNumber)
				sum += intVal
				currNumber = ""
				isPartNumber = false
			} else {
				currNumber = ""
				isPartNumber = false
			}
		}
		if currNumber != "" {
			isPartNumber = isPartNumber || checkIsPartNumber(i, len(row)-1, matrix)
			if isPartNumber {
				intVal, _ := strconv.Atoi(currNumber)
				sum += intVal
			}
		}
	}

	return strconv.Itoa(sum)
}

func checkIsPartNumber(i int, j int, matrix [][]rune) bool {
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if k >= 0 && k < len(matrix) && l >= 0 && l < len(matrix[k]) {
				if k == i && l == j {
					continue
				}
				if !unicode.IsDigit(matrix[k][l]) && matrix[k][l] != '.' {
					return true
				}
			}
		}
	}
	return false
}

func Day3Part2(input string) string {
	lines := strings.Split(input, "\n")
	matrix := utils.CreateMatrix(lines)

	sum := 0
	starGears := make(map[Coordinate][]int)
	for i, row := range matrix {
		currNumber := ""
		var currStar *Coordinate
		for j, char := range row {
			if unicode.IsDigit(char) {
				currNumber += string(char)
				if currStar == nil {
					currStar = getStarCoordinates(i, j, matrix)
				}
			} else if currStar != nil && currNumber != "" {
				intVal, _ := strconv.Atoi(currNumber)
				starGears[*currStar] = append(starGears[*currStar], intVal)
				currNumber = ""
				currStar = nil
			} else {
				currNumber = ""
				currStar = nil
			}
		}
		if currNumber != "" && currStar != nil {
			intVal, _ := strconv.Atoi(currNumber)
			starGears[*currStar] = append(starGears[*currStar], intVal)
		}
	}

	for _, gears := range starGears {
		if len(gears) == 2 {
			sum += gears[0] * gears[1]
		}
	}

	return strconv.Itoa(sum)
}

func getStarCoordinates(i int, j int, matrix [][]rune) *Coordinate {
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if k >= 0 && k < len(matrix) && l >= 0 && l < len(matrix[k]) {
				if k == i && l == j {
					continue
				}
				if matrix[k][l] == '*' {
					return &Coordinate{x: k, y: l}
				}
			}
		}
	}
	return nil
}

type Coordinate struct {
	x int
	y int
}
