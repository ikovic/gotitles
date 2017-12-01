package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/oz/osdb"
	"gopkg.in/h2non/filetype.v1"
)

var osdbClient, _ = osdb.NewClient()
var languages = []string{"hrv"}

func searchSubtitles(path string, languages []string) {
	subs, _ := osdbClient.FileSearch(path, languages)

	if len(subs) == 0 {
		return
	}

	directory := filepath.Dir(path)
	fullPath := directory + subs[0].SubFileName

	if err := osdbClient.DownloadTo(&subs[0], fullPath); err != nil {
		fmt.Printf("\nError saving %v", fullPath)
	} else {
		fmt.Printf("\nSaved %s", fullPath)
	}
}

func analyzeFile(path string, info os.FileInfo) {
	if !info.IsDir() {
		file, _ := os.Open(path)
		head := make([]byte, 261)
		file.Read(head)

		if filetype.IsVideo(head) {
			searchSubtitles(path, languages)
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

	osdbClient.LogIn("", "", "")

	filepath.Walk(path, walk)

	// don't stop
	var input string
	fmt.Scanln(&input)
}
