package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	LOG_FILE := "C:\\Users\\frank\\Documents\\logs\\log.txt"

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("logging to custom file")

}
