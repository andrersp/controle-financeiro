package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/andrersp/controle-financeiro/src/crud"
	"github.com/andrersp/controle-financeiro/src/database"
	"github.com/andrersp/controle-financeiro/src/models"
	"github.com/andrersp/controle-financeiro/src/response"
	"github.com/gorilla/mux"
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
		statusCode, err := NormalizerErrorDB(err)
		response.Erro(w, statusCode, err)
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

func SelectUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	crud := crud.NewCrudUser(db)

	result, err := crud.SelectUser(uint(userID))

	if err != nil {

		statusCode, err := NormalizerErrorDB(err)

		response.Erro(w, statusCode, err)
		return
	}

	response.Sucess(w, http.StatusOK, result)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

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

	if err := crud.UpdateUser(uint(userID), user); err != nil {
		statusCode, err := NormalizerErrorDB(err)
		response.Erro(w, statusCode, err)
		return
	}

	response.Sucess(w, http.StatusOK, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	crud := crud.NewCrudUser(db)

	if err := crud.DeleteUser(uint(userID)); err != nil {
		statusCode, err := NormalizerErrorDB(err)

		response.Erro(w, statusCode, err)
		return
	}

	response.Sucess(w, http.StatusOK, nil)

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 32)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	var paswdModel models.UserPassword

	if err := json.NewDecoder(r.Body).Decode(&paswdModel); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return

	}

	db, err := database.Connect()

	crud := crud.NewCrudUser(db)

	passwdSaved, err := crud.GetUserPassword(uint(userID))

	if err != nil {

		statusCode, err := NormalizerErrorDB(err)
		response.Erro(w, statusCode, err)
		return
	}

	if err := core.CheckPassword(passwdSaved, paswdModel.Old); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := core.HashGenerator(paswdModel.New)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := crud.UpdateUserPassword(uint(userID), string(hashedPassword)); err != nil {

		statusCode, err := NormalizerErrorDB(err)
		response.Erro(w, statusCode, err)
		return
	}

	response.Sucess(w, http.StatusOK, nil)

}
