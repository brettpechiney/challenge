package challenge

import (
	"bytes"
	"net/http"
	"testing"
)

const (
	// ValidModel is a serialized valid model implementing the Okay interface.
	ValidModel = `{"firstName":"first","lastName":"last","username":"user","password":"pass","role":"customer"}`

	// InvalidModel is a serialized invalid model implementing the Okay interface.
	InvalidModel = `{"firstName":"first","lastName":"last","username":"user","role":"customer"}`
)

// MakeRequest creates an *http.Request with an optional body.
func MakeRequest(t *testing.T, method string, url string, body string) *http.Request {
	t.Helper()
	var jsonBytes []byte
	if len(body) > 0 {

	}
	jsonBytes = []byte(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Errorf("could not create request: %v", err)
	}
	return req
}
