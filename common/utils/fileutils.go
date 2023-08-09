package utils

import (
	"os"
	"sort"
)

func GetAllFile(path string) (files []string, err error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range dir {
		files = append(files, file.Name())
	}

	sort.Strings(files)
	return
}
