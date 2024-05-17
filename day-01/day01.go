package day01

import (
	"bufio"
	"strconv"
	"unicode"

	utils "github.com/srychert/advent-of-code-2023/internal/utils"
)

func Day01() int {
	file := utils.GetInputFile("day-01")
	defer file.Close()

	sum := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()        // GET the line string
		runes := make([]rune, 0) // initialize empty rune slice

		for _, char := range line {
			if !unicode.IsDigit(char) {
				continue
			}

			switch len(runes) {
			case 0, 1:
				runes = append(runes, char)
			case 2:
				runes[1] = char
			}
		}

		if len(runes) == 1 { // if we have only one digit add a copy of it to form 2 digit number at the end
			runes = append(runes, runes[0])
		}

		i, err := strconv.Atoi(string(runes)) // convert slice to integer
		utils.Check(err)
		sum += i
	}

	err := sc.Err()
	utils.Check(err)

	return sum
}
