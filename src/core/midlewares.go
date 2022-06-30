package core

import (
	"log"
	"net/http"

	"github.com/andrersp/controle-financeiro/src/response"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)

	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if erro := ValidateToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
