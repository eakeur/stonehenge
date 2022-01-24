package middlewares

import (
	"net/http"
	"stonehenge/app/gateway/api/rest"
)

func (m middleware) Authorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		rest.Handler(func (r *http.Request) rest.Response {
			token := r.Header.Get("Authorization")
			acc, err := m.am.ExtractAccessFromToken(token)
			if err != nil {
				return rest.BuildErrorResult(err)
			}
			acc, err = m.am.Create(acc.AccountID)
			if err != nil {
				return rest.BuildErrorResult(err)
			}
			ctx = m.am.AssignAccessToContext(ctx, acc)
			return rest.Response{}.
				AddHeaders("Authorization", "Bearer " + acc.Token)
		}).ServeHTTP(rw, req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}