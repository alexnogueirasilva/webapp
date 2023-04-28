package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookeis"
)

// Logger logs the requests
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookeis.Read(r); err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		next(w, r)
	}
}
