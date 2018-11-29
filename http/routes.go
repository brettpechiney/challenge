package http

func (s *server) addRoutes() {
	s.r.Handle("/signup", s.HandleSignupV1()).Methods("POST")
	s.r.Handle("/login", s.HandleAuthorizeV1()).Methods("POST")
	s.r.Handle("/users", Adapt(s.HandleFindAllUsersV1(), CheckAuth("customer", s.signingKey))).
		Methods("GET")
	s.r.Handle("/attestations", Adapt(s.HandleCreateAttestationV1(), CheckAuth("admin", s.signingKey))).
		Methods("POST")
	s.r.Handle("/attestations", Adapt(s.HandleFindAttestationsByUserV1(), CheckAuth("customer", s.signingKey))).
		Queries("fname", "{fname:[a-zA-Z]+}", "lname", "{lname:[a-zA-Z]+}").
		Methods("GET")
}
