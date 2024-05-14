package wordcounter

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/juhanas/golang-training-advanced/pkg/helpers"
)

func TestConcurrent(t *testing.T) {
	wordToFind := "line"

	countChan := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go Concurrent(wordToFind, "../../data/test.txt", countChan, wg)
	go helpers.CloseChan(countChan, wg)

	wordsFound := 0
	for {
		count, more := <-countChan
		if more {
			wordsFound += count
		} else {
			break
		}
	}
	assert.Equal(t, wordsFound, 4)
}
