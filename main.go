package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/h2non/filetype.v1"
)

func walk(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		buf, _ := ioutil.ReadFile(path)
		head := buf[:261]

		if filetype.IsVideo(head) {
			fmt.Println(path)
		}
	}
	return nil
}

func main() {
	args := os.Args[1:]
	path := strings.Join(args, " ")

	filepath.Walk(path, walk)
}
