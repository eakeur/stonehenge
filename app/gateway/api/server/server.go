package server

import (
	"net/http"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/gateway/api/accounts"
	"stonehenge/app/gateway/api/middlewares/authorization"
	"stonehenge/app/gateway/api/transfers"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func New(workspaces *WorkspaceWrapper, accessFactory access.Factory) *Server {
	router := chi.NewRouter()

	authMiddleware := func(h http.Handler) http.Handler {
		return authorization.Middleware(h, accessFactory)
	}

	accountsController := accounts.New(workspaces.Accounts)

	router.Route("/accounts", func(router chi.Router) {
		router.Use(authMiddleware)
		router.Get("/", accountsController.List)
		router.Post("/", accountsController.Create)
		router.Route("/{accountId}", func(router chi.Router) {
			router.Get("/balance", accountsController.GetBalance)
		})
	})

	router.Post("/login", accountsController.Authenticate)

	router.Route("/transfers", func(router chi.Router) {
		transfersController := transfers.New(workspaces.Transfers)

		router.Use(authMiddleware)
		router.Get("/", transfersController.List)
		router.Post("/", transfersController.Create)
	})

	return &Server{
		Router: router,
	}
}

func TokenMiddleware()
