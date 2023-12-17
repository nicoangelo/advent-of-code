package day6

import (
	"log/slog"
	"os"
	"testing"
)

var testData = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
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
	want := 288
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 71503
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
