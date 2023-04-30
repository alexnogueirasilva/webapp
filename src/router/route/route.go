package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"webapp/src/middlewares"
)

// Route represents all routes from the WebApp
type Route struct {
	URI         string
	Method      string
	Function    func(w http.ResponseWriter, r *http.Request)
	RequestAuth bool
}

// Setup configures all routes from the WebApp
func Setup(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, routesUsers...)
	routes = append(routes, routeHome)
	routes = append(routes, routePublications...)

	for _, route := range routes {

		if route.RequestAuth {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
			continue
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))

	return router

}
