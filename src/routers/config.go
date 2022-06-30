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
	userRouters = append(userRouters, loginRouter)

	for _, router := range userRouters {

		if router.RequireAuth {
			r.HandleFunc(router.URI, core.Logger(core.Authenticate(router.Func))).Methods(router.Method)

		} else {
			r.HandleFunc(router.URI, core.Logger(router.Func)).Methods(router.Method)
		}

	}

	return r
}

func Load() *mux.Router {
	r := mux.NewRouter()

	return configureRouters(r)
}
