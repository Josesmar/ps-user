package postgres

import (
	"database/sql"
	"ps-user/src/adapter/api/domain/configuration"

	_ "github.com/lib/pq"
)

// Conectar open conection in database and return data
func Conection() (*sql.DB, error) {

	db, err := sql.Open("postgres", configuration.StringConectionBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}