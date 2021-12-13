package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 * Using the sum of frequencies for all
 * segments, and given the fact all digits
 * are present in the input, we can easily
 * resolve the digits.
 *   aaaa    =>    8888
 *  b    c   =>   6    8
 *  b    c   =>   6    8
 *   dddd    =>    7777
 *  e    f   =>   4    9
 *  e    f   =>   4    9
 *   gggg    =>    7777
 */

var digitSums = map[int]int{
	17: 1,
	34: 2,
	39: 3,
	30: 4,
	37: 5,
	41: 6,
	25: 7,
	49: 8,
	45: 9,
	42: 0,
}

var uniqueSegmentCountDigits = [4]int{
	1, 4, 7, 8,
}

func main() {
	input_path := flag.String("input", "./input", "The input data")
	entries := getFileContents(*input_path)
	part1 := countUniqueOutputs(entries)
	fmt.Println("Part 1:", part1)

	numbers := resolveOutputToNumbers(entries)
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("Part 2:", sum)
}

func getFileContents(filepath string) (entries []*Entry) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		entryTokens := strings.Split(line, "|")

		if len(entryTokens) != 2 {
			fmt.Printf("Line '%s' must contain exactly two coordinates. Skipping", line)
			continue
		}
		newEntry := &Entry{
			parseSignals(entryTokens[0]),
			parseSignals(entryTokens[1]),
		}
		entries = append(entries, newEntry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return entries
}

func parseSignals(token string) (signals []*Signal) {
	signal_tokens := strings.Split(token, " ")
	for _, v := range signal_tokens {
		signal := new(Signal)
		if v != "" {
			signal.Segments = make(map[rune]bool, 0)
			signal.activateSegmentsByString(v)
			signal.Digit = -1
			signals = append(signals, signal)
		}
	}
	return signals
}

func countUniqueOutputs(entries []*Entry) int {
	sum := 0
	for _, entry := range entries {
		for _, output := range entry.OutputValuePatterns {
			patternLength := output.GetSegmentCount()
			if patternLength == 2 ||
				patternLength == 3 ||
				patternLength == 4 ||
				patternLength == 7 {
				sum++
			}
		}
	}
	return sum
}

func resolveOutputToNumbers(entries []*Entry) (numbers []int) {
	numbers = make([]int, 0)
	for _, entry := range entries {
		inputRuneFrequency := make(map[rune]int, 0)
		for _, signal := range entry.Signals {
			signal.addToRuneCount(inputRuneFrequency)
		}
		outputNumber := 0
		for _, output := range entry.OutputValuePatterns {
			frequency := output.sumWithFrequencies(inputRuneFrequency)
			outputNumber = outputNumber*10 + digitSums[frequency]
		}
		numbers = append(numbers, outputNumber)
	}
	return numbers
}
