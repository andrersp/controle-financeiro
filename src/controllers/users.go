package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/andrersp/controle-financeiro/src/crud"
	"github.com/andrersp/controle-financeiro/src/database"
	"github.com/andrersp/controle-financeiro/src/models"
	"github.com/andrersp/controle-financeiro/src/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	db, err := database.Connect()

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	crud := crud.NewCrudUser(db)

	result, err := crud.Create(user)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Sucess(w, http.StatusCreated, result)

}

func SearchUsers(w http.ResponseWriter, r *http.Request) {

	nameorEmail := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	crud := crud.NewCrudUser(db)
	users, err := crud.SearchUsers(nameorEmail, r)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Sucess(w, http.StatusOK, users)
}
