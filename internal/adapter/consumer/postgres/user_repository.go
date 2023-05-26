package postgres

import (
	"database/sql"
	"ps-user/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

/*
Create inserts a user into the database
the return is uint64 because after being created an id must be returned
*/
func (repository UserRepository) Create(user domain.User) (uint64, error) {

	var lastInsertId = 0
	err := repository.db.QueryRow("insert into users (name, nick, email, password) values ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Nick, user.Email, user.PassWord).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil

}

// FindById return user in database
func (repository UserRepository) FindById(userID uint64) (domain.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, createIn from users where id =$1", userID,
	)
	if err != nil {
		return domain.User{}, err
	}
	defer lines.Close()

	var user domain.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateIn,
		); err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (repository UserRepository) FindAllUser() ([]domain.User, error) {
	SQL := `select id, name, nick, email, createIn from users ORDER BY id;`
	rows, err := repository.db.Query(SQL)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User

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

func (repository UserRepository) ObterNumeroProdutos() int {

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

func (repository UserRepository) FindByUserAndPassword(userName string) (domain.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, password from users where nick =$1", userName,
	)
	if err != nil {
		return domain.User{}, err
	}
	defer lines.Close()

	var user domain.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.PassWord,
		); err != nil {
			return domain.User{}, err
		}

	}

	return user, nil
}

func (repository UserRepository) DeleteById(userID uint64) error {
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

func (repository UserRepository) Update(user domain.User) error {
	statement, err := repository.db.Prepare(
		"update users set name =$1 where id=$2",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(user.Name, user.ID); err != nil {
		return err
	}
	return nil
}

func (repository UserRepository) DeleteListId(IDs string) error {
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

func (repository UserRepository) FindListUsers(IDs string) ([]domain.User, error) {
	SQL := "select id, name, nick, email, createIn from users where id in(" + IDs + ")"
	rows, err := repository.db.Query(SQL)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User

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
