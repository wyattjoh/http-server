package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// isDir tests a filepath if it is a directory or not.
func isDir(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		return false
	}

	return d.IsDir()
}

// getDirectory returns the directory that should be served. This is sourced
// from the flags that are provided.
func getDirectory(dir string) (string, error) {
	if !isDir(dir) {
		return "", fmt.Errorf("%s is not a directory or does not exist", dir)
	}

	if !path.IsAbs(dir) {
		absdir, err := filepath.Abs(dir)
		if err != nil {
			return "", fmt.Errorf("%s cannot be resolved to an absolute directory: %s", dir, err.Error())
		}

		dir = absdir
	}

	// Clean up the filepath if we can and return it.
	return filepath.Clean(dir), nil
}
