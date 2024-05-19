package day02

import (
	"bufio"
	"strconv"
	"strings"

	utils "github.com/srychert/advent-of-code-2023/internal/utils"
)

type Color string

const (
	RED   Color = "red"
	GREEN Color = "green"
	BLUE  Color = "blue"
)

func handleCubes(cubes string, color Color, minCubes *int) {
	before, found := strings.CutSuffix(cubes, " "+string(color))

	if found {
		n, _ := strconv.Atoi(before)

		if n > *minCubes {
			*minCubes = n
		}
	}
}

func Day02() int {
	file := utils.GetInputFile("day-02")
	defer file.Close()

	sum := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		g := strings.Split(line, ": ")

		minRed := 0
		minGreen := 0
		minBlue := 0

		rounds := strings.Split(g[1], "; ")

		for _, round := range rounds {
			r := strings.Split(round, ", ")
			for _, cubes := range r {
				handleCubes(cubes, RED, &minRed)
				handleCubes(cubes, GREEN, &minGreen)
				handleCubes(cubes, BLUE, &minBlue)
			}
		}

		sum += minRed * minGreen * minBlue
	}

	return sum
}
