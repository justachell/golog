package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const (
	requestDelay = 2 // Delay between requests in seconds.
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("golog: No URL given\n")
	}

	url := args[1] // URL of web-site to check.

	for {
		resp, err := http.Get(url)
		if err != nil {
			// TODO: Error handing...
		}
		// TODO: Info output...
		time.Sleep(requestDelay * time.Second)
	}
}
