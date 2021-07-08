package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

type HashResults struct {
	filename string
	hash     string
}

func GetMD5ForFiles(files []string) []HashResults {
	results := []HashResults{}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		defer f.Close()
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			os.Exit(-1)
		}
		results = append(results, HashResults{f.Name(), fmt.Sprintf("%x", h.Sum(nil))})
	}
	return results
}
