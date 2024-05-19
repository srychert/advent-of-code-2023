package day02

import (
	"bufio"
	"strconv"
	"strings"

	utils "github.com/srychert/advent-of-code-2023/internal/utils"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

type Color string

const (
	RED   Color = "red"
	GREEN Color = "green"
	BLUE  Color = "blue"
)

var m = map[Color]int{
	RED:   MAX_RED,
	GREEN: MAX_GREEN,
	BLUE:  MAX_BLUE,
}

func handleCubes(cubes string, color Color) bool {
	before, found := strings.CutSuffix(cubes, " "+string(color))

	if found {
		n, _ := strconv.Atoi(before)

		return !(n > m[color])
	}

	return true
}

func Day02() int {
	file := utils.GetInputFile("day-02")
	defer file.Close()

	sum := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		g := strings.Split(strings.Replace(line, "Game ", "", 1), ": ")
		id, _ := strconv.Atoi(g[0])

		rounds := strings.Split(g[1], "; ")
		gameIsValid := true

	game:
		for _, round := range rounds {
			r := strings.Split(round, ", ")
			for _, cubes := range r {
				gameIsValid = handleCubes(cubes, RED) &&
					handleCubes(cubes, GREEN) &&
					handleCubes(cubes, BLUE)

				if !gameIsValid {
					break game
				}
			}
		}

		if gameIsValid {
			sum += id
		}
	}

	return sum
}
