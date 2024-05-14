package dirtraveler

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/juhanas/golang-training-advanced/pkg/helpers"
)

func TestConcurrent(t *testing.T) {
	filesChan := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go Concurrent("../../data", filesChan, wg)
	go helpers.CloseChan(filesChan, wg)

	filesFound := []string{}
	for {
		file, more := <-filesChan
		if more {
			filesFound = append(filesFound, file)
		} else {
			break
		}
	}
	assert.Equal(t, 18, len(filesFound))
	assert.Contains(t, filesFound, "../../data/books/3/blue-castle.txt")
}
