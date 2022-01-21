package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"stonehenge/app"
	"stonehenge/app/gateway/api/accounts"
	"stonehenge/app/gateway/api/middlewares/authorization"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfers"
)

type Server struct {
	Router *chi.Mux
}

func New(application *app.Application) *Server {
	authMiddleware := func(h http.Handler) http.Handler {
		return authorization.Middleware(h, application.AccessManager)
	}

	acc := accounts.NewController(application.Accounts)
	trf := transfers.NewController(application.Transfers)

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/accounts", func(r chi.Router) {
			r.Get("/", rest.Handler{Func: acc.List}.Handle)
			r.Get("/{accountID}/balance", rest.Handler{Func: acc.GetBalance}.Handle)
		})
		r.Route("/transfers", func(r chi.Router) {
			r.Post("/", rest.Handler{Func: trf.Create}.Handle)
			r.Get("/", rest.Handler{Func: trf.List}.Handle)
		})
	})
	router.Post("/login", rest.Handler{Func: acc.Authenticate}.Handle)
	router.Post("/accounts", rest.Handler{Func: acc.Create}.Handle)

	return &Server{
		Router: router,
	}
}
