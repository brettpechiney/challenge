package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/brettpechiney/challenge/challenge"
)

// HandleCreateAttestationV1 creates an attestation for a user.
func (s *server) HandleCreateAttestationV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a challenge.NewAttestation
		if err := DecodeRequest(r, &a); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		id, err := s.attestRepo.Insert(&a)
		if err != nil {
			log.Printf("failed to create attestation: %v", err)
			const Message = "something went wrong"
			RespondWithError(w, http.StatusInternalServerError, Message)
		}
		RespondWithJSON(w, http.StatusCreated, id)
	}
}

// HandleFindAttestationsByUserV1 finds all attestations for a user.
func (s *server) HandleFindAttestationsByUserV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ua, err := s.attestRepo.FindByUser(vars["fname"], vars["lname"])
		if err != nil {
			log.Printf("failed to retrieve attestations: %v", err)
			const Message = "something went wrong"
			RespondWithError(w, http.StatusInternalServerError, Message)
		}
		RespondWithJSON(w, http.StatusOK, ua)
	}
}
