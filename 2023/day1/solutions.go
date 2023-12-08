package day1

import (
	"log/slog"
	"strings"
	"unicode"
)

var digitStrings = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func part1(lines []string) (sum int) {
	sum = 0
	for _, v := range lines {
		sum += getDigitCalibrationValues(v)
	}

	return sum
}

func getDigitCalibrationValues(line string) int {
	var first, last int
	hasFirst := false

	for _, r := range line {
		if !unicode.IsDigit(r) {
			continue
		}
		intVal := int(r - '0')

		if !hasFirst {
			first = intVal
			hasFirst = true
		}
		last = intVal
	}
	return (first * 10) + last
}

func part2(lines []string) (sum int) {
	sum = 0
	for _, v := range lines {
		cal := getAllCalibrationValues(v)
		sum += cal
	}

	return sum
}

func getAllCalibrationValues(line string) (sum int) {
	var first, last int
	firstPos, lastPos := len(line), -1

	for pos, r := range line {
		if !unicode.IsDigit(r) {
			continue
		}
		intVal := int(r - '0')

		if pos < firstPos {
			first = intVal
			firstPos = pos
		}
		last = intVal
		lastPos = pos
	}

	for num, ds := range digitStrings {
		currentDigit := num + 1
		pos := strings.Index(line, ds)
		rpos := strings.LastIndex(line, ds)

		if pos == -1 {
			continue
		}

		if pos < firstPos {
			first = currentDigit
			firstPos = pos
		}

		if rpos > lastPos {
			last = currentDigit
			lastPos = rpos
		}
	}
	sum = (first * 10) + last
	slog.Debug("All calibration values", "line", line, "first", first, "last", last, "firstPos", firstPos, "lastPos", lastPos, "sum", sum)
	return sum
}
