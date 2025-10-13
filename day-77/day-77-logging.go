package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Simple log message
	log.Println("This is a log message")

	// Log with Fatal (program exit after logging)
	// log.Fatal("Something went wrong!")

	// Log with Panic (logs + panic)
	// log.Panic("Unexpected panic happened!")

	// Custom logger writing to file
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	customLogger := log.New(file, "TigerXApp: ", log.Ldate|log.Ltime|log.Lshortfile)
	customLogger.Println("This is a custom log entry")
	fmt.Println("Log written to app.log ✅")
}
