package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juhanas/golang-training-advanced/internal/dirtraveler"
	"github.com/juhanas/golang-training-advanced/internal/wordcounter"
)

// Path in reference to where the program is run - now "main.go" in root
// Using global private variable to allow tests to change it
var dirPath = "./data"

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
		wordsFound, err = getWordsConcurrently(wordToFind)
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
	// TODO: Get all file paths concurrently

	// TODO: Count words concurrently

	wordsFound := 0
	// TODO: Aggregate word counts

	return wordsFound, nil
}
