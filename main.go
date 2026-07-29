package main

import (
	"fmt"
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
		time.Sleep(requestDelay * time.Second)

		// Receiving web-site response.
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf("[ERR] %v\n", err)
			continue
		}

		// Prints web-site status information.
		fmt.Printf("[INFO] status: %d %s\n", resp.StatusCode, resp.Status)
		resp.Body.Close()
	}
}
