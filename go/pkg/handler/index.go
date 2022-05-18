package handler

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
	fmt.Fprintf(w, "Hello world!")
}
