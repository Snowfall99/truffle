package server

import "truffle.io/handlers"

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
}
