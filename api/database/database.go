package database

import (
	"database/sql"
	"ps-user/configuration"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar open conection in database and return data
func Conection() (*sql.DB, error) {
	db, err := sql.Open("mysql", configuration.StringConectionBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
