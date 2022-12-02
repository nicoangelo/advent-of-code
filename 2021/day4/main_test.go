package main

import "testing"

const amountNumbersDrawn = 27
const winningScore = 4512
const lastWinningScore = 1924

func TestGetFileContents(t *testing.T) {
	numbers_drawn, _ := getFileContents("./example")
	if len(numbers_drawn) != amountNumbersDrawn {
		t.Fatalf("Amount of drawn numbers wrong, got: %d, want: %d", len(numbers_drawn), amountNumbersDrawn)
	}
}

func TestPart1(t *testing.T) {
	numbers_drawn, boards := getFileContents("./example")
	scores := getScores(numbers_drawn, boards)
	if scores[0] != winningScore {
		t.Fatalf("Score of winning board wrong, got: %d, want: %d", scores[0], winningScore)
	}
}

func TestPart2(t *testing.T) {
	numbers_drawn, boards := getFileContents("./example")
	scores := getScores(numbers_drawn, boards)
	if scores[len(scores)-1] != lastWinningScore {
		t.Fatalf("Score of winning board wrong, got: %d, want: %d", scores[len(scores)-1], lastWinningScore)
	}
}
