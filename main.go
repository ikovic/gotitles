package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/oz/osdb"
	"github.com/urfave/cli"
	"gopkg.in/h2non/filetype.v1"
)

// AppName is the application name, used for naming a folder where subs will be placed
const AppName = "gotitles"

var osdbClient, _ = osdb.NewClient()
var languages = []string{"hrv"}
var wg sync.WaitGroup

func createAppDirectory(path string) {
	appDirectory := filepath.Join(path, AppName)
	if _, err := os.Stat(appDirectory); os.IsNotExist(err) {
		os.Mkdir(appDirectory, os.ModePerm)
	}
}

func searchSubtitles(path string, languages []string) {
	subs, _ := osdbClient.FileSearch(path, languages)

	if len(subs) == 0 {
		return
	}

	directory := filepath.Dir(path)
	fullPath := filepath.Join(directory, AppName, subs[0].SubFileName)
	createAppDirectory(directory)

	if err := osdbClient.DownloadTo(&subs[0], fullPath); err != nil {
		fmt.Printf("\nError saving %v", fullPath)
	} else {
		fmt.Printf("\nSaved %s", fullPath)
	}
}

func analyzeFile(path string, info os.FileInfo) {
	defer wg.Done()
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
	wg.Add(1)
	go analyzeFile(path, info)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = "Download subtitles for all movie files in path"
	app.Action = func(c *cli.Context) error {
		path := c.Args().Get(0)
		osdbClient.LogIn("", "", "")
		filepath.Walk(path, walk)
		return nil
	}

	app.Run(os.Args)

	wg.Wait()
	fmt.Println("\nDone!")
}
