package repositories

import (
	"database/sql"
	"ps-user/models"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

// NewRepositoryUsers create a user repository
func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

/*
Create inserts a user into the database
the return is uint64 because after being created an id must be returned
*/
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ? ,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(user.Name, user.Nick, user.Email, user.PassWord)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}

// FindById return user in database
func (repository Users) FindById(userID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, createIn from users where id =$1", userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateIn,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
