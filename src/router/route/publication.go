package route

import (
	"net/http"
	"webapp/src/controllers"
)

var routePublications = []Route{
	{
		URI:         "/publications",
		Method:      http.MethodPost,
		Function:    controllers.CreatePublication,
		RequestAuth: true,
	},
}
