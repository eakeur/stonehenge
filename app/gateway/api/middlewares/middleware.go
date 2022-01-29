package middlewares

import (
	"net/http"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/types/logger"
)

func NewMiddleware(manager access.Manager, lg logger.Logger) Middleware {
	return middleware{am: manager, logger: lg}
}

type Middleware interface {
	Authorization(http.Handler) http.Handler
	Logger(http.Handler) http.Handler
}

type middleware struct {
	am access.Manager
	logger logger.Logger
}
