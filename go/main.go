package main

import (
	"fmt"
	"log"
	"net/http"
)

type Index struct {
	Logger *log.Logger
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.Logger.Printf("Received request on '%s'", r.URL.Path)
	fmt.Fprintf(w, HelloWorld())
}

func HelloWorld() string {
	return "Hello world!"
}

func main() {
	port := ":8080"
	logger := log.Default()

	index := &Index{Logger: logger}
	http.Handle("/", index)

	logger.Printf("Listening on port %s", port)
	http.ListenAndServe(port, nil)
}
