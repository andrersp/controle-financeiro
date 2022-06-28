package core

import (
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)

	}
}
