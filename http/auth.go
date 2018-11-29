package http

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/brettpechiney/challenge/challenge"
)

// HandleAuthorizeV1 attempts to authorize an alleged registered user
// of the application.
func (s server) HandleAuthorizeV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u challenge.UserLogin
		if err := DecodeRequest(r, &u); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		pw, err := s.usrRepo.GetPassword(u.Username)
		if err != nil {
			const Message = "something went wrong"
			RespondWithError(w, http.StatusInternalServerError, Message)
		}
		err = bcrypt.CompareHashAndPassword([]byte(pw), []byte(u.Password))
		if err != nil {
			const Message = "could not authenticate"
			RespondWithError(w, http.StatusUnauthorized, Message)
		}
		RespondWithJSON(w, http.StatusOK, "Login")
		// TODO: return token
	}
}
