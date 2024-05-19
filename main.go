package main

import (
	"flag"
	"fmt"
	"os"

	day01 "github.com/srychert/advent-of-code-2023/day-01"
	day02 "github.com/srychert/advent-of-code-2023/day-02"
)

func main() {
	var day int

	flag.IntVar(&day, "day", 0, "Advent of code day (required). Accepted values: 1-25")
	flag.Parse()

	switch day {
	case 1:
		fmt.Println(day01.Day01())
	case 2:
		fmt.Println(day02.Day02())
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		panic(fmt.Sprintf("There is no solution for day %d yet\n", day))
	default:
		fmt.Printf("Error: No solution for day %d\n", day)
		flag.PrintDefaults()
		os.Exit(1)
	}
}
