package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileContents(t *testing.T) {
	heightMap := getFileContents("./example")
	test := len(heightMap)
	assert.Equal(t, 5, test)
}

func TestCalculateRiskLevel(t *testing.T) {
	heightMap := getFileContents("./example")
	test := calculateRiskLevel(heightMap)
	assert.Equal(t, 15, test)
}
