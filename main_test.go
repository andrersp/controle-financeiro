package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/andrersp/controle-financeiro/src/database"
	"github.com/andrersp/controle-financeiro/src/routers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var r *mux.Router

func TestMain(m *testing.M) {

	core.LoadConfig()

	if err := database.SetupAPP(); err != nil {
		log.Fatal(err)

	}

	r = routers.Load()

	fmt.Printf("Open on port %d \n", core.API_PORT)
	exitVal := m.Run()

	os.Exit(exitVal)

}

// func TestGetUsers(t *testing.T) {

// 	req, _ := http.NewRequest("GET", "/users", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	body := getResponseBody(response)

// 	assert.Equal(t, true, body["success"], "Error response sucess")

// }

func TestCreateuser(t *testing.T) {

	var strUser = []byte(`{"name": "Andre Luis", "email": "rsp.assistencia@gmail.com"`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(strUser))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

}

// func TestGetUser(t *testing.T) {

// 	req, _ := http.NewRequest("GET", "/users/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// }

// func TestUpdateUser(t *testing.T) {

// 	req, _ := http.NewRequest("PUT", "/users/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// }

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual, "Test Error")

}

func getResponseBody(req *httptest.ResponseRecorder) map[string]interface{} {

	response := make(map[string]interface{})
	json.Unmarshal(req.Body.Bytes(), &response)
	return response

}
