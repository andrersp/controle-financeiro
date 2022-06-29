package main_test

import (
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

func clearTable(tablename string) {
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	db.Exec(fmt.Sprintf("DELETE FROM %s", tablename))
	db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1", tablename))
}

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
