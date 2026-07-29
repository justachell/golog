package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("golog: No URL given\n")
	}

	url := args[1] // URL of web-site to check.

}
