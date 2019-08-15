package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, rec *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func Router() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}