package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"ps-user/src/adapter/api/domain/db/postgres"
	"ps-user/src/adapter/api/domain/models"
	"ps-user/src/adapter/api/domain/responses"
	"ps-user/src/adapter/api/domain/security"
	"ps-user/src/infrastructure/repositories"

	"strings"

	"github.com/lib/pq"

	"strconv"

	"github.com/gorilla/mux"
)

// CreatUser inserts a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err, "")
		return
	}

	if err = user.Prepare("cadastro"); err != nil {
		responses.Err(w, http.StatusBadRequest, err, "")
		return
	}

	db, err := postgres.Conection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				responses.Err(w, http.StatusUnprocessableEntity, errors.New(fmt.Sprintf("Usuario %s já existe. Entre com um usuário e um e-mail diferente.", user.Nick)), "nick")
				return
			}
		}

		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	responses.Sucess(w, http.StatusCreated, fmt.Sprintf("Usuário %s - id: %d criado com sucesso!", user.Nick, user.ID))

}

// GetUser get user in database
func GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("Id user deve ser inteiro"), param["userID"])
		return
	}
	db, erro := postgres.Conection()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	user, erro := repository.FindById(userID)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	if user.ID == 0 {
		responses.Err(w, http.StatusNotFound, errors.New("Registro não encontrado"), string(userID))
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {

	db, erro := postgres.Conection()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	users, err := repository.FindAllUser()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

// ValidCredentials validation credentials user for oauth
func ValidCredentials(w http.ResponseWriter, r *http.Request) {
	user := strings.ToLower(r.URL.Query().Get("userName"))
	password := strings.ToLower(r.URL.Query().Get("password"))

	db, err := postgres.Conection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	userValid, err := repository.FindByUserAndPassword(user)

	if err = security.VerifyPassword(userValid.PassWord, password); err != nil {
		responses.JSON(w, http.StatusOK, false)
		return
	}
	if userValid.ID != 0 {
		responses.JSON(w, http.StatusOK, true)
		return
	}

}

// DeleteUser delete user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("Id user deve ser inteiro"), param["userID"])
		return
	}
	db, err := postgres.Conection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	user, err := repository.FindById(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	if user.ID == 0 {
		responses.Err(w, http.StatusNotFound, errors.New("Registro não encontrado"), string(userID))
		return
	}
	if err = repository.DeleteById(userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}

//UpdateUser update data user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("Id user deve ser inteiro"), param["userID"])
		return
	}
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Err(w, http.StatusUnprocessableEntity, erro, "")
		return
	}
	var user models.User
	if err = json.Unmarshal(corpoRequisicao, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, erro, "")
		return
	}

	if err = user.Prepare("edicao"); err != nil {
		responses.Err(w, http.StatusBadRequest, err, "")
		return
	}

	db, err := postgres.Conection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	if err := repository.Update(userID, user); err != nil {
		responses.Err(w, http.StatusInternalServerError, erro, "")
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser delete users in database
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	IDs := string(r.URL.Query().Get("ids"))

	if IDs == "" {
		responses.Err(w, http.StatusBadRequest, errors.New("Ids devem ser informados"), "")
		return
	}

	db, err := postgres.Conection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	err = repository.DeleteListId(IDs)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
