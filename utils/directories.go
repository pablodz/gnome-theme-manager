package utils

import (
	"os"
)

func GetDirectories(path string) (out []string, err error) {

	// Use ReadDir to get a list of the directories in the path
	directories, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Iterate over the list of directories and print their names
	for _, directory := range directories {
		if directory.IsDir() {
			out = append(out, directory.Name())
		}
	}

	return out, nil
}
