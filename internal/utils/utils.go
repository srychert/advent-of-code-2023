package internal

import (
	"os"
	"path/filepath"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInputFile(dayDir string) *os.File {
	// Get the current working directory
	currentDir, err := os.Getwd()
	Check(err)

	// Construct the input file path
	inputFilePath := filepath.Join(currentDir, dayDir, "input.txt")

	// Open the input file
	file, err := os.OpenFile(inputFilePath, os.O_RDONLY, os.ModePerm)
	Check(err)

	return file
}
