package main

import "fmt"

type Signal struct {
	Segments map[rune]bool
	Digit    int
}

func (signal Signal) String() string {
	if signal.Digit != -1 {
		return fmt.Sprint(signal.Digit)
	}
	return "."
}

func (signal1 *Signal) GetDiff(signal2 *Signal) *Signal {
	diff := &Signal{}
	for rune, active := range signal1.Segments {
		if !(active && signal2.Segments[rune]) {
			diff.Segments[rune] = true
		}
	}
	return diff
}

func (signal *Signal) GetSegmentCount() int {
	sum := 0
	for _, segment := range signal.Segments {
		if segment {
			sum++
		}
	}
	return sum
}

func (signal *Signal) activateSegmentsByString(segments string) {
	for _, rune := range segments {
		signal.Segments[rune] = true
	}
}

func (signal *Signal) addToRuneCount(runeFrequencies map[rune]int) {
	for rune, active := range signal.Segments {
		if active {
			runeFrequencies[rune]++
		}
	}
}

func (signal *Signal) sumWithFrequencies(runeFrequencies map[rune]int) int {
	sum := 0
	for rune, active := range signal.Segments {
		if active {
			sum += runeFrequencies[rune]
		}
	}
	return sum
}
