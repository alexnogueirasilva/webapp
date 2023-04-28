package route

import (
	"net/http"
	"webapp/src/controllers"
)

var routeHome = Route{
	URI:         "/home",
	Method:      http.MethodGet,
	Function:    controllers.LoadHomePage,
	RequestAuth: true,
}
