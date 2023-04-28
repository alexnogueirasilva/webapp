package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Err represents an error response
type Err struct {
	Err string `json:"error"`
}

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// HandleErrorResponses returns a JSON response to the request
func HandleErrorResponses(w http.ResponseWriter, r *http.Response) {
	var err Err
	fail := json.NewDecoder(r.Body).Decode(&err)
	if fail != nil {
		return
	}
	JSON(w, r.StatusCode, err)
}
