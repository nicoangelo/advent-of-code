package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileContents(t *testing.T) {
	heightMap := getFileContents("./example")
	test := len(heightMap.Values)
	assert.Equal(t, 5, test)
}

func TestCalculateRiskLevel(t *testing.T) {
	heightMap := getFileContents("./example")
	lp := findLowPoints(heightMap)
	test := calculateRiskLevel(lp)
	assert.Equal(t, 15, test)
}
