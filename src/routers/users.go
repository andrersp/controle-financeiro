package routers

import (
	"net/http"

	"github.com/andrersp/controle-financeiro/src/controllers"
)

var userRouters = []Router{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Func:        controllers.SearchUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Func:        controllers.CreateUser,
		RequireAuth: false,
	},
}
