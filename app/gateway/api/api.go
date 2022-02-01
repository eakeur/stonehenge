package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"stonehenge/app"
	"stonehenge/app/gateway/api/account"
	"stonehenge/app/gateway/api/authentication"
	"stonehenge/app/gateway/api/middlewares"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfer"
)

type Server struct {
	Router         *chi.Mux
	accounts       account.Controller
	transfers      transfer.Controller
	authentication authentication.Controller
	middlewares    middlewares.Middleware
}

func (s *Server) Serve(address string) error {
	return http.ListenAndServe(address, s.Router)
}

func (s *Server) AssignRoutes() {

	s.Router.Use(s.middlewares.CORS, s.middlewares.RequestTracer)

	s.Router.Route("/accounts", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(s.middlewares.Authorization)
			r.Method("GET", "/", rest.Handler(s.accounts.List))
			r.Method("GET", "/{accountID}/balance", rest.Handler(s.accounts.GetBalance))
		})
		r.Method("POST", "/", rest.Handler(s.accounts.Create))
	})

	s.Router.Route("/transfers", func(r chi.Router) {
		r.Use(s.middlewares.Authorization)
		r.Method("POST", "/", rest.Handler(s.transfers.Create))
		r.Method("GET", "/", rest.Handler(s.transfers.List))
	})

	s.Router.Method("POST", "/login", rest.Handler(s.authentication.Authenticate))
}

func NewServer(application *app.Application) *Server {

	responseBuilder := rest.ResponseBuilder{
		Access: application.AccessManager,
		Logger: application.Logger,
	}

	srv := &Server{
		Router:         chi.NewRouter(),
		accounts:       account.NewController(application.Accounts, responseBuilder),
		transfers:      transfer.NewController(application.Transfers, application.TransfersWorker, responseBuilder),
		authentication: authentication.NewController(application.Authentication, responseBuilder),
		middlewares:    middlewares.NewMiddleware(application.AccessManager, responseBuilder),
	}

	srv.AssignRoutes()
	return srv
}
