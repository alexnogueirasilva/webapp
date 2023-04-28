package cookeis

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

var s *securecookie.SecureCookie

// Setup sets up the secure cookie
func Setup() {
	s = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

// Save saves the cookie in the user's browser
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
