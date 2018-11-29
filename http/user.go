package http

import (
	"log"
	"net/http"
)

// HandleFindAllUsersV1 finds all users if the requesting user
// is an administrator.
func (s server) HandleFindAllUsersV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.usrRepo.FindAll()
		if err != nil {
			log.Printf("failed to retrieve users: %v", err)
			const Message = "something went wrong"
			RespondWithError(w, http.StatusInternalServerError, Message)
		}
		RespondWithJSON(w, http.StatusOK, users)
	}
}
