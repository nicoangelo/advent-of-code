package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	input_path := flag.String("input", "./input", "The input data")
	getFileContents(*input_path)
}

func getFileContents(filepath string) (chunks [][]rune) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	chunks = make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		chunks = append(chunks, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return chunks
}
