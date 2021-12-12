package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var digitSegmentCount = map[int]int{
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
	0: 6,
}

var uniqueSegmentCountDigits = [4]int{
	1, 4, 7, 8,
}

type Entry struct {
	SignalPatterns      []*Signal
	OutputValuePatterns []*Signal
}

func (entry Entry) String() string {
	res := ""
	for _, v := range entry.SignalPatterns {
		res += v.String()
	}
	res += "|"
	for _, v := range entry.OutputValuePatterns {
		res += v.String()
	}
	return res
}

func (signal Signal) String() string {
	if signal.Digit != -1 {
		return fmt.Sprint(signal.Digit)
	}
	return "."
}

type Signal struct {
	Pattern string
	Digit   int
}

func main() {
	input_path := flag.String("input", "./input", "The input data")
	entries := getFileContents(*input_path)
	part1 := countUniqueOutputs(entries)
	fmt.Println("Part 1:", part1)
	resolveSignalsToDigits(entries)
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
		if v != "" {
			signals = append(signals, &Signal{v, -1})
		}
	}
	return signals
}

func countUniqueOutputs(entries []*Entry) int {
	sum := 0
	for _, entry := range entries {
		for _, output := range entry.OutputValuePatterns {
			patternLength := len(output.Pattern)
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

func resolveSignalsToDigits(entries []*Entry) int {
	for _, entry := range entries {
		for _, signal := range entry.SignalPatterns {
			if v, ok := getUniqueCandidate(signal.Pattern); ok {
				signal.Digit = v
			}
		}
		for _, output := range entry.OutputValuePatterns {
			if v, ok := getUniqueCandidate(output.Pattern); ok {
				output.Digit = v
			}
		}
	}
	return 0
}

func getUniqueCandidate(pattern string) (digit int, ok bool) {
	for _, digit := range uniqueSegmentCountDigits {
		if len(pattern) == digitSegmentCount[digit] {
			return digit, true
		}
	}
	return -1, false
}
