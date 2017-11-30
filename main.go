package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ikovic/gotitles/queue"
)

func main() {
	args := os.Args[1:]

	path := strings.Join(args, " ")
	stuff := []string{"test", "test3"}
	// hash, err := hash.HashFile(file)

	traverse(path, stuff)

	// fmt.Println(hash)
}

func traverse(path string, nodes []string) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	queue := new(queue.FileQueue)
	queue.Enqueue(*file)

	for !queue.IsEmpty() {
		f := queue.Dequeue()
		fileInfo, _ := f.Stat()
		fmt.Println(fileInfo.Name())
		// if dir, get children info
		if fileInfo.IsDir() {
			childrenNames, _ := f.Readdirnames(0)
			for _, name := range childrenNames {
				// find the exact path
				filePath, _ := filepath.Abs(filepath.Dir(name))
				fmt.Println(filePath)
				child, _ := os.Open(filePath)
				queue.Enqueue(*child)
			}
		}

	}

}
