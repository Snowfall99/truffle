package server

import (
	"truffle.io/handlers"
	"truffle.io/model"
)

func (s *Server) setupRoutes(bfts []model.BFT) {
	handlers.Health(s.mux)
	handlers.FrontPage(s.mux, s.log, bfts)
	for _, bft := range bfts {
		handlers.NormalPage(s.mux, s.log, bft)
	}
}
