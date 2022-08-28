package models

import (
	"errors"
	"ps-user/src/adapter/api/domain/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	PassWord string    `json:"password,omitempty"`
	CreateIn time.Time `json:"creatIn,omitempty"`
}

type PageUser struct {
	totalElements int
	totalPages    int
}

// Prepare will call methods to validate and format the received user
func (user *User) Prepare(etapa string) error {
	if erro := user.validate(etapa); erro != nil {
		return erro
	}
	if erro := user.format(etapa); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(etapa string) error {

	switch etapa {
	case "edicao":
		if user.Name == "" {
			return errors.New("O nome é obrigatório e não pode estar em branco")
		}
	case "cadastro":
		if user.Name == "" {
			return errors.New("O nome é obrigatório e não pode estar em branco")
		}
		if user.Nick == "" {
			return errors.New("O nick é obrigatório e não pode estar em branco")
		}
		if user.Email == "" {
			return errors.New("O email é obrigatório e não pde estar em branco")
		}
		if erro := checkmail.ValidateFormat(user.Email); erro != nil {
			return errors.New("O e-mail inserido é inválido")
		}
		if user.PassWord == "" {
			return errors.New("O password é obrigatório e não pode estar em branco")
		}
	default:
		return nil
	}
	return nil
}

func (user *User) format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(user.PassWord)
		if erro != nil {
			return erro
		}

		user.PassWord = string(senhaComHash)
	}

	return nil
}
