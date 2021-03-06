package routers

import (
	"net/http"

	"github.com/andrersp/controle-financeiro/src/controllers"
)

var userRouters = []Router{

	{
		URI:         "/users",
		Method:      http.MethodPost,
		Func:        controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Func:        controllers.SearchUsers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodGet,
		Func:        controllers.SelectUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPut,
		Func:        controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodDelete,
		Func:        controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPost,
		Func:        controllers.UpdatePassword,
		RequireAuth: true,
	},
}
