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

func TestCalculateBasinSize(t *testing.T) {
	heightMap := getFileContents("./example")
	lp := findLowPoints(heightMap)
	basins := findBasinsFromLowPoints(lp, heightMap)
	test := getFinalTopNBasinSizes(3, basins)
	assert.Equal(t, 1134, test)
}
