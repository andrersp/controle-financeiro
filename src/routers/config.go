package routers

import (
	"net/http"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/gorilla/mux"
)

type Router struct {
	URI         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func configureRouters(r *mux.Router) *mux.Router {
	userRouters := userRouters

	for _, router := range userRouters {
		r.HandleFunc(router.URI, core.Logger(router.Func)).Methods(router.Method)
	}

	return r
}

func Load() *mux.Router {
	r := mux.NewRouter()

	return configureRouters(r)
}
