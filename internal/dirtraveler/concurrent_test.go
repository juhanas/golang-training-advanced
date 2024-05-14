package dirtraveler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrent(t *testing.T) {
	// TODO: Get files concurrently

	filesFound := []string{}
	// TODO: Append the files to the filesFound slice

	assert.Equal(t, 18, len(filesFound))
	assert.Contains(t, filesFound, "../../data/books/3/blue-castle.txt")
}
