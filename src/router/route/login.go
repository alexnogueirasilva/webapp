package route

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		URI:         "/",
		Method:      http.MethodGet,
		Function:    controllers.LoadLoginPage,
		RequestAuth: false,
	},
	{
		URI:         "/login",
		Method:      http.MethodGet,
		Function:    controllers.LoadLoginPage,
		RequestAuth: false,
	},
	{
		URI:         "/login",
		Method:      http.MethodPost,
		Function:    controllers.Login,
		RequestAuth: false,
	},
}
