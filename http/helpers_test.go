package http_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brettpechiney/challenge/challenge"
	myhttp "github.com/brettpechiney/challenge/http"
)

func TestDecodeRequest(t *testing.T) {
	testCases := []struct {
		Name      string
		JSON      string
		ShouldErr bool
	}{
		{
			Name:      "ValidBody",
			JSON:      challenge.ValidModel,
			ShouldErr: false,
		},
		{
			Name:      "InvalidModel",
			JSON:      challenge.InvalidModel,
			ShouldErr: true,
		},
		{
			Name:      "InvalidModelPassed",
			JSON:      "laksjd",
			ShouldErr: true,
		},
	}

	for _, tc := range testCases {
		var u challenge.NewUser
		t.Run(fmt.Sprintf("%s", tc.Name), func(t *testing.T) {
			req := challenge.MakeRequest(t, "POST", "url", tc.JSON)
			err := myhttp.DecodeRequest(req, &u)
			failure := ((err == nil) && tc.ShouldErr) || ((err != nil) && !tc.ShouldErr)
			if failure {
				t.Error(err)
			}
		})
	}
}

type assertion func(w httptest.ResponseRecorder, t *testing.T)

func TestRespondWithJSON(t *testing.T) {
	testCases := []struct {
		Name      string
		Code      int
		Payload   string
		Assertion assertion
	}{
		{
			Name:    "SetsContentType",
			Code:    http.StatusOK,
			Payload: "test",
			Assertion: func(w httptest.ResponseRecorder, t *testing.T) {
				actual := w.Header().Get("Content-Type")
				if actual != "application/json" {
					t.Errorf("expected application/json; got %s", actual)
				}
			},
		},
		{
			Name:    "SetsStatusCode",
			Code:    http.StatusCreated,
			Payload: "test",
			Assertion: func(w httptest.ResponseRecorder, t *testing.T) {
				actual := w.Code
				if actual != http.StatusCreated {
					t.Errorf("expected %d; got %d", http.StatusCreated, actual)
				}
			},
		},
		{
			Name:    "WritesResponse",
			Code:    http.StatusUnauthorized,
			Payload: "response",
			Assertion: func(w httptest.ResponseRecorder, t *testing.T) {
				res := w.Result()
				defer res.Body.Close()

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("could not read response: %v", err)
				}

				actual := string(bytes.Trim(body, "\x22"))
				if actual != "response" {
					t.Errorf("expected 'response'; got %s", actual)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.Name), func(t *testing.T) {
			w := httptest.NewRecorder()
			myhttp.RespondWithJSON(w, tc.Code, tc.Payload)
			tc.Assertion(*w, t)
		})
	}
}

func TestRespondWithError(t *testing.T) {
	testCases := []struct {
		Name      string
		Code      int
		Message   string
		Assertion assertion
	}{
		{
			Name:    "SetsContentType",
			Code:    http.StatusOK,
			Message: "test",
			Assertion: func(w httptest.ResponseRecorder, t *testing.T) {
				actual := w.Header().Get("Content-Type")
				if actual != "application/json" {
					t.Errorf("expected application/json; got %s", actual)
				}
			},
		},
		{
			Name:    "SetsStatusCode",
			Code:    http.StatusBadRequest,
			Message: "test",
			Assertion: func(w httptest.ResponseRecorder, t *testing.T) {
				actual := w.Code
				if actual != http.StatusBadRequest {
					t.Errorf("expected %d; got %d", http.StatusBadRequest, actual)
				}
			},
		},
		{
			Name:    "SetsErrorMessage",
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Assertion: func(w httptest.ResponseRecorder, t *testing.T) {
				res := w.Result()
				defer res.Body.Close()

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("could not read response: %v", err)
				}

				actual := string(bytes.Trim(body, "\x22"))
				if actual != "unauthorized" {
					t.Errorf("expected 'unauthorized'; got %s", actual)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.Name), func(t *testing.T) {
			w := httptest.NewRecorder()
			myhttp.RespondWithJSON(w, tc.Code, tc.Message)
			tc.Assertion(*w, t)
		})
	}
}
