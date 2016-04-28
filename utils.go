package main

import (
	"fmt"
	"os"
)

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return fileInfo.IsDir(), err
	}

	return false, fmt.Errorf("No such directory: %s", path)
}
