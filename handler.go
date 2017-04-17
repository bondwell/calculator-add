package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bondwell/calculator-add/add"
)

type addIntRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type addIntResponse struct {
	Result int `json:"result"`
}

func addIntHandler(w http.ResponseWriter, r *http.Request) {
	var request addIntRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("error decoding request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sum := add.Integers(request.A, request.B)
	response := addIntResponse{Result: sum}
	log.Printf("%v + %v = %v\n", request.A, request.B, sum)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
