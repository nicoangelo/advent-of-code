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

// ReadInputIntoStruct reads the file from filepath and uses the provided reader
// functions to transform the input into a single struct.
// Empty lines will use the next reader function
func ReadInputIntoStruct[T any](filepath string, readers ...func(string, *T)) (res *T) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	if len(readers) == 0 {
		log.Fatalln("Must provide at least one reader function.")
	}
	scanner := bufio.NewScanner(file)
	res = new(T)
	readerIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if readerIndex == len(readers)-1 {
				readerIndex = 0
			} else {
				readerIndex++
			}
			continue
		}
		readers[readerIndex](line, res)

	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return res
}
