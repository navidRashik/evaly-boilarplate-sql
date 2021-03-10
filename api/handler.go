package api

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func brandsRouter(ctrl *BrandsController) http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/", ctrl.ListBrand)
		r.Post("/", ctrl.AddBrand)
	})

	return h
}

func healthRouter(ctrl *SystemController) http.Handler {
	log.Println("healthRouter")
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/api", ctrl.apiCheck)
	})
	return h
}
