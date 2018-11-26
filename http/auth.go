package http

import "net/http"

// HandleAuthorizeV1 attempts to authorize an alleged registered user
// of the application.
func HandleAuthorizeV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondWithJSON(w, http.StatusOK, "Login")
		// TODO: implement after token
	}
}
