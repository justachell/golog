package main

import (
	"fmt"
	"io"
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

	requestDelay = 2 // Delay between web-site requests in seconds.
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
			message := "Server is not responding"
			logStatus(os.Stderr, LogErr, message)
			continue
		}

		// Prints web-site status information.
		message := "status " + resp.Status
		logStatus(os.Stdout, LogInfo, message) // Logging status.
		resp.Body.Close()
	}
}

func logStatus(stream io.Writer, l Log, msg string) {
	curTime := time.Now().Format("15:04:05")
	fmt.Fprintf(stream, "[%s]\t%s\t%s\n", l, curTime, msg)
}
