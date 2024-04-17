package wordcounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	count, err := Simple("line", "../../data/books/3/alices-adventures-in-wonderland.txt")
	assert.Nil(t, err)
	assert.Equal(t, count, 4)
}

func TestSimpleFileNotFound(t *testing.T) {
	count, err := Simple("line", "./not_found.txt")
	assert.Equal(t, err.Error(), "open ./not_found.txt: The system cannot find the file specified.")
	assert.Equal(t, count, 0)
}
