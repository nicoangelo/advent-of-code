package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileContents(t *testing.T) {
	heightMap := getFileContents("./example")
	test := len(heightMap.Octopuses)
	assert.Equal(t, 10, test)
}

func TestStepOneIteration(t *testing.T) {
	heightMap := getFileContents("./example")
	test := heightMap.Step()
	assert.Equal(t, 0, test)
}

func TestStepTwoIterations(t *testing.T) {
	heightMap := getFileContents("./example")
	test := 0
	test += heightMap.Step()
	test += heightMap.Step()
	assert.Equal(t, 35, test)
}

func TestStepHundredIterations(t *testing.T) {
	heightMap := getFileContents("./example")
	test := 0
	for i := 1; i <= 100; i++ {
		test += heightMap.Step()
	}
	assert.Equal(t, 1656, test)
}

func TestStepUntilSynchronized(t *testing.T) {
	heightMap := getFileContents("./example")
	i := 0
	for i = 1; true; i++ {
		flashes := heightMap.Step()
		if flashes == 100 {
			break
		}
	}
	assert.Equal(t, 195, i)
}
