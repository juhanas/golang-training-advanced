//go:build mage

package main

import (
	"github.com/magefile/mage/sh"

	"github.com/juhanas/golang-training-advanced/pkg/buildhelpers"
)

var targetPath = "./bin/"

// Cleans the build dir and removes all data
func Clean() error {
	return sh.Rm(targetPath)
}

// Clean build dir, build executable and copy all data files to the build dir
func Build() error {
	err := Clean()
	if err != nil {
		return err
	}

	err = sh.Run("go", "build", "-o", targetPath)
	if err != nil {
		return err
	}

	return buildhelpers.CopyDataFolder("./data", targetPath+"data")
}

// Run all tests
func Test() error {
	return sh.Run("go", "test", "./...")
}

// Run only integration tests
func TestIntegration() error {
	return sh.Run("go", "test", "-run", "Integration", "./...")
}

// Run only unit tests
func TestUnit() error {
	return sh.Run("go", "test", "-short", "./...")
}
