package handlers

import (
	"own/models"
)

type ReqCreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type application struct {
	Models models.Models
}
/*
func CreateTask(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateTask
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	var task models.Task

	q := `INSERT INTO tasks (
			title, 
			description, 
		) VALUES (
		 	$1, 
			$2, 
			false, 
			NOW(), 
			NOW()
		) RETURNING id, title, description, completed, created_at, updated_at`

	err = DB.QueryRow(
			q, 
			req.Title, 
			req.Description, 
			).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
	*/