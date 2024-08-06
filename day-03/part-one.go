package day03

import (
	"bufio"
	"strconv"
	"unicode"

	config "github.com/srychert/advent-of-code-2023/internal/config"
	utils "github.com/srychert/advent-of-code-2023/internal/utils"
)

type Coord struct {
	X, Y int
}

type CoordSE struct {
	S     Coord
	E     Coord
	Value int
}

func isAdjecent(ces *CoordSE, m *map[Coord]bool) bool {
	offsets := []Coord{
		{-1, -1}, // Top-left diagonal
		{-1, 0},  // Top
		{-1, 1},  // Top-right diagonal
		{0, -1},  // Left
		{0, 1},   // Right
		{1, -1},  // Bottom-left diagonal
		{1, 0},   // Bottom
		{1, 1},   // Bottom-right diagonal
	}

	for _, offset := range offsets {
		adjacentS := Coord{X: ces.S.X + offset.X, Y: ces.S.Y + offset.Y}
		if _, exists := (*m)[adjacentS]; exists {
			return true
		}

		adjacentE := Coord{X: ces.E.X + offset.X, Y: ces.E.Y + offset.Y}
		if _, exists := (*m)[adjacentE]; exists {
			return true
		}
	}
	return false
}

func PartOne(flags config.Flags) int {
	file := utils.GetInputFile("day-03", "one", flags.Ex)
	defer file.Close()

	lineIndex := 0
	sum := 0
	current := ""
	coord := Coord{-1, -1}
	coords := make([]CoordSE, 0)
	symbolsMap := make(map[Coord]bool)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		for pos, char := range line {
			if char == '.' {
				if current != "" {
					n, _ := strconv.Atoi(current)
					coords[len(coords)-1].Value = n
					coords[len(coords)-1].E = coord
					current = ""
				}

				continue
			}

			if unicode.IsDigit(char) {
				coord.X = lineIndex
				coord.Y = pos
				if current == "" {
					coords = append(coords, CoordSE{coord, coord, 0})
				}
				current += string(char)
				continue
			}

			symbolsMap[Coord{lineIndex, pos}] = true

			if current != "" {
				n, _ := strconv.Atoi(current)
				coords[len(coords)-1].Value = n
				coords[len(coords)-1].E = coord
				current = ""
			}
		}

		lineIndex++
	}

	for _, ces := range coords {
		if isAdjecent(&ces, &symbolsMap) {
			sum += ces.Value
		}
	}

	return sum
}
