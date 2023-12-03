package utils

func CreateMatrix(lines []string) [][]rune {
	var matrix [][]rune
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	return matrix
}
