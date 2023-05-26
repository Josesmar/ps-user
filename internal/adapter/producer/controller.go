package producer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"ps-user/internal/adapter/security"
	"ps-user/internal/application"
	"ps-user/internal/domain"
	"ps-user/internal/infra/rest/response"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

type Controller struct {
	userService application.UserService
}

func NewController(userService application.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

// CreatUser inserts a user into the database
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	var user domain.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err, "")
		return
	}

	if err = user.Prepare("cadastro"); err != nil {
		response.Err(w, http.StatusBadRequest, err, "")
		return
	}
	userID, err := c.userService.Create(user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				message := fmt.Sprintf("usuario %s já existe. Entre com um usuário e um e-mail diferente.", user.Nick)
				response.Err(w, http.StatusUnprocessableEntity, errors.New(message), "nick")
				return
			}
		}

		response.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	response.Sucess(w, http.StatusCreated, fmt.Sprintf("usuário id: %d criado com sucesso!", userID))

}

// GetUser get user in database
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, errors.New("id user deve ser inteiro"), param["userID"])
		return
	}
	user, erro := c.userService.FindById(userID)
	if erro != nil {
		response.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	if user.ID == 0 {
		response.Err(w, http.StatusNotFound, errors.New("Registro não encontrado"), fmt.Sprint(userID))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func (c *Controller) GetAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := c.userService.FindAllUser()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	response.JSON(w, http.StatusOK, users)

}

// ValidCredentials validation credentials user for oauth
func (c *Controller) ValidCredentials(w http.ResponseWriter, r *http.Request) {
	user := strings.ToLower(r.URL.Query().Get("userName"))
	password := strings.ToLower(r.URL.Query().Get("password"))

	userValid, err := c.userService.FindByUserAndPassword(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err, "error ao buscar usuário")
		return
	}

	if err = security.VerifyPassword(userValid.PassWord, password); err != nil {
		response.JSON(w, http.StatusOK, false)
		return
	}
	if userValid.ID != 0 {
		response.JSON(w, http.StatusOK, true)
		return
	}

}

// DeleteUser delete user in database
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, errors.New("id user deve ser inteiro"), param["userID"])
		return
	}
	user, err := c.userService.FindById(userID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err, "")
		return
	}

	if user.ID == 0 {
		response.Err(w, http.StatusNotFound, errors.New("registro não encontrado"), fmt.Sprint(userID))
		return
	}
	if err = c.userService.DeleteById(userID); err != nil {
		response.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

// UpdateUser update data user
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, errors.New("id user deve ser inteiro"), param["userID"])
		return
	}
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Err(w, http.StatusUnprocessableEntity, erro, "")
		return
	}
	var user domain.User
	if err = json.Unmarshal(corpoRequisicao, &user); err != nil {
		response.Err(w, http.StatusBadRequest, erro, "")
		return
	}

	if err = user.Prepare("edicao"); err != nil {
		response.Err(w, http.StatusBadRequest, err, "")
		return
	}

	if err := c.userService.Update(user); err != nil {
		response.Err(w, http.StatusInternalServerError, erro, "")
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser delete users in database
func (c *Controller) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	IDs := string(r.URL.Query().Get("ids"))

	if IDs == "" {
		response.Err(w, http.StatusBadRequest, errors.New("ids devem ser informados"), "")
		return
	}

	user, erro := c.userService.FindListUsers(IDs)
	if erro != nil {
		response.Err(w, http.StatusInternalServerError, errors.New("error in get all users"), "")
		return
	}

	if user == nil {
		response.Err(w, http.StatusNotFound, errors.New("registro não encontrado para o(s) id(s) informado(s)"), IDs)
		return
	}

	err := c.userService.DeleteListId(IDs)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err, "")
		return
	}
	response.JSON(w, http.StatusOK, nil)
}
