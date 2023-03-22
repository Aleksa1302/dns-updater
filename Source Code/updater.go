package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func sendPush(url string) {
	fmt.Println("Sending push request...")
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error sending push:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Push sent successfully")
}

func main() {
	args := os.Args[1:]

	// check if enough arguments are provided
	if len(args) < 2 {
		fmt.Println("Usage: app.exe <URL> <interval in seconds>")
		fmt.Println("Example: app.exe http://url:3001/api/push/m1lE6JIXQm?status=up&msg=OK 300")
		os.Exit(1)
	}

	url := args[0]

	// parse interval from command line argument
	interval, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid interval provided:", err)
		os.Exit(1)
	}

	sendPush(url)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		sendPush(url)
	}
}
