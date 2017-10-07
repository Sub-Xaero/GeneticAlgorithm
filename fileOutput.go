package ga

import (
	"os"
	"strconv"
	"strings"
)

// GetOutputFilenameGenerator closure that returns a func that increments a counter and returns filename and NotExist error
func GetOutputFilenameGenerator() func() (string, error) {
	filename := "output"
	filenum := 0
	extension := ".txt"
	return func() (string, error) {
		filenum++
		fullFileName := strings.Join([]string{filename, strconv.Itoa(filenum), extension}, "")
		_, err := os.Stat(fullFileName)
		return fullFileName, err
	}
}

// GetUnusedOutputFile returns the filename of a numbered output file that doesn't yet exist
func GetUnusedOutputFile() string {
	generateFilename := GetOutputFilenameGenerator()

	fullFileName, err := generateFilename()
	for !os.IsNotExist(err) {
		fullFileName, err = generateFilename()
	}
	return fullFileName
}
