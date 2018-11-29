package http

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/brettpechiney/challenge/challenge"
)

// HandleSignupV1 registers a non-privileged user.
func (s server) HandleSignupV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u challenge.NewUser
		if err := DecodeRequest(r, &u); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		if err != nil {
			const Message = "something went wrong"
			RespondWithError(w, http.StatusInternalServerError, Message)
		}
		u.Password = string(hashedPw)
		id, err := s.usrRepo.Insert(&u)
		if err != nil {
			log.Printf("failed to create user: %v", err)
			const Message = "something went wrong"
			RespondWithError(w, http.StatusInternalServerError, Message)
		}
		RespondWithJSON(w, http.StatusCreated, id)
	}
}

// HandlePrivilegedSignupV1 registers a privileged user.
func (s server) HandlePrivilegedSignupV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement after token
	}
}
