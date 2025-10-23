package handlers

import (
	"encoding/json"
	"net/http"
	"own/models"
)

type ReqCreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}


func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
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
			completed
		) VALUES (
		 	$1, 
			$2, 
			false
		) RETURNING id, title, description, completed, created_at, updated_at`

	err = h.Models.Task.DB.QueryRow(
			q, 
			req.Title, 
			req.Description, 
			).Scan(
				&task.ID,
				&task.Title,
				&task.Description,
				&task.Completed,
				&task.CreatedAt,
				&task.UpdatedAt,
    )

	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}