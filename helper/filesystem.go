package helper

import (
	"fmt"
	"os"
)

// IsDirectory returns true if the path is a directory
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return fileInfo.IsDir(), err
	}

	return false, fmt.Errorf("No such directory: %s", path)
}
