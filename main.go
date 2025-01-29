package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &profilehandler{})

	logger := log.Default()

	logger.Println("Serving HTTP server in port :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("An unknown error occurred")
	}

}
