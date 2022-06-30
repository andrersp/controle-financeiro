package main_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	clearTable("users")
	var strUser = []byte(`{"name": "Andre Luis", 
						   "email": "rsp.assistencia@gmail.com", 
						   "password": "andre123",
						   "admin": true, "enable": true}`)

	var strUserUpdate = []byte(`{"name": "Andre Luis França",
						   "email": "rsp.assistencia@gmail.com",
						   "admin": true, "enable": true}`)

	var strPSWD = []byte(`{"old": "andre123", "new": "admin123"}`)

	var strLogin = []byte(`{ "email": "rsp.assistencia@gmail.com", "password": "andre123"}`)

	t.Run("CreateUser", func(t *testing.T) {

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(strUser))

		response := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, response.Code)

		body := getResponseBody(response)

		msg := body

		assert.Equal(t, true, msg["success"], fmt.Sprintf("%s", msg))

	})

	t.Run("CreateError", func(t *testing.T) {

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(strUser))

		response := executeRequest(req)
		checkResponseCode(t, http.StatusBadRequest, response.Code)

		body := getResponseBody(response)

		msg := body

		assert.Equal(t, false, msg["success"], fmt.Sprintf("%s", msg))

	})

	t.Run("LoginSucess", func(t *testing.T) {

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(strLogin))

		response := executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		body := getResponseBody(response)

		msg := body

		Token = fmt.Sprint(msg["token"])

		assert.Equal(t, true, msg["success"], fmt.Sprintf("%s", msg))

	})

	t.Run("GetAllUser", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users", nil)

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)

		body := getResponseBody(response)

		assert.Equal(t, true, body["success"], body["error"])

	})

	t.Run("GetUserSuccess", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)
		body := getResponseBody(response)

		assert.Equal(t, true, body["success"], body["success"])
	})

	t.Run("GetUserNotFound", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/2", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusNotFound, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, false, body["success"], body["success"])
	})

	t.Run("UpdateUser", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(strUserUpdate))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, true, body["success"], body["success"])

	})

	t.Run("UpdateUserNotFound", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/users/2", bytes.NewBuffer(strUserUpdate))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusNotFound, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, false, body["success"], body["success"])

	})

	t.Run("UpdatePasswordSuccess", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/users/1", bytes.NewBuffer(strPSWD))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, true, body["success"], body["error"])
	})

	t.Run("UpdatePasswordErrorPasswd", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/users/1", bytes.NewBuffer(strPSWD))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusBadRequest, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, false, body["success"], body["error"])
	})

	t.Run("UpdatePasswordUserNotFound", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/users/2", bytes.NewBuffer(strPSWD))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusNotFound, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, false, body["success"], body["error"])
	})

	t.Run("LoginError", func(t *testing.T) {

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(strLogin))

		response := executeRequest(req)
		checkResponseCode(t, http.StatusUnauthorized, response.Code)

		body := getResponseBody(response)

		msg := body

		assert.Equal(t, false, msg["success"], fmt.Sprintf("%s", msg))

	})

	t.Run("DeleteUser", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/users/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusOK, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, true, body["success"], body["error"])
	})

	t.Run("DeleteUserErrNotFound", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/users/2", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
		response := executeRequest(req)
		checkResponseCode(t, http.StatusNotFound, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, false, body["success"], body["error"])
	})

	t.Run("DeleteUserErrUnauthorized", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/users/2", nil)

		response := executeRequest(req)
		checkResponseCode(t, http.StatusUnauthorized, response.Code)
		body := getResponseBody(response)
		assert.Equal(t, false, body["success"], body["error"])
	})

}
