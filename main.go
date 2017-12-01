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
var languages []string
var wg sync.WaitGroup

// create the special directory for saving subtitles
func createAppDirectory(path string) {
	appDirectory := filepath.Join(path, AppName)
	if _, err := os.Stat(appDirectory); os.IsNotExist(err) {
		os.Mkdir(appDirectory, os.ModePerm)
	}
}

// Search and download subtitles for the given file
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

// Check if the file is a video, download subtitles if it is
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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "language, l",
			Value: "eng",
			Usage: "3 letter code",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 1 {
			fmt.Println("Path to a folder should be provided")
			return nil
		}
		path := c.Args().Get(0)
		languages = append(languages, c.String("language"))

		// log in to OSDB first
		osdbClient.LogIn("", "", "")

		// walk the FS and search for video files
		filepath.Walk(path, walk)
		return nil
	}

	app.Run(os.Args)

	wg.Wait()
	fmt.Println("\nDone!")
}
