package router

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"net/http"
	"os"
)

func sayhello(w http.ResponseWriter, rec *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func Router() {
	http.HandleFunc("/", sayhello)
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("Logging to a file in Go!")

}