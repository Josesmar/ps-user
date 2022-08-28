package repositories

import (
	"database/sql"
	"ps-user/src/adapter/api/domain/models"
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

func (repository Users) FindAllUser() ([]models.User, error) {
	SQL := `select id, name, nick, email, createIn from users ORDER BY id;`
	rows, err := repository.db.Query(SQL)

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

func (repository Users) ObterNumeroProdutos() int {

	var sql string = `SELECT count(0) FROM users`

	rows, err := repository.db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var totalUsers int = 0
	for rows.Next() {
		rows.Scan(&totalUsers)
	}

	return totalUsers
}

func (repository Users) FindByUserAndPassword(userName string) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, password from users where nick =$1", userName,
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
			&user.PassWord,
		); err != nil {
			return models.User{}, err
		}

	}

	return user, nil
}

func (repository Users) DeleteById(userID uint64) error {
	statement, erro := repository.db.Prepare("delete from users where id =$1")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}
	return nil
}

func (repository Users) Update(userID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name =$1 where id=$2",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(user.Name, userID); err != nil {
		return err
	}
	return nil
}

func (repository Users) DeleteListId(IDs string) error {
	sql := "delete from users where id in(" + IDs + ")"
	statement, err := repository.db.Prepare(sql)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(); err != nil {
		return err
	}
	return nil
}

func (repository Users) FindListUsers(IDs string) ([]models.User, error) {
	SQL := "select id, name, nick, email, createIn from users where id in(" + IDs + ")"
	rows, err := repository.db.Query(SQL)

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
