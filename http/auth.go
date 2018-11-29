package http

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/brettpechiney/challenge/challenge"
)

// HandleAuthorizeV1 attempts to authorize an alleged registered user
// of the application.
func (s server) HandleAuthorizeV1() http.HandlerFunc {
	type response struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		const InternalErrorMessage = "something went wrong"

		// Verify credentials
		var ul challenge.UserLogin
		if err := DecodeRequest(r, &ul); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		// TODO: add specific error checks
		pw, err := s.usrRepo.GetPassword(ul.Username)
		if err != nil {
			const Message = "could not authenticate"
			RespondWithError(w, http.StatusUnauthorized, Message)
		}
		err = bcrypt.CompareHashAndPassword([]byte(pw), []byte(ul.Password))
		if err != nil {
			const Message = "could not authenticate"
			RespondWithError(w, http.StatusUnauthorized, Message)
		}

		// Generate JWT
		u, err := s.usrRepo.FindByUsername(ul.Username)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, InternalErrorMessage)
		}

		// TODO: use goroutine
		u.LastLogin = time.Now()
		_, err = s.usrRepo.Update(u)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, InternalErrorMessage)
		}

		// TODO: use goroutine
		tokenString, err := GenerateToken(u.Role, s.signingKey)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, InternalErrorMessage)
		}

		res := response{tokenString}
		RespondWithJSON(w, http.StatusOK, res)
	}
}
