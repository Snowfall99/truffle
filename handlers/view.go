package handlers

import (
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"truffle.io/model"
	"truffle.io/views"
)

func FrontPage(mux chi.Router, log *zap.Logger, bfts []model.BFT) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request 8081")
		_ = views.FrontPage(bfts).Render(w)
	})
}

func NormalPage(mux chi.Router, log *zap.Logger, bft model.BFT) {
	mux.Get(path.Join("/", bft.Name), func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request 8081/" + bft.Name)
		_ = views.NormalPage(bft).Render(w)
	})
}
