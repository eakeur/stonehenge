package main

import (
	"stonehenge/handler"
	"stonehenge/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeServer() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/accounts", routeAccounts)
	router.Route("/login", routeAuthentication)
	router.Route("/transfers", routeTransfers)

	return router

}

func routeAccounts(router chi.Router) {
	router.Use(middlewares.TokenValidatorMiddleware)

	router.Get("/", handler.GetAllAcounts)
	router.Post("/", handler.CreateAccount)
	// Routes with the accountId as parameter
	router.Route("/{accountId}", func(router chi.Router) {
		router.Get("/", handler.GetAccountById)
		router.Get("/balance", handler.GetAccountBalance)
	})

}

func routeAuthentication(router chi.Router) {
	router.Post("/", handler.Authenticate)
}

func routeTransfers(router chi.Router) {
	router.Use(middlewares.TokenValidatorMiddleware)
	router.Get("/", handler.GetAllTransfers)

	router.Post("/", handler.RequestTransfer)
}
