package route

import (
	"github.com/gorilla/mux"
	"net/http"
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

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))

	return router

}
