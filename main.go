package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ikovic/gotitles/hash"
	"gopkg.in/h2non/filetype.v1"
)

func getHash(path string) uint64 {
	file, _ := os.Open(path)
	hash, _ := hash.HashFile(file)
	return hash
}

func analyzeFile(path string, info os.FileInfo) {
	if !info.IsDir() {
		buf, _ := ioutil.ReadFile(path)
		head := buf[:261]

		if filetype.IsVideo(head) {
			hash := getHash(path)
			fmt.Println(hash)
		}
	}
}

func walk(path string, info os.FileInfo, err error) error {
	go analyzeFile(path, info)
	return nil
}

func main() {
	args := os.Args[1:]
	path := strings.Join(args, " ")

	filepath.Walk(path, walk)

	// don't stop
	var input string
	fmt.Scanln(&input)
}
