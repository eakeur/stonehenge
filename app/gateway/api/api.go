package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"stonehenge/app"
	"stonehenge/app/gateway/api/accounts"
	"stonehenge/app/gateway/api/authentication"
	"stonehenge/app/gateway/api/middlewares"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfers"
)

type Server struct {
	Router *chi.Mux
}

func New(application *app.Application) *Server {

	m := middlewares.NewMiddleware(application.AccessManager)

	aut := authentication.NewController(application.Authentication, application.Logger)
	acc := accounts.NewController(application.Accounts, application.Logger)
	trf := transfers.NewController(application.Transfers, application.Logger)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Group(func(r chi.Router) {
		r.Use(m.Authorization)
		router.Route("/accounts", func(r chi.Router) {
			r.Method("GET", "/", rest.Handler(acc.List))
			r.Method("GET", "/{accountID}/balance", rest.Handler(acc.GetBalance))
		})
		router.Route("/transfers", func(r chi.Router) {
			r.Method("POST", "/", rest.Handler(trf.Create))
			r.Method("GET", "/", rest.Handler(trf.List))
		})
	})
	router.Method("POST", "/login", rest.Handler(aut.Authenticate))
	router.Method("POST", "/accounts", rest.Handler(acc.Create))

	return &Server{
		Router: router,
	}
}
