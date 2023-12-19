package solutions

import (
	"regexp"
	"strconv"
	"strings"
)

func Day15Part1(input string) string {
	parts := strings.Split(input, ",")

	sum := 0
	for _, part := range parts {
		sum += hash(part)
	}

	return strconv.Itoa(sum)
}

func hash(input string) int {
	currValue := 0
	for _, char := range input {
		currValue += int(char)
		currValue *= 17
		currValue = currValue % 256
	}
	return currValue
}

type Lens struct {
	label       string
	focalLength int
}

func Day15Part2(input string) string {
	instructions := strings.Split(input, ",")
	labelRe := regexp.MustCompile(`^\w+`)
	operatorRe := regexp.MustCompile(`[-=]`)
	focalLengthRe := regexp.MustCompile(`\d+$`)
	boxes := make(map[int][]Lens)

	for _, instruction := range instructions {
		label := labelRe.FindString(instruction)
		labelHash := hash(label)
		operator := operatorRe.FindString(instruction)

		var focalLength int
		if operator == "=" {
			focalLengthStr := focalLengthRe.FindString(instruction)
			focalLength, _ = strconv.Atoi(focalLengthStr)

			lenses := boxes[labelHash]
			foundExistingLens := false
			for i, lens := range lenses {
				if lens.label == label {
					boxes[labelHash][i].focalLength = focalLength
					foundExistingLens = true
				}
			}

			if !foundExistingLens {
				boxes[labelHash] = append(lenses, Lens{label, focalLength})
			}
		} else if operator == "-" {
			lenses := boxes[labelHash]
			for i, lens := range lenses {
				if lens.label == label {
					boxes[labelHash] = append(lenses[:i], lenses[i+1:]...)
				}
			}
		}
	}

	totalPower := 0
	for boxNumber, lenses := range boxes {
		for i, lens := range lenses {
			totalPower += (boxNumber + 1) * (i + 1) * lens.focalLength
		}
	}

	return strconv.Itoa(totalPower)
}
