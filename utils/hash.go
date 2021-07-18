package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"log"
	"os"
)

type HashResults struct {
	filename string
	hash     string
}

func GetMD5ForFiles(files []string, algo string) []HashResults {
	results := []HashResults{}
	var h hash.Hash
	switch algo {
	case "md5sum":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "crc32":
		h = crc32.NewIEEE()
	default:
		log.Fatalf("Unknown hash algorithm: %s", algo)
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		results = append(results, HashResults{f.Name(), fmt.Sprintf("%x", h.Sum(nil))})
		h.Reset()
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
