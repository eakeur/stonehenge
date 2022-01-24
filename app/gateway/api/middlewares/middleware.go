package middlewares

import (
	"net/http"
	"stonehenge/app/core/entities/access"
)

func NewMiddleware(manager access.Manager) Middleware {
	return middleware{am: manager}
}

type Middleware interface {
	Authorization(http.Handler) http.Handler
}

type middleware struct {
	am access.Manager
}
