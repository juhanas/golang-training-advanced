package wordcounter

import (
	"bufio"
	"os"
	"strings"
)

// Counts the number of words in the file found in the path.
// The counting is done with simple looping in a single thread.
func Simple(wordToFind, filePath string) (int, error) {
	file, err := os.Open(filePath)
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return 0, err
	}

	wordsFound := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count := strings.Count(line, wordToFind)
		wordsFound += count
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return wordsFound, nil
}
