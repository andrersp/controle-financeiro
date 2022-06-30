package routers

import (
	"net/http"

	"github.com/andrersp/controle-financeiro/src/controllers"
)

var loginRouter = Router{
	URI:         "/login",
	Method:      http.MethodPost,
	Func:        controllers.Login,
	RequireAuth: false,
}
