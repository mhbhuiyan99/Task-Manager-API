package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"own/database"
	"own/models"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Learning by doing"))
}


type ReqCreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type application struct {
	Models models.Models
}

func (app *application) createTask(w http.ResponseWriter, r *http.Request) {
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

	err = app.Models.Task.DB.QueryRow(
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
	json.NewEncoder(w).Encode(task)
}


func main() {

	dsn := "user=postgres dbname=TaskAPI password=101010 sslmode=disable"

	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// 1. Initialize models
	app := &application{
		Models: models.NewModel(db),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /tasks", app.createTask)
	//mux.HandleFunc("GET /tasks", getAllTask)
	//mux.HandleFunc("GET /tasks/{id}", getTask)
	//mux.HandleFunc("PUT /tasks/{id}", updateTask)
	//mux.HandleFunc("DELETE /tasks/{id}", deleteTask)

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":4000", mux))

}