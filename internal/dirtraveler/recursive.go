package dirtraveler

import (
	"fmt"
	"os"
	"strings"
)

// Finds all files inside the given directory, and all sub-directories
// The full path to each file is returned and can be directly used to e.g. open the file
func Recursive(dirName string) ([]string, error) {
	dirItems, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	filesFound := []string{}
	for _, dirItem := range dirItems {
		itemName := dirItem.Name()
		path := dirName + "/" + itemName
		if strings.Contains(itemName, ".") {
			filesFound := append(filesFound, path)
			fmt.Sprintln("debug filesFound:", filesFound) // TODO: Remove debug once issue is fixed
		} else {
			newFiles, err := Recursive(path)
			if err != nil {
				return filesFound, err
			}
			filesFound = append(filesFound, newFiles...)
		}
	}
	return filesFound, nil
}
