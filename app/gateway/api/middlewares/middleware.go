package middlewares

import (
	"net/http"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/gateway/api/rest"
)

func NewMiddleware(manager access.Manager, builder rest.ResponseBuilder) Middleware {
	return middleware{access: manager, builder: builder}
}

type Middleware interface {
	Authorization(http.Handler) http.Handler
	RequestTracer(http.Handler) http.Handler
}

type middleware struct {
	access  access.Manager
	builder rest.ResponseBuilder
}
