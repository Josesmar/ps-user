package repositories

import (
	"database/sql"
	"ps-user/models"
	"strconv"
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

	var lastInsertId = 0
	err := repository.db.QueryRow("insert into users (name, nick, email, password) values ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Nick, user.Email, user.PassWord).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil

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

func (repository Users) FindAllUser(page string) ([]models.User, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	limit := 5
	offset := limit*pageInt - 1
	SQL := `select id, name, nick, email, createIn from users ORDER BY id DESC OFFSET $1 LIMIT $2;`

	rows, err := repository.db.Query(SQL, offset, limit)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateIn,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
