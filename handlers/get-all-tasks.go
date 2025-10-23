package handlers

import (
	"encoding/json"
	"net/http"
	"own/models"
)

func (h Handler) getAllTasks() ([]models.Task, error) {

	q := `SELECT id, title, description, completed, created_at, updated_at FROM tasks`

	rows, err := h.Models.Task.DB.Query(q)
	if err != nil {	
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task 
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil,  err
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.getAllTasks()
	if err != nil {	
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}