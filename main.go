package main

import (
	"fmt"
	"log"
	"net/http"
	"own/database"
	"own/models"
	"own/handlers"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Learning by doing"))
}


func main() {

	dsn := "user=postgres dbname=TaskAPI password=101010 sslmode=disable"

	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// 1. Initialize models
	h := &handlers.Handler{
		Models: models.NewModel(db),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	//mux.HandleFunc("GET /tasks", h.getAllTask)
	//mux.HandleFunc("GET /tasks/{id}", getTask)
	//mux.HandleFunc("PUT /tasks/{id}", updateTask)
	//mux.HandleFunc("DELETE /tasks/{id}", deleteTask)

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":4000", mux))

}