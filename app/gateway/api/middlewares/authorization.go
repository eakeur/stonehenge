package middlewares

import (
	"net/http"
	"stonehenge/app/gateway/api/rest"
)

func (m middleware) Authorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		const operation = "Controllers.Middleware.Authorization"
		ctx := req.Context()
		rest.Handler(func(r *http.Request) rest.Response {
			token := r.Header.Get("Authorization")
			acc, err := m.am.ExtractAccessFromToken(token)
			if err != nil {
				m.logger.Error(ctx, operation, err.Error())
				return rest.BuildErrorResult(err)
			}
			acc, err = m.am.Create(acc.AccountID)
			if err != nil {
				m.logger.Error(ctx, operation, err.Error())
				return rest.BuildErrorResult(err)
			}
			ctx = m.am.AssignAccessToContext(ctx, acc)
			m.logger.Trace(ctx, operation, "finished operation successfully")
			return rest.Response{}.
				AddHeaders("Authorization", "Bearer "+acc.Token)
		}).ServeHTTP(rw, req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
