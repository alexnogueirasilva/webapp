package router

import "github.com/gorilla/mux"

// Generate returns a new mux.Router
func Generate() *mux.Router {
	return mux.NewRouter()
}
