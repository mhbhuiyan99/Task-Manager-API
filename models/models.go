package models

import "database/sql"

type Models struct {
	Task TaskModel
}

type Handler struct {
	Models Models
}

func NewModel(db *sql.DB) Models {
	return Models{
		Task: TaskModel{DB: db},
	}
}