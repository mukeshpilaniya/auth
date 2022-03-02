package models

import (
	"database/sql"
)

// DBModel is a type for database connection value
type DBModel struct {
	DB *sql.DB
}

// Models is the wrappers for all models
type Models struct {
	DB DBModel
}

// NewModel return a model with database connection pool
func NewModel(db *sql.DB) Models{
	return Models{
		DB: DBModel{DB: db},
	}
}