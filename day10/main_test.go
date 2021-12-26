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
