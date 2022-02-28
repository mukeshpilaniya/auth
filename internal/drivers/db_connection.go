package drivers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DBConnection(dsn string) (*sql.DB, error){
	db , err := sql.Open("postgres", dsn)

	if err !=nil {
		return nil, err
	}
	err = db.Ping()
	if err !=nil {
		return nil, err
	}
	return db, nil
}
