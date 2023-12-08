package day2

import (
	"log/slog"
	"os"
	"testing"

	"github.com/nicoangelo/aoc-pkg/reader"
)

var testData []string

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
	want := 8
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 2286
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
