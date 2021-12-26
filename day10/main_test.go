package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileContents(t *testing.T) {
	navChunks := getFileContents("./example")
	test := len(navChunks)
	assert.Equal(t, 10, test)
}

func TestFindFirstUnbalancedChunks(t *testing.T) {
	navChunks := getFileContents("./example")
	test := findFirstUnbalancedChunks(navChunks)
	assert.Equal(t, []rune("})])>"), test)
}

func TestCalculateUnbalancedScore(t *testing.T) {
	navChunks := getFileContents("./example")
	unbalChunks := findFirstUnbalancedChunks(navChunks)
	test := CalculateUnbalancedScore(unbalChunks)
	assert.Equal(t, 26397, test)
}
