package dirtraveler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursive(t *testing.T) {
	filesFound, err := Recursive("../../data")
	assert.Nil(t, err)
	assert.Equal(t, len(filesFound), 18)
	assert.Contains(t, filesFound, "../../data/books/3/blue-castle.txt")
}
