package handlers

import (
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"

	"truffle.io/model"
	"truffle.io/views"
)

func FrontPage(mux chi.Router, bfts []model.BFT) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_ = views.FrontPage(bfts).Render(w)
	})
}

func NormalPage(mux chi.Router, bft model.BFT) {
	mux.Get(path.Join("/", bft.Name), func(w http.ResponseWriter, r *http.Request) {
		_ = views.NormalPage(bft).Render(w)
	})
}
