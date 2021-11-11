package main

import (
	"stonehenge/app"
	"stonehenge/controllers"
	"stonehenge/infra/persistence"
	"stonehenge/middlewares"

	"github.com/go-chi/chi/v5"
)

type StonehengeServer struct {
	accounts controllers.AccountController

	transfers controllers.TransferController

	logins controllers.IdentityController

	router *chi.Mux
}

func NewStonehengeServer(workspace *persistence.Workspace) *StonehengeServer {
	accounts := app.NewAccountApp(workspace.Accounts)
	logins := app.NewIdentityApp(workspace.Identity)
	transfers := app.NewTransferApp(workspace.Transfers, workspace.Accounts)
	return &StonehengeServer{
		accounts:  controllers.NewAccountController(&accounts, &logins),
		logins:    controllers.NewIdentityController(&logins),
		transfers: controllers.NewTransferController(&transfers),
		router:    chi.NewRouter(),
	}
}

func (s *StonehengeServer) MapControllers() *chi.Mux {

	s.router.Route("/accounts", func(router chi.Router) {
		router.Use(middlewares.TokenValidatorMiddleware)

		router.Get("/", s.accounts.GetAllAcounts)
		router.Post("/", s.accounts.CreateAccount)
		// Routes with the accountId as parameter
		router.Route("/{accountId}", func(router chi.Router) {
			router.Get("/", s.accounts.GetAccountById)
			router.Get("/balance", s.accounts.GetAccountBalance)
		})
	})

	s.router.Route("/login", func(router chi.Router) {
		router.Post("/", s.logins.Authenticate)
	})

	s.router.Route("/transfers", func(router chi.Router) {
		router.Use(middlewares.TokenValidatorMiddleware)
		router.Get("/", s.transfers.GetAllTransfers)

		router.Post("/", s.transfers.RequestTransfer)
	})

	return s.router

}
