package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func (s *TaskStore) createTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	var input CreateTaskInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(input.Title) == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	task := Task{
		ID:          s.NextID,
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}

	s.tasks = append(s.tasks, task)
	s.NextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}

func (s *TaskStore) getTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	for _, task := range s.tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "task not found", http.StatusNotFound)
}

func (s *TaskStore) updateTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var input UpdateTaskInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if input.Title != nil && strings.TrimSpace(*input.Title) == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	for i := range s.tasks {
		task := &s.tasks[i]
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			if input.Title != nil {
				task.Title = *input.Title
			}

			if input.Description != nil {
				task.Description = *input.Description
			}

			if input.Completed != nil {
				task.Completed = *input.Completed
			}
			json.NewEncoder(w).Encode(task)
			return

		}
	}
	http.Error(w, "task not found", http.StatusNotFound)

}
