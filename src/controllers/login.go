package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/andrersp/controle-financeiro/src/crud"
	"github.com/andrersp/controle-financeiro/src/database"
	"github.com/andrersp/controle-financeiro/src/models"
	"github.com/andrersp/controle-financeiro/src/response"
)

func Login(w http.ResponseWriter, r *http.Request) {

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

	result, err := crud.SearchByEmail(user.Email)

	if err != nil {
		statusCode, err := NormalizerErrorDB(err)
		response.Erro(w, statusCode, err)
		return
	}

	if err := core.CheckPassword(result.Password, user.Password); err != nil {
		err = errors.New("Email or Passowrd Invalid!")
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := core.CreateToken(result.ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Sucess(w, http.StatusOK, struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Token   string `json:"token"`
		Success bool   `json:"success"`
	}{
		result.ID,
		result.Name,
		string([]byte(token)),
		true,
	})

}
