package buildhelpers

import (
	"os"
	"strings"

	"github.com/juhanas/golang-training-advanced/internal/dirtraveler"
)

// Copies the data in the sourcePath to the targetPath.
// Copied data includes all files and the folder structure is maintained
func CopyDataFolder(sourcePath, targetPath string) error {
	os.Mkdir(targetPath, os.ModePerm)

	filePaths, err := dirtraveler.Recursive(sourcePath)
	if err != nil {
		return err
	}

	for _, filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		newFilePath := strings.Replace(filePath, sourcePath, targetPath, 1)
		idx := strings.LastIndex(newFilePath, "/")
		foldersPath := newFilePath[:idx+1]
		err = os.MkdirAll(foldersPath, os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(newFilePath, data, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
