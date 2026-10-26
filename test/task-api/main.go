package main

import (
	"encoding/json"
	"fmt"
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
	tasks := &TaskStore{
		tasks:  make([]Task, 0, 10),
		NextID: 1,
	}
	fmt.Println(tasks)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	log.Println("Server running on :8090")
	log.Fatal(http.ListenAndServe(":8090", mux))

}
