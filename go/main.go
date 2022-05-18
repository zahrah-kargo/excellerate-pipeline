package main

import (
	"log"
	"net/http"

	"github.com/kargotech/excellerate-pipeline/go/pkg/handler"
)

func main() {
	port := ":8080"
	logger := log.Default()

	index := &handler.Index{Logger: logger}
	http.Handle("/", index)

	logger.Printf("Listening on port %s", port)
	http.ListenAndServe(port, nil)
}
