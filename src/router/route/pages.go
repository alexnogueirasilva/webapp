package route

import (
	"net/http"
	"webapp/src/controllers"
)

var routesUsers = []Route{
	{
		URI:         "/create-user",
		Method:      http.MethodGet,
		Function:    controllers.LoadUsersPage,
		RequestAuth: false,
	},
}
