package models

import "database/sql"

type Models struct {
	Task TaskModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Task: TaskModel{DB: db},
	}
}