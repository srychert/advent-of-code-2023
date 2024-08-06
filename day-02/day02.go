package day02

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	config "github.com/srychert/advent-of-code-2023/internal/config"
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

func Day02(flags config.Flags) int {
	if utils.IsPartOne(flags.Part) {
		log.Fatal("Part one not available")
	}
	if flags.Ex {
		log.Fatal("Example file not available")
	}
	file := utils.GetInputFile("day-02", "", false)
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
