package internal

import (
	"log"
	"os"
)

var openFlags = os.O_APPEND | os.O_CREATE | os.O_WRONLY
var infoFlags = log.Ldate | log.Ltime
var errorFlags = log.Ldate | log.Ltime
var fatalFlags = log.Ldate | log.Ltime

func InfoLog(message string) {
	file, err := os.OpenFile("logs.txt", openFlags, 0666)
	if err != nil {
		log.Fatal("Error loading logs.txt file")
	}
	InfoLogger := log.New(file, "INFO: ", infoFlags)
	InfoLogger.Println(message)
}
func ErrorLog(message string) {
	file, err := os.OpenFile("logs.txt", openFlags, 0666)
	if err != nil {
		log.Fatal("Error loading logs.txt file")
	}
	ErrorLogger := log.New(file, "ERROR: ", errorFlags)
	ErrorLogger.Println(message)
}
func FatalLog(message string) {
	file, err := os.OpenFile("logs.txt", openFlags, 0666)
	if err != nil {
		log.Fatal("Error loading logs.txt file")
	}
	ErrorLogger := log.New(file, "FATAL: ", fatalFlags)
	ErrorLogger.Println(message)
	os.Exit(1)
}
