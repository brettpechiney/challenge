package http

import (
	"github.com/gorilla/mux"

	"github.com/brettpechiney/challenge/challenge"
)

// Server provides an interface to set up an HTTP server.
type Server interface {
	addRoutes()
	Start()
}

// server holds and decouples dependencies.
type server struct {
	r          *mux.Router
	usrRepo    challenge.UserRepository
	attestRepo challenge.AttestationRepository
}

// NewServer returns a new instance of a server object.
func NewServer(userRepo challenge.UserRepository, attestRepo challenge.AttestationRepository) Server {
	router := mux.NewRouter()
	return &server{
		r:          router,
		usrRepo:    userRepo,
		attestRepo: attestRepo,
	}
}
