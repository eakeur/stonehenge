package api

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"io/ioutil"
	"net/http"

	"stonehenge/app"
	"stonehenge/app/gateway/api/account"
	"stonehenge/app/gateway/api/authentication"
	"stonehenge/app/gateway/api/middlewares"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/gateway/api/transfer"
	_ "stonehenge/docs"
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

	s.Router.Route("/api/v1", func(r chi.Router) {
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

		s.Router.Method("GET", "/swagger/{*}", httpSwagger.WrapHandler)

		s.Router.Get("/swagger/swagger.json", func(rw http.ResponseWriter, _ *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			doc, _ := ioutil.ReadFile("docs/swagger.json")
			rw.Write(doc)
		})

		s.Router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "/swagger/index.html", http.StatusSeeOther)
		})
	})

}

func NewServer(application *app.Application) *Server {

	responseBuilder := rest.ResponseBuilder{
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
