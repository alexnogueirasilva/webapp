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

// Read reads the cookie from the user's browser
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	var data map[string]string
	if err = s.Decode("data", cookie.Value, &data); err != nil {
		return nil, err
	}

	return data, nil
}
