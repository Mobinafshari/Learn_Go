package main

import (
	"encoding/json"
	"net/http"
)

func (s *TaskStore) listTasks(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(s.tasks); err != nil {
		return
	}
}

