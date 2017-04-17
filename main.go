package main

import (
	"log"
	"net/http"
)

func main() {
	var handler http.HandlerFunc
	handler = addIntHandler
	log.Fatal(http.ListenAndServe(":8080", handler))
}
