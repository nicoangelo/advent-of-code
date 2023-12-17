package almanac

import (
	"cmp"
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type MapTable struct {
	conversions []ConversionEntry
	seeds       []Seed
}

func (t *MapTable) ResetEntries() {
	t.conversions = make([]ConversionEntry, 0)
}

func (t *MapTable) HasEntries() bool {
	return t.conversions != nil && len(t.conversions) > 0
}

func (t *MapTable) AddEntryFromLine(line string) {
	tokens := strings.Split(line, " ")
	if len(tokens) != 3 {
		panic("Line must contain exactly 3 tokens.")
	}
	destStart, _ := strconv.Atoi(tokens[0])
	sourceStart, _ := strconv.Atoi(tokens[1])
	rng, _ := strconv.Atoi(tokens[2])
	t.conversions = append(t.conversions, ConversionEntry{
		DestinationStart: destStart,
		SourceStart:      sourceStart,
		SourceEnd:        sourceStart + rng - 1,
	})
	sort.Slice(t.conversions, func(i, j int) bool { return t.conversions[i].SourceStart < t.conversions[j].SourceStart })
}

func (t *MapTable) TranslateSeeds() {
	for i := range t.seeds {
		t.processSeed(&t.seeds[i])
	}
	// for _, s := range t.seeds {
	// 	t.processSeed(&s)
	// }
	sort.Slice(t.seeds, func(i, j int) bool { return t.seeds[i].Start < t.seeds[j].Start })
	t.PrintSeeds()
}

func (t *MapTable) processSeed(s *Seed) {
	slog.Info("Processing seed", "start", s.Start, "end", s.End)
	for _, c := range t.conversions {
		translatedStart := c.TranslateNumber(s.Start)
		translatedEnd := c.TranslateNumber(s.End)
		// Range does not contain start or end
		if translatedStart == -1 && translatedEnd == -1 {
			continue
		}
		slog.Info("Seed translated", "start", s.Start, "end", s.End, "newStart", translatedStart, "newEnd", translatedEnd)
		if translatedStart != -1 {
			s.Start = translatedStart
		} else {
			slog.Info("Must split seed at start", "at", c.SourceEnd)
			splitSeed := s.SplitAtStart(c.SourceStart)
			t.processSeed(&splitSeed)
			s.Start = c.TranslateNumber(s.Start)
			t.AddSeed(splitSeed)
		}
		if translatedEnd != -1 {
			s.End = translatedEnd // does not update
		} else {
			slog.Info("Must split seed at end", "at", c.SourceEnd)
			splitSeed := s.SplitAtEnd(c.SourceEnd)
			t.processSeed(&splitSeed)
			s.End = c.TranslateNumber(s.End)
			t.AddSeed(splitSeed)
		}
		if s.End < s.Start {
			panic("Something's off")
		}
		break
	}
}

func (t *MapTable) GetMinimumSeed() int {
	return slices.MinFunc(t.seeds, func(a, b Seed) int { return cmp.Compare(a.Start, b.Start) }).Start
}

func (t *MapTable) SetSeeds(s []Seed) {
	t.seeds = s
	sort.Slice(t.seeds, func(i, j int) bool { return t.seeds[i].Start < t.seeds[j].Start })
}

func (t *MapTable) AddSeed(s Seed) {
	t.seeds = append(t.seeds, s)
}

func (t *MapTable) PrintSeeds() {
	fmt.Println("Seed values:")
	for _, s := range t.seeds {
		fmt.Println(s.Start, " - ", s.End)
	}
}

type Seed struct {
	Start int
	End   int
}

func (s *Seed) SplitAtEnd(n int) Seed {
	originalEnd := s.End
	s.End = n
	return Seed{Start: n + 1, End: originalEnd}
}

func (s *Seed) SplitAtStart(n int) Seed {
	originalStart := s.Start
	s.Start = n
	return Seed{Start: originalStart, End: n - 1}
}

type ConversionEntry struct {
	SourceStart      int
	SourceEnd        int
	DestinationStart int
}

func (e *ConversionEntry) Contains(input int) bool {
	return input >= e.SourceStart && input <= e.SourceEnd
}

func (e *ConversionEntry) TranslateNumber(input int) int {
	if !e.Contains(input) {
		return -1
	}

	return e.DestinationStart + (input - e.SourceStart)
}
