package main

import (
	"stonehenge/handler"
	"stonehenge/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Assigns all routes to their handlers
func InitializeServer() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/accounts", func(router chi.Router) {
		router.Use(middlewares.TokenValidatorMiddleware)

		router.Get("/", handler.GetAllAcounts)
		router.Post("/", handler.CreateAccount)
		// Routes with the accountId as parameter
		router.Route("/{accountId}", func(router chi.Router) {
			router.Get("/", handler.GetAccountById)
			router.Get("/balance", handler.GetAccountBalance)
		})
	})

	router.Route("/login", func(router chi.Router) {
		router.Post("/", handler.Authenticate)
	})

	router.Route("/transfers", func(router chi.Router) {
		router.Use(middlewares.TokenValidatorMiddleware)
		router.Get("/", handler.GetAllTransfers)

		router.Post("/", handler.RequestTransfer)
	})

	return router

}
