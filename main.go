package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	day01 "github.com/srychert/advent-of-code-2023/day-01"
	day02 "github.com/srychert/advent-of-code-2023/day-02"
	day03 "github.com/srychert/advent-of-code-2023/day-03"
	config "github.com/srychert/advent-of-code-2023/internal/config"
	utils "github.com/srychert/advent-of-code-2023/internal/utils"
)

func main() {
	var day int
	var part string
	var ex bool

	flag.IntVar(&day, "day", 0, "Advent of code day (required). Accepted values: 1-25")
	flag.StringVar(&part, "part", "two", "Part of the puzzle. Accepted values: one, two")
	flag.BoolVar(&ex, "ex", false, "Use the example file for part.")
	flag.Parse()

	if !(utils.IsPartOne(part) || utils.IsPartTwo(part)) {
		log.Fatal("Invalid 'part' value")
	}

	flags := config.Flags{Day: day, Part: part, Ex: ex}

	switch day {
	case 1:
		fmt.Println(day01.Day01(flags))
	case 2:
		fmt.Println(day02.Day02(flags))
	case 3:
		fmt.Println(day03.Day03(flags))
	case 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		panic(fmt.Sprintf("There is no solution for day %d yet\n", day))
	default:
		fmt.Printf("Error: No solution for day %d\n", day)
		flag.PrintDefaults()
		os.Exit(1)
	}
}
