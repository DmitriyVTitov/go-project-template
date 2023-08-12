package main

import (
	"log"
	"net/http"

	"github.com/DmitriyVTitov/go-project-template/internal/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
