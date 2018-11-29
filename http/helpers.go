package http

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

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

// GenerateToken generates a new JWT token.
func GenerateToken(role string, key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := challenge.AppClaims{
		// TODO: enhance claims
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "challenge",
			IssuedAt:  time.Now().Unix(),
		},
		Role: role,
	}
	token.Claims = claims

	return token.SignedString([]byte(key))
}
