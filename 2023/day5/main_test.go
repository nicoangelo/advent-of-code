package day5

import (
	"log/slog"
	"os"
	"testing"

	"github.com/nicoangelo/aoc-pkg/reader"
)

var testData = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)

	testData = reader.ReadInputFile("./input_test")
	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 35
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 46
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
