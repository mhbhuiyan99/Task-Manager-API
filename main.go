package main

import (
	"log"
	"own/database"
	"own/handlers"
	"own/models"
	"own/cmd"
)

func main() {

	dsn := "user=postgres dbname=TaskAPI password=101010 sslmode=disable"

	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Initialize models
	h := &handlers.Handler{
		Models: models.NewModel(db),
	}

	server := &cmd.Server{
		Handlers: h,
	}

	if err := server.Serve(); err != nil {
		log.Fatal("Cannot start server:", err)
	}

}