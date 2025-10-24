package handlers

import (
	"encoding/json"
	"net/http"
	"own/models"
	"strconv"
)

func (h *Handler) getTask (id int) (models.Task, error) {
	var task models.Task

	q := `SELECT id, title, description, completed, created_at, updated_at 
	FROM tasks 
	WHERE id = $1`

	err := h.Models.Task.DB.QueryRow(q, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.getTask(id)

	if err != nil {
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}