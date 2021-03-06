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
		var message string
		if err := DecodeRequest(r, &u); err != nil {
			message = err.Error()
		}
		if u.Role != "customer" && len(message) == 0 {
			message = "invalid role"
		}
		if len(message) > 0 {
			RespondWithError(w, http.StatusBadRequest, message)
		} else {
			signup(s, u, w)
		}
	}
}

// HandleRegisterPrivilegedV1 registers a privileged user.
func (s server) HandleRegisterPrivilegedV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u challenge.NewUser
		if err := DecodeRequest(r, &u); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			signup(s, u, w)
		}
	}
}

func signup(s server, u challenge.NewUser, w http.ResponseWriter) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		const Message = "something went wrong"
		RespondWithError(w, http.StatusInternalServerError, Message)
		return
	}
	u.Password = string(hashedPw)
	id, err := s.usrRepo.Insert(&u)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		const Message = "something went wrong"
		RespondWithError(w, http.StatusInternalServerError, Message)
		return
	}
	RespondWithJSON(w, http.StatusCreated, id)
}
