package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var openingRunes = []rune("([{<")
var closingRunes = []rune(")]}>")
var chunkScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}
var closingCounterparts = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	input_path := flag.String("input", "./input", "The input data")
	navChunks := getFileContents(*input_path)
	unbalancedChunks := findFirstUnbalancedChunks(navChunks)
	part1 := CalculateUnbalancedScore(unbalancedChunks)
	fmt.Println(part1)
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

func findFirstUnbalancedChunks(chunks [][]rune) (unbalancedChunks []rune) {
	unbalancedChunks = make([]rune, 0)
	for _, line := range chunks {
		mustCloseChunks := []rune{}
		for _, rune := range line {
			if isOpeningRune(rune) {
				mustCloseChunks = append(mustCloseChunks, closingCounterparts[rune])
			}
			if isClosingRune(rune) {
				if rune == mustCloseChunks[len(mustCloseChunks)-1] {
					mustCloseChunks = mustCloseChunks[:len(mustCloseChunks)-1]
				} else {
					unbalancedChunks = append(unbalancedChunks, rune)
					break
				}
			}
		}
	}
	return unbalancedChunks
}

func isOpeningRune(rune rune) bool {
	for _, v := range openingRunes {
		if v == rune {
			return true
		}
	}
	return false
}

func isClosingRune(rune rune) bool {
	for _, v := range closingRunes {
		if v == rune {
			return true
		}
	}
	return false
}

func CalculateUnbalancedScore(unbalancedChunks []rune) int {
	score := 0
	for _, v := range unbalancedChunks {
		score += chunkScore[v]
	}
	return score
}
