package http

import (
	"encoding/json"
	"net/http"

	"github.com/brettpechiney/challenge/challenge"
)

// DecodeRequest attempts to parse and validate an object implementing
// the Okay interface contained in an HTTP request.
func DecodeRequest(r *http.Request, v challenge.Okay) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}

// RespondWithJSON creates a response to an HTTP request by writing JSON
// to a ResponseWriter and setting response headers.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError creates an error response to an HTTP request by
// writing to a ResponseWriter.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
