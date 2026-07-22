package routes

import (
	"dvdrental-api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/films", func(r chi.Router) {
		r.Get("/", handlers.GetFilms)
	})

	return r
}
