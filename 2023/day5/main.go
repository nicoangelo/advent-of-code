package day5

import (
	"log"
	"log/slog"
	"strconv"
	"strings"

	"github.com/nicoangelo/advent-of-code-2023/day5/almanac"
	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day5/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

const (
	StateSeeds      = iota
	StateBlockStart = iota
	StateBlockData  = iota
)

func part1(lines []string) int {
	return ReadAndTranslateInput(lines, readSeedsAsSingleNumbers)
}

func part2(lines []string) int {
	return ReadAndTranslateInput(lines, readSeedsAsRanges)
}

func ReadAndTranslateInput(lines []string, seedReader func(string) []almanac.Seed) int {
	state := StateSeeds
	t := almanac.MapTable{}
	stepName := ""
	for _, l := range lines {
		if len(l) == 0 {
			if t.HasEntries() {
				slog.Info("Starting translation", "step", stepName)
				t.TranslateSeeds()
				t.ResetEntries()
			}
			// every empty line marks the beginning of a new block
			state = StateBlockStart
			continue
		}

		if state == StateSeeds && l[0:6] == "seeds:" {
			t.SetSeeds(seedReader(l[7:]))
			t.PrintSeeds()
		} else if state == StateBlockStart && (l[0] <= '0' || l[0] >= '9') {
			stepName = l[:len(l)-1]
			state = StateBlockData
			continue
		} else if state == StateBlockData {
			t.AddEntryFromLine(l)
		}
	}
	slog.Info("Starting translation", "step", stepName)
	t.TranslateSeeds()

	return t.GetMinimumSeed()
}

func readSeedsAsSingleNumbers(line string) (res []almanac.Seed) {
	tokens := strings.Split(line, " ")
	res = make([]almanac.Seed, len(tokens))
	for i, v := range tokens {
		res[i].Start, _ = strconv.Atoi(v)
		res[i].End = res[i].Start

	}
	return res
}

func readSeedsAsRanges(line string) (res []almanac.Seed) {
	tokens := strings.Split(line, " ")
	res = make([]almanac.Seed, 0)
	start := 0
	for i, v := range tokens {
		if i%2 == 1 {
			length, _ := strconv.Atoi(v)
			res = append(res, almanac.Seed{Start: start, End: start + length - 1})
		} else {
			start, _ = strconv.Atoi(v)
		}
	}
	return res
}
