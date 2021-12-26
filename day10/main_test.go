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
	test, _ := analyseNavigationRows(navChunks)
	assert.Equal(t, []rune("})])>"), test)
}

func TestCalculateUnbalancedScore(t *testing.T) {
	navChunks := getFileContents("./example")
	unbalChunks, _ := analyseNavigationRows(navChunks)
	test := CalculateUnbalancedScore(unbalChunks)
	assert.Equal(t, 26397, test)
}

func TestCalculateMissingChunkScore(t *testing.T) {
	navChunks := getFileContents("./example")
	_, missingChunks := analyseNavigationRows(navChunks)
	test := CalculateMissingChunkMiddleScore(missingChunks)
	assert.Equal(t, 288957, test)
}
