package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type (
	Log string
)

const (
	// Types of log output.
	LogInfo Log = "INFO"
	LogWarn Log = "WARN"
	LogErr  Log = "ERR"

	requestDelay = 2 // Delay between requests in seconds.
	timeoutDelay = 3 // Delay of web-site request timeout in seconds.
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("golog: No URL given\n")
	}

	url := args[1] // URL of web-site to check.
	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	client := &http.Client{
		Timeout: timeoutDelay * time.Second,
	}

	for {
		time.Sleep(requestDelay * time.Second)

		// Receiving web-site response.
		resp, err := client.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERR] %v\n", err)
			continue
		}

		// Prints web-site status information.
		fmt.Printf("[INFO] status: %s\n", resp.StatusCode, resp.Status)
		resp.Body.Close()
	}
}
