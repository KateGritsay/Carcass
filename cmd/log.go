package cmd

import (
	"os"
	log "github.com/sirupsen/logrus"

)

func Loger() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("Logging to a file in Go!")
}
