package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type dataResult struct {
	Success bool   `json:"success"`
	Err     string `json:"error,omitempty"`
}

// JSON return response for req	uest
func responder(w http.ResponseWriter, statusCode int, dados interface{}) {

	w.WriteHeader(statusCode)

	if dados != nil {

		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}
func Sucess(w http.ResponseWriter, statusCode int, dados interface{}) {

	if dados == nil && statusCode == 204 {
		responder(w, statusCode, nil)
		return
	}

	responder(w, statusCode, dados)

}

func Erro(w http.ResponseWriter, statusCode int, erro error) {

	var result dataResult
	result.Success = false
	result.Err = erro.Error()

	responder(w, statusCode, result)

}
