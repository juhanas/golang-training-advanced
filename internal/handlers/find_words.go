package handlers

import (
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/juhanas/golang-training-advanced/internal/dirtraveler"
	"github.com/juhanas/golang-training-advanced/internal/wordcounter"
	"github.com/juhanas/golang-training-advanced/pkg/helpers"
	"golang.org/x/sync/errgroup"
)

// Path in reference to where the program is run - now "main.go" in root
// Using global private variable to allow tests to change it
var dirPath = "./data"

var withErrors = true // Should concurrent functions detect errors

// Gin route handler for finding a specific word in the data.
// The response (success or fail) will be written into the context.
// Expects query param 'word' (string) and 'concurrent' (bool) in the request url.
func FindWordHandler(c *gin.Context) {
	wordToFind := c.Query("word")
	concurrent := c.Query("concurrent")
	useConcurrent := concurrent == "true"

	slog.Debug(
		"incoming request: FindWordHandler",
		slog.String("wordToFind", wordToFind),
		"concurrent", concurrent,
	)

	if wordToFind == "" {
		c.String(http.StatusBadRequest, "missing query param: word")
		slog.Warn(
			"validation error: FindWordHandler",
			slog.String("wordToFind", wordToFind),
		)
		return
	}

	var wordsFound int
	var err error
	if useConcurrent {
		if withErrors {
			wordsFound, err = getWordsConcurrentlyWithError(wordToFind)
		} else {
			wordsFound, err = getWordsConcurrently(wordToFind)
		}
	} else {
		wordsFound, err = getWordsRecursively(wordToFind)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		slog.Error(
			"error: FindWordHandler",
			slog.String("wordToFind", wordToFind),
			"concurrent", concurrent,
			slog.String("error", err.Error()),
		)
		return
	}

	slog.Info(
		"successful request: FindWordHandler",
		slog.String("wordToFind", wordToFind),
		"concurrent", concurrent,
		slog.Int("wordsFound", wordsFound),
	)
	c.String(http.StatusOK, "Found word '%s' %d times with concurrent %t", wordToFind, wordsFound, useConcurrent)
}

func getWordsRecursively(wordToFind string) (int, error) {
	filePaths, err := dirtraveler.Recursive(dirPath)
	if err != nil {
		slog.Error(
			"error from dirtraveler.Recursive in function: handlers.getWordsRecursive",
			slog.String("wordToFind", wordToFind),
			slog.String("filePaths", strings.Join(filePaths, ",")),
		)
		return 0, err
	}

	var wordsFound = 0
	for _, filePath := range filePaths {
		words, err := wordcounter.Simple(wordToFind, filePath)
		if err != nil {
			slog.Error(
				"error from wordcounter.Simple in function: handlers.getWordsRecursively",
				slog.String("wordToFind", wordToFind),
				slog.String("error", err.Error()),
			)
			return 0, err
		}
		wordsFound += words
	}
	return wordsFound, nil
}

func getWordsConcurrently(wordToFind string) (int, error) {
	wgDir := new(sync.WaitGroup)
	wgWords := new(sync.WaitGroup)
	filesChan := make(chan string)
	countChan := make(chan int)

	wgDir.Add(1)
	go dirtraveler.Concurrent(dirPath, filesChan, wgDir)
	go helpers.CloseChan(filesChan, wgDir)

	for {
		filePath, more := <-filesChan
		if more {
			wgWords.Add(1)
			go wordcounter.Concurrent(wordToFind, filePath, countChan, wgWords)
		} else {
			break
		}
	}
	go helpers.CloseChan(countChan, wgWords)

	wordsFound := 0
	for {
		count, more := <-countChan
		if more {
			wordsFound += count
		} else {
			break
		}
	}

	return wordsFound, nil
}

func getWordsConcurrentlyWithError(wordToFind string) (int, error) {
	egDir := new(errgroup.Group)
	egWords := new(errgroup.Group)
	filesChan := make(chan string)
	countChan := make(chan int)
	errChan := make(chan error)

	egDir.Go(func() error {
		return dirtraveler.ConcurrentWithError(dirPath, filesChan, egDir)
	})
	go func() {
		err := egDir.Wait()
		close(filesChan)
		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	// Name the for loop. Name can be anything, but it must be unique.
loop:
	for {
		select {
		case err := <-errChan:
			if err != nil {
				close(countChan)
				return 0, err
			}
			// Break out of the outer statement with name "loop"
			break loop
		case filePath := <-filesChan:
			egWords.Go(func() error {
				if filePath != "" {
					return wordcounter.ConcurrentWithError(wordToFind, filePath, countChan)
				}
				return nil
			})
		}
	}
	errChan = make(chan error)
	go func() {
		err := egWords.Wait()
		close(countChan)
		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	wordsFound := 0
loop2:
	for {
		select {
		case err := <-errChan:
			if err != nil {
				return 0, err
			}
			break loop2
		case count := <-countChan:
			wordsFound += count
		}
	}

	return wordsFound, nil
}
