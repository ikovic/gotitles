package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	client := http.Client{Timeout: time.Second * 10}

	fmt.Println(strings.Join(args, " "))
}
