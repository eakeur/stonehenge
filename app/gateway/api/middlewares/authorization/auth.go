package authorization

import (
	"net/http"
	"stonehenge/app/core/entities/access"
)

func Middleware(next http.Handler, factory access.Factory) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		access, err := factory.ExtractAccessFromToken(token)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}
		access, err = factory.Create(access.AccountID)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		wr.Header().Add("Authorization", "Bearer "+access.Token)
		ctx := factory.AssignAccessToContext(r.Context(), access)
		next.ServeHTTP(wr, r.WithContext(ctx))
	})
}
