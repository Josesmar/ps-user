package postgres

import (
	"database/sql"
	"ps-user/internal/adapter/config"

	_ "github.com/lib/pq"
)

type Postgres struct{}

// Conectar open conection in database and return data
func (p *Postgres) Conection() (*sql.DB, error) {

	db, err := sql.Open("postgres", config.StringConectionBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
