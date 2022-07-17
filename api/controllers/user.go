package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"ps-user/database"
	"ps-user/infrastructure/repositories"
	"ps-user/models"
	"ps-user/responses"

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

	db, err := database.Conection()
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
				responses.Err(w, http.StatusInternalServerError, errors.New(fmt.Sprintf("Usuario %s já existe. Entre com um usuário e um e-mail diferente.", user.Nick)), "nick")
				return
			}
		}

		responses.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	responses.JSON(w, http.StatusCreated, fmt.Sprintf("Usuário %s - id: %d criado com sucesso!", user.Nick, user.ID))

}

//GetUser get user in database
func GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("Id user deve ser inteiro"), param["userID"])
		return
	}
	db, erro := database.Conection()
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
