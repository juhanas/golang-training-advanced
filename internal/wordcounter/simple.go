package wordcounter

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

// Counts the number of words in the file found in the path.
// The counting is done with simple looping in a single thread.
func Simple(wordToFind, filePath string) (int, error) {
	slog.Debug(
		"start function: wordcounter.Simple",
		slog.String("wordToFind", wordToFind),
		slog.String("filePath", filePath),
	)

	file, err := os.Open(filePath)
	if err != nil {
		slog.Error(
			"error from os.Open in function: wordcounter.Simple",
			slog.String("error", err.Error()),
		)
		return 0, err
	}
	defer func() {
		file.Close()
	}()

	wordsFound := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count := strings.Count(line, wordToFind)
		wordsFound += count
	}
	if err := scanner.Err(); err != nil {
		slog.Error(
			"error from scanner in function: wordcounter.Simple",
			slog.String("error", err.Error()),
		)
		return 0, err
	}
	return wordsFound, nil
}
