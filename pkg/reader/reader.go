package reader

import (
	"bufio"
	"log"
	"os"
)

// ReadInputFile reads the file from filepath and returns an array
// of string with each element representing one line
func ReadInputFile(filepath string) (lines []string) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return lines
}
