package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LFTrip/ProjectLFTrip/api/auth"
	"github.com/LFTrip/ProjectLFTrip/api/models"
	"github.com/LFTrip/ProjectLFTrip/api/responses"
	"github.com/LFTrip/ProjectLFTrip/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

type response struct {
	Token	string 			`json:"token"`
	User 	models.User 	`json:"user"`
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (response, error) {

	var err error
	var res response
	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return res, err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, err
	}

	token, err := auth.CreateToken(user.ID)
	res = response{
        Token:  token,
		User:  user}

	return res, err
}