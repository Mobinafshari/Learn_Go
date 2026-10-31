package main

type Task struct {
	ID          int  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
