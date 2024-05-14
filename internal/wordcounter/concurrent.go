package wordcounter

import (
	"bufio"
	"errors"
	"os"
	"sync"
)

// Counts the number of words in the file found in the path.
// The counting is done in separate goroutines to increase performance with large datasets.
func Concurrent(wordToFind, filePath string, countChan chan int, wgCount *sync.WaitGroup) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// TODO: Get all lines concurrently

	// TODO: Count words on each line

	// TODO: Indicate that the operation is done

	return nil
}

func getLines(file *os.File, lineChan chan string) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "failure" {
			return errors.New("error happened when reading file")
		}
		// TODO: Add line to channel
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	// TODO: Close channel
	return nil
}
