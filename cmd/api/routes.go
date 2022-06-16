package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		router.Get("/healthcheck", app.healthcheckHandler)
		router.Post("/movies", app.createMovieHandler)
		router.Get("/movies/{id}", app.showMovieHandler)
	})

	return router
}
