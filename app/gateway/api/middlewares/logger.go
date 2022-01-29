package middlewares

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"stonehenge/app/core/types/logger"
)

func (m middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), logger.TraceKey, uuid.NewString())
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
