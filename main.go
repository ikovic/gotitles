package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ikovic/gotitles/hash"
	"gopkg.in/h2non/filetype.v1"
)

func analyzeFile(path string, info os.FileInfo) {
	if !info.IsDir() {
		file, _ := os.Open(path)
		head := make([]byte, 261)
		file.Read(head)

		if filetype.IsVideo(head) {
			hash, _ := hash.HashFile(file)
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
