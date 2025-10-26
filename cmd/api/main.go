package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", server.handleHealth)
	mux.HandleFunc("/suggest", server.handleSuggest)

	fmt.Println("Server running at: http://localhost:9696")
	log.Fatal(http.ListenAndServe(":9696", mux))
}
