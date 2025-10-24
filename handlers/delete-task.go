package handlers

import (
	"net/http"
	"strconv"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	
	q := `DELETE FROM tasks WHERE id = $1`

	_, err = h.Models.Task.DB.Exec(q, id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}