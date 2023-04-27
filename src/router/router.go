package router

import (
	"github.com/gorilla/mux"
	"webapp/src/router/route"
)

// Generate returns a new mux.Router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return route.Setup(r)
}
