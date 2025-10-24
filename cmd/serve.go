package cmd

import (
	"fmt"
	"net/http"
	"own/handlers"
)

type Server struct {
	Handlers *handlers.Handler
}

func (s *Server) Serve() error {
	srv := http.Server{
		Handler: s.routers(),
		Addr:    ":4000",
	}

	fmt.Println("Server is running...")
	return srv.ListenAndServe()
}

func (s *Server) routers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /tasks", s.Handlers.CreateTask)
	mux.HandleFunc("GET /tasks", s.Handlers.GetAllTasks)
	mux.HandleFunc("GET /tasks/{id}", s.Handlers.GetTask)
	mux.HandleFunc("PUT /tasks/{id}", s.Handlers.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", s.Handlers.DeleteTask)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Learning by doing"))
}