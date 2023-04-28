package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/request"
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
	response, err := request.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)

	fmt.Println(response.StatusCode, err)
	utils.ExecuteTemplate(w, "home.html", nil)
}
