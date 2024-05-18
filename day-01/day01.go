package day01

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"

	utils "github.com/srychert/advent-of-code-2023/internal/utils"
)

func assignRune(runes *[]rune, char rune) {
	switch len(*runes) {
	case 0, 1:
		*runes = append(*runes, char)
	case 2:
		(*runes)[1] = char
	}
}

func Day01() int {
	file := utils.GetInputFile("day-01")
	defer file.Close()

	sum := 0
	var m = map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		bound := 0
		runes := make([]rune, 0)

		for i, char := range line {
			if unicode.IsDigit(char) {
				bound = i + 1
				assignRune(&runes, char)
				continue
			}

			for k, v := range m {
				if strings.Contains(line[bound:i+1], k) {
					// move bound to last rune of matched substr
					// e.g. twone -> one
					bound = i
					assignRune(&runes, v)
				}
			}
		}

		if len(runes) == 1 { // if we have only one digit add a copy of it to form 2 digit number at the end
			runes = append(runes, runes[0])
		}

		i, err := strconv.Atoi(string(runes))
		utils.Check(err)
		sum += i
	}

	err := sc.Err()
	utils.Check(err)

	return sum
}
