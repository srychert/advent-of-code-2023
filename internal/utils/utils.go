package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInputFile(dayDir string, part string, ex bool) *os.File {
	// Get the current working directory
	currentDir, err := os.Getwd()
	Check(err)

	filePath := "input"
	if part != "" {
		filePath += fmt.Sprintf("-%s", part)
	}
	if ex {
		filePath += "-ex"
	}
	filePath += ".txt"

	// Construct the input file path
	inputFilePath := filepath.Join(currentDir, dayDir, filePath)

	// Open the input file
	file, err := os.OpenFile(inputFilePath, os.O_RDONLY, os.ModePerm)
	Check(err)

	return file
}

func IsPartOne(part string) bool {
	return part == "one"
}

func IsPartTwo(part string) bool {
	return part == "two"
}
