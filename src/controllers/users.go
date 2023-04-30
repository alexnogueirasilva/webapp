package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/config"
	"webapp/src/response"
)

// RegisterUser registers a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nickname"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(user),
	)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Err: err.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err)
		}
	}(resp.Body)

	if resp.StatusCode >= 400 {
		response.HandleErrorResponses(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}
