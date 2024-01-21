package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a log file for errors
	errorLog, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open error log file:", err)
	}
	defer errorLog.Close() // defers runs at the end of main()

	// Create a log file for the response
	responseLog, err := os.OpenFile("response.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open response log file:", err)
	}
	defer responseLog.Close()

	// Create a logger for the response
	responseLogger := log.New(responseLog, "", log.LstdFlags)

	// Send GET request to the API
	response, err := http.Get("https://api.cadegray.dev/joke")
	if err != nil {
		log.SetOutput(errorLog)
		log.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.SetOutput(errorLog)
		log.Println("Error:", err)
		return
	}

	// Log the response body
	responseLogger.Println(string(body))
}
