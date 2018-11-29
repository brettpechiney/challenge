package http

func (s *server) addRoutes() {
	s.r.HandleFunc("/signup", s.HandleSignupV1()).Methods("POST")
	s.r.HandleFunc("/login", s.HandleAuthorizeV1()).Methods("POST")
	s.r.HandleFunc("/users", s.HandleFindAllUsersV1()).Methods("GET")
	s.r.HandleFunc("/attestations", s.HandleCreateAttestationV1()).Methods("POST")
	s.r.HandleFunc("/attestations", s.HandleFindAttestationsByUserV1()).
		Queries("fname", "{fname:[a-zA-Z]+}", "lname", "{lname:[a-zA-Z]+}").
		Methods("GET")
}
