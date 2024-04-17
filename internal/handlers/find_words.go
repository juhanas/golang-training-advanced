package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juhanas/golang-training-advanced/internal/dirtraveler"
	"github.com/juhanas/golang-training-advanced/internal/wordcounter"
)

// Path in reference to where the program is run - now "main.go" in root
// Using global private variable to allow tests to change it
var dirPath = "./data"

// Gin route handler for finding a specific word in the data.
// The response (success or fail) will be written into the context.
// Expects query param 'word' (string) in the request url.
func FindWordHandler(c *gin.Context) {
	wordToFind := c.Query("word")

	if wordToFind == "" {
		c.String(http.StatusBadRequest, "missing query param: word")
		return
	}

	var wordsFound int
	var err error
	wordsFound, err = getWordsRecursively(wordToFind)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "Found word '%s' %d times", wordToFind, wordsFound)
}

func getWordsRecursively(wordToFind string) (int, error) {
	filePaths, _ := dirtraveler.Recursive(dirPath)

	var wordsFound = 0
	for _, filePath := range filePaths {
		words, err := wordcounter.Simple(wordToFind, filePath)
		if err != nil {
			return 0, err
		}
		wordsFound += words
	}
	return wordsFound, nil
}
