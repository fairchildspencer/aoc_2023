package solutions

import (
	"regexp"
	"strconv"
	"strings"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

var blueRe = regexp.MustCompile(`\d+ blue`)
var redRe = regexp.MustCompile(`\d+ red`)
var greenRe = regexp.MustCompile(`\d+ green`)
var numRe = regexp.MustCompile(`\d+`)

func Day2Part1(input string) string {
	games := strings.Split(input, "\n")
	sum := 0
	for _, game := range games {
		blues := blueRe.FindAllString(game, -1)
		reds := redRe.FindAllString(game, -1)
		greens := greenRe.FindAllString(game, -1)

		bluesValid := checkValidColor(blues, maxBlue)
		redsValid := checkValidColor(reds, maxRed)
		greensValid := checkValidColor(greens, maxGreen)

		if bluesValid && redsValid && greensValid {
			id := getGameId(game)
			sum += id
		}
	}

	return strconv.Itoa(sum)
}

func Day2Part2(input string) string {
	games := strings.Split(input, "\n")

	sum := 0
	for _, game := range games {
		blues := blueRe.FindAllString(game, -1)
		reds := redRe.FindAllString(game, -1)
		greens := greenRe.FindAllString(game, -1)

		bluesMax := getMaxColorValue(blues)
		redsMax := getMaxColorValue(reds)
		greensMax := getMaxColorValue(greens)

		power := bluesMax * redsMax * greensMax
		sum += power
	}

	return strconv.Itoa(sum)
}

func checkValidColor(colorValues []string, maxValue int) bool {
	for _, colorValue := range colorValues {
		values := numRe.FindAllString(colorValue, -1)
		for _, value := range values {
			intVal, _ := strconv.Atoi(value)
			if intVal > maxValue {
				return false
			}
		}
	}

	return true
}

func getGameId(game string) int {
	values := numRe.FindAllString(game, 1)
	intVal, _ := strconv.Atoi(values[0])
	return intVal
}

func getMaxColorValue(colorValues []string) int {
	biggest := 0
	for _, colorValue := range colorValues {
		values := numRe.FindAllString(colorValue, -1)
		for _, value := range values {
			intVal, _ := strconv.Atoi(value)
			if intVal > biggest {
				biggest = intVal
			}
		}
	}

	return biggest
}
