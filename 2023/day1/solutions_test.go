package day1

import (
	"log/slog"
	"os"
	"testing"
)

var testData1 = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)
	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	res := part1(testData1)
	want := 142
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

var testData2 = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func TestPart2(t *testing.T) {
	res := part2(testData2)
	want := 281
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
