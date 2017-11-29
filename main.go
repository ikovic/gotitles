package main

import (
	"fmt"
	"log"
	"os"
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
	fileInfo, err := file.Stat()

	kju := new(queue.FileQueue)

	fmt.Println(kju.IsEmpty())

	kju.Enqueue(fileInfo)

	fmt.Println(kju.IsEmpty())

	infou := *kju.Dequeue()

	fmt.Println(infou.Name())
}
