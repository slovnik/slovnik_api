package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(router)))
}
