package wordcounter

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Counts the number of words in the file found in the path.
// The counting is done in separate goroutines to increase performance with large datasets.
func Concurrent(wordToFind, filePath string, countChan chan int, wgCount *sync.WaitGroup) error {
	defer wgCount.Done()
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	lineChan := make(chan string)

	go getLines(file, lineChan)

	for {
		line, more := <-lineChan
		if more {
			if count := strings.Count(line, wordToFind); count > 0 { // Without this if, execution takes a lot of extra time
				countChan <- count
			}
		} else {
			break
		}
	}
	return nil
}

// Counts the number of words in the file found in the path.
// The counting is done in separate goroutines to increase performance with large datasets.
func ConcurrentWithError(wordToFind, filePath string, countChan chan int) error {
	eg := new(errgroup.Group)
	errChan := make(chan error)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	lineChan := make(chan string)

	eg.Go(func() error {
		return getLines(file, lineChan)
	})
	go func() {
		err := eg.Wait()
		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()
	for {
		select {
		case err := <-errChan:
			return err // Returns nil or error if it happened
		case line := <-lineChan:
			if count := strings.Count(line, wordToFind); count > 0 { // Without this if, execution takes a lot of time
				countChan <- count
			}
		}
	}
}

func getLines(file *os.File, lineChan chan string) error {
	defer close(lineChan)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "failure" {
			return errors.New("error happened when reading file")
		}
		lineChan <- line
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
