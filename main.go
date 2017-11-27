package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ikovic/gotitles/hash"
)

func main() {
	args := os.Args[1:]

	path := strings.Join(args, " ")

	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	hash, err := hash.HashFile(file)

	fmt.Println(hash)
}
