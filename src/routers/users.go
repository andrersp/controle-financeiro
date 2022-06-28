package routers

import (
	"net/http"
)

var userRouters = []Router{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Func:        func(w http.ResponseWriter, r *http.Request) {},
		RequireAuth: false,
	},
}
