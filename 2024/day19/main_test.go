package day19

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = &Towels{
	availablePatterns: map[string]int{"r": 1, "wr": 2, "b": 1, "g": 1, "bwu": 3, "rb": 2, "gb": 2, "br": 2},
	desiredPatterns: []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	},
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)
	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 6
	assert.Equal(t, want, res)
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 16
	assert.Equal(t, want, res)
}
