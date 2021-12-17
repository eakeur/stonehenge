package server

import (
	"stonehenge/app/gateway/api/accounts"
	"stonehenge/app/gateway/api/middlewares/authorization"
	"stonehenge/app/gateway/api/transfers"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func New(workspaces *WorkspaceWrapper) *Server {
	router := chi.NewRouter()
	router.Route("/accounts", func(router chi.Router) {

		controller := accounts.New(workspaces.Accounts)

		router.Use(authorization.Middleware)
		router.Get("/", controller.List)
		router.Post("/", controller.Create)
		router.Route("/{accountId}", func(router chi.Router) {
			router.Get("/balance", controller.GetBalance)
		})
	})

	router.Route("/transfers", func(router chi.Router) {
		controller := transfers.New(workspaces.Transfers)

		router.Use(authorization.Middleware)
		router.Get("/", controller.List)
		router.Post("/", controller.Create)
	})

	return &Server{
		Router:    router,
	}
}
