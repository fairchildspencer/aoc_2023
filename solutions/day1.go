package solutions

import (
	"strconv"
	"strings"
	"unicode"
)

func Day1Part1(input string) string {
	lines := strings.Split(input, "\n")

	var sum int
	var first string
	var last string
	for _, line := range lines {
		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			}
		}
		combined, _ := strconv.Atoi(first + last)
		sum += combined
		first = ""
		last = ""
	}

	return strconv.Itoa(sum)
}

func Day1Part2(input string) string {
	lines := strings.Split(input, "\n")
	var b strings.Builder
	for _, line := range lines {
		b.WriteString(replaceWordsWithDigits(line))
		b.WriteString("\n")
	}
	return Day1Part1(b.String())
}

func replaceWordsWithDigits(input string) string {
	var formattedBuilder strings.Builder
	var wholeBuilder strings.Builder
	for _, char := range input {
		wholeBuilder.WriteRune(char)
		if unicode.IsDigit(char) {
			formattedBuilder.WriteRune(char)
		} else {
			formattedBuilder.WriteString(getNumberSuffix(wholeBuilder.String()))
		}
	}
	return formattedBuilder.String()
}

func getNumberSuffix(str string) string {
	if strings.HasSuffix(str, "one") {
		return "1"
	} else if strings.HasSuffix(str, "two") {
		return "2"
	} else if strings.HasSuffix(str, "three") {
		return "3"
	} else if strings.HasSuffix(str, "four") {
		return "4"
	} else if strings.HasSuffix(str, "five") {
		return "5"
	} else if strings.HasSuffix(str, "six") {
		return "6"
	} else if strings.HasSuffix(str, "seven") {
		return "7"
	} else if strings.HasSuffix(str, "eight") {
		return "8"
	} else if strings.HasSuffix(str, "nine") {
		return "9"
	}
	return ""
}
