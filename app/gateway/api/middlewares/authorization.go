package middlewares

import (
	"net/http"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/gateway/api/rest"
	"strings"
)

func (m middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		const operation = "Controllers.Middleware.Authorization"
		ctx := req.Context()
		failed := false
		rest.Handler(func(r *http.Request) rest.Response {
			token := r.Header.Get("Authorization")
			acc, err := m.access.ExtractAccessFromToken(strings.Replace(token, "Bearer ", "", 1))
			if err != nil {
				failed = true
				err = erring.Wrap(err, operation)
				return m.builder.BuildErrorResult(err).WithErrorLog(ctx)
			}
			acc, err = m.access.Create(acc.AccountID)
			if err != nil {
				failed = true
				err = erring.Wrap(err, operation, erring.AdditionalData{Key: "actor", Value: acc.AccountID.String()})
				return m.builder.BuildErrorResult(err).WithErrorLog(ctx)
			}
			ctx = m.access.AssignAccessToContext(ctx, acc)
			return rest.Response{}.
				AddHeaders("Authorization", "Bearer "+acc.Token)
		}).ServeHTTP(rw, req)
		if !failed {
			next.ServeHTTP(rw, req.WithContext(ctx))
		}
	})
}
