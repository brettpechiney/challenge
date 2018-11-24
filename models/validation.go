package models

import (
	"encoding/json"
	"net/http"
)

// okay is an interface implemented by types that can validate
// themselves.
type okay interface {
	OK() error
}

// DecodeRequest attempts to parse and validate an object implementing
// the okay interface contained in an HTTP request.
func DecodeRequest(r *http.Request, v okay) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}
