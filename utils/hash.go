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

// HashResults store the computed hash of a filename
type HashResults struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
}

// JSONHashResults are a json representation of the results
type JSONHashResults struct {
	Results []HashResults `json:"results"`
	Match   bool          `json:"matched"`
	Algo    string        `json:"algo"`
}

// GetHashesForFiles returns the hashes for a given set of files
func GetHashesForFiles(files []string, algo string) []HashResults {
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
	if len(hr1) != len(hr2) {
		fmt.Println("Number of files do not match")
		os.Exit(-1)
	}
	for i, v := range hr1 {
		if v.Hash != hr2[i].Hash {
			mismatches = append(mismatches, v, hr2[i])
		}
	}
	return len(mismatches) == 0, mismatches
}
