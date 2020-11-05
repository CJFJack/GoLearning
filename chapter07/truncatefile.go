package main

import (
	"log"
	"os"
)

func main() {
	logfile, err := os.OpenFile("password.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}

	defer logfile.Close()

	log.SetOutput(logfile)
	log.Println("test1")
}
