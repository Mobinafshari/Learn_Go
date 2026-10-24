package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	log.Println("Server running on :8090")
	log.Fatal(http.ListenAndServe(":8090", mux))
}
