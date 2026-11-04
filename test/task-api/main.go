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
	mux.HandleFunc("GET /tasks", tasks.listTasks)
	mux.HandleFunc("POST /tasks", tasks.createTask)
	mux.HandleFunc("GET /tasks/{id}", tasks.getTask)
	mux.HandleFunc("POST /tasks/{id}", tasks.updateTask)

	log.Println("Server running on :8090")
	log.Fatal(http.ListenAndServe(":8090", mux))

}
