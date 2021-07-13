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
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		results = append(results, HashResults{f.Name(), fmt.Sprintf("%x", h.Sum(nil))})
		f.Close()
	}
	return results
}

func HashesAreEqual(hr1 []HashResults, hr2 []HashResults) (bool, []HashResults) {
	mismatches := []HashResults{}
	for i, v := range hr1 {
		if v.hash != hr2[i].hash {
			mismatches = append(mismatches, v, hr2[i])
		}
	}
	return len(mismatches) == 0, mismatches
}
