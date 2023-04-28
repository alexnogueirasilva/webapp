package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookeis"
	"webapp/src/models"
	"webapp/src/response"
)

// Login is responsible for authenticating a user in the system
func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	users, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.Err{Err: err.Error()})
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(users),
	)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Err: err.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode >= 400 {
		response.HandleErrorResponses(w, res)
		return
	}

	var authenticationData models.Authentication
	if err = json.NewDecoder(res.Body).Decode(&authenticationData); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Err: err.Error()})
		return
	}

	if err = cookeis.Save(w, authenticationData.ID, authenticationData.Token); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Err: err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, nil)
}
