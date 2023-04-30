package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookeis"
	"webapp/src/models"
	"webapp/src/request"
	"webapp/src/response"
	"webapp/src/utils"
)

// LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadUsersPage loads the users page
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

// LoadHomePage loads the home page
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	res, err := request.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Err: err.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err)
		}
	}(res.Body)

	if res.StatusCode >= 400 {
		response.HandleErrorResponses(w, res)
		return
	}

	var publications []models.Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Err: err.Error()})
		return
	}

	cookie, _ := cookeis.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}
