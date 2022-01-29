package middlewares

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"stonehenge/app/core/types/logger"
)

func (m middleware) RequestTracer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqID := uuid.NewString()
		ctx := context.WithValue(req.Context(), logger.RequestTracerContextKey, reqID)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
