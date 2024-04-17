package dirtraveler

import (
	"log/slog"
	"os"
	"strings"
)

// Finds all files inside the given directory, and all sub-directories
// The full path to each file is returned and can be directly used to e.g. open the file
func Recursive(dirName string) ([]string, error) {
	slog.Debug(
		"start function: dircounter.Recursive",
		slog.String("dirName", dirName),
	)

	dirItems, err := os.ReadDir(dirName)
	if err != nil {
		slog.Error(
			"error from os.ReadDir in function: dircounter.Recursive",
			slog.String("dirName", dirName),
			slog.String("error", err.Error()),
		)
		return nil, err
	}
	slog.Info(
		"debug1 in function: dircounter.Recursive",
		slog.Int("dirItems.length", len(dirItems)),
	)

	filesFound := []string{}
	for _, dirItem := range dirItems {
		itemName := dirItem.Name()
		path := dirName + "/" + itemName
		slog.Info(
			"debug2 in function: dircounter.Recursive",
			slog.String("itemName", itemName),
			slog.String("path", path),
		)
		if strings.Contains(itemName, ".") {
			filesFound := append(filesFound, path)
			slog.Info(
				"debug3 in function: dircounter.Recursive",
				slog.String("filesFound", strings.Join(filesFound, ",")),
			)
		} else {
			newFiles, err := Recursive(path)
			if err != nil {
				slog.Error(
					"error from dirtraveler.Recursive in function: dirtraveler.Recursive",
					slog.String("error", err.Error()),
				)
				return filesFound, err
			}
			filesFound = append(filesFound, newFiles...)
		}
	}
	return filesFound, nil
}
