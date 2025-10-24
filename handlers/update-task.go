package handlers

import (
	"encoding/json"
	"net/http"
	"own/models"
	"strconv"
)

type ReqUpdateTask struct {
	Completed   bool   `json:"completed"`
}

func (h Handler) updateTask(id int, completed bool) (models.Task, error) {
	var task models.Task

	q := `UPDATE tasks SET 
			completed = $1,
			updated_at = NOW() AT TIME ZONE 'UTC'
		WHERE id = $2
		RETURNING id, title, description, completed, created_at, updated_at`
	
	err := h.Models.Task.DB.QueryRow(q, completed, id).Scan(
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

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req ReqUpdateTask
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	update, err := h.updateTask(id, req.Completed)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(update)


}