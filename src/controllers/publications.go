package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/config"
	"webapp/src/request"
	"webapp/src/response"
)

// CreatePublication creates a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	fail := r.ParseForm()
	if fail != nil {
		return
	}

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.Err{Err: err.Error()})
	}

	url := fmt.Sprintf("%s/publications", config.APIURL)
	res, err := request.MakeRequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(publication))
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

	response.JSON(w, res.StatusCode, nil)

}
