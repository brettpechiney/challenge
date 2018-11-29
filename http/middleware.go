package http

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// An Adapter takes in and returns an http.Handler. It is very useful when
// writing middleware.
type Adapter func(http.Handler) http.Handler

// Adapt wraps an http.Handler with a series of Adapter types, returning
// the result of the first adapter. Note that the adapters are applied in
// reverse of the specified order.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// CheckAuth attempts to inspect the JWT token contained in the request
// header to see if it is valid and contains sufficient privileges.
func CheckAuth(role string, secret string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const AuthErrorMessage = "could not authorize"

			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				RespondWithError(w, http.StatusUnauthorized, AuthErrorMessage)
				return
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if ve, ok := err.(*jwt.ValidationError); ok {
				switch {
				case ve.Errors&jwt.ValidationErrorMalformed != 0:
					RespondWithError(w, http.StatusUnauthorized, "bad token")
				case ve.Errors&jwt.ValidationErrorAudience != 0:
					RespondWithError(w, http.StatusUnauthorized, "AUD validation failed")
				case ve.Errors&jwt.ValidationErrorExpired != 0:
					RespondWithError(w, http.StatusUnauthorized, "EXP validation failed")
				case ve.Errors&jwt.ValidationErrorIssuedAt != 0:
					RespondWithError(w, http.StatusUnauthorized, "IAT validation failed")
				case ve.Errors&jwt.ValidationErrorIssuer != 0:
					RespondWithError(w, http.StatusUnauthorized, "ISS validation failed")
				case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
					RespondWithError(w, http.StatusUnauthorized, "NBF validation failed")
				case ve.Errors&jwt.ValidationErrorId != 0:
					RespondWithError(w, http.StatusUnauthorized, "JTI validation failed")
				case ve.Errors&jwt.ValidationErrorClaimsInvalid != 0:
					RespondWithError(w, http.StatusUnauthorized, "claims validation failed")
				default:
					RespondWithError(w, http.StatusUnauthorized, AuthErrorMessage)
				}
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				RespondWithError(w, http.StatusUnauthorized, "invalid token")
				return
			}

			userRole := claims["role"].(string)
			if userRole != role {
				RespondWithError(w, http.StatusForbidden, "insufficient privileges")
				return
			}

			newTokenString, err := GenerateToken(userRole, secret)
			if err != nil {
				RespondWithError(w, http.StatusInternalServerError, "failed to generate token")
			}
			w.Header().Set("Authorization", newTokenString)

			h.ServeHTTP(w, r)
		})
	}
}
