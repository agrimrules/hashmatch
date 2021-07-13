package utils

import (
	"fmt"
	"os"

	"github.com/karrick/godirwalk"
)

func IsDirectory(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return fi.IsDir()
}

func ReturnFilesInFolder(folder string) []string {
	files := []string{}
	godirwalk.Walk(folder, &godirwalk.Options{
		Callback: func(osPathname string, directoryEntry *godirwalk.Dirent) error {
			if directoryEntry.IsRegular() {
				files = append(files, osPathname)
			}
			return nil
		}, Unsorted: false,
	})
	return files
}
