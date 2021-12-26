package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var openingRunes = []rune("([{<")
var closingRunes = []rune(")]}>")
var unbalancedChunkScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}
var incompleteChunkScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
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
	unbalancedChunks, expectedChunks := analyseNavigationRows(navChunks)
	part1 := CalculateUnbalancedScore(unbalancedChunks)
	fmt.Println(part1)
	part2 := CalculateMissingChunkMiddleScore(expectedChunks)
	fmt.Println(part2)
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

func analyseNavigationRows(chunks [][]rune) (corruptingChunks []rune, expectedChunks [][]rune) {
	corruptingChunks = make([]rune, 0)
	expectedChunks = make([][]rune, 0)
	for _, line := range chunks {
		mustCloseChunks := []rune{}
		isCorrupt := false
		for _, rune := range line {
			if isOpeningRune(rune) {
				mustCloseChunks = append(mustCloseChunks, closingCounterparts[rune])
			}
			if isClosingRune(rune) {
				if rune == mustCloseChunks[len(mustCloseChunks)-1] {
					mustCloseChunks = mustCloseChunks[:len(mustCloseChunks)-1]
				} else {
					corruptingChunks = append(corruptingChunks, rune)
					isCorrupt = true
					break
				}
			}
		}
		if !isCorrupt {
			expectedChunks = append(expectedChunks, ReverseRuneArray(mustCloseChunks))
		}
	}
	return corruptingChunks, expectedChunks
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

func ReverseRuneArray(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func CalculateUnbalancedScore(unbalancedChunks []rune) int {
	score := 0
	for _, v := range unbalancedChunks {
		score += unbalancedChunkScore[v]
	}
	return score
}

func CalculateMissingChunkMiddleScore(missing [][]rune) int {
	scores := []int{}
	for _, line := range missing {
		score := 0
		for _, v := range line {
			score *= 5
			score += incompleteChunkScore[v]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	middleIndex := len(scores) / 2
	return scores[middleIndex]
}
