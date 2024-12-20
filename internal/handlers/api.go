package handlers

import (
	"github.com/Karthik150502/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capitalized function mnme enables the us to import the function in other modules, functions with un-capitalized function names, can only be used inside the package.
func Handler(r *chi.Mux) {
	// Parent middleware
	r.Use(chimiddle.StripSlashes) //Removes any trailing slashes in the endpoint.

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authotization)
		router.Get("/coins", GetCoinBalance)
	})
}
