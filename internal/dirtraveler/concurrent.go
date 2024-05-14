package dirtraveler

import (
	"os"
	"strings"
	"sync"
)

// Reads through a directory concurrently, by using goroutines and channels.
// This allows for a faster traversal of the directory with large datasets.
func Concurrent(dirName string, filesChan chan string, wg *sync.WaitGroup) error {
	dirItems, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	for _, dirItem := range dirItems {
		itemName := dirItem.Name()
		path := dirName + "/" + itemName
		if strings.Contains(itemName, ".") {
			if strings.Contains(itemName, ".txt") {
				filesChan <- path
			}
		} else {
			wg.Add(1)
			go Concurrent(path, filesChan, wg)
		}
	}
	wg.Done()
	return nil
}
