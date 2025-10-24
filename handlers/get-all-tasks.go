package handlers

import (
	"encoding/json"
	"net/http"
	"own/models"
)

func (h *Handler) getAllTasks(filter models.Filter) ([]models.Task, models.Metadata, error) {

	q := `SELECT COUNT(*) OVER(), id, title, description, completed, created_at, updated_at 
		  FROM tasks ORDER BY id ASC
		  LIMIT $1 OFFSET $2`

	rows, err := h.Models.Task.DB.Query(q, filter.Limit(), filter.Offset())
	if err != nil {	
		return nil, models.Metadata{}, err
	}
	defer rows.Close()

	var tasks []models.Task 
	var totalRec int
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&totalRec,
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, models.Metadata{}, err
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, models.Metadata{}, err
	}
	return tasks, models.Metadata{}, nil
}

func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	
	f := models.Filter{
		Page: 1,
		PageSize: 5,
	}

	tasks, _, err := h.getAllTasks(f)
	if err != nil {	
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}