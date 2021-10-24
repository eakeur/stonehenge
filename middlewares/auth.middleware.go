package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"stonehenge/model"
	"stonehenge/shared"
	"strings"
)

// A middleware that listens to the JWT of the request and validates it
func TokenValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/accounts" && r.Method == http.MethodPost {
			next.ServeHTTP(wr, r)
			return
		}
		token := GetRequestToken(r)
		accountId, err := shared.GetAccountIdByToken(token)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(wr, model.ErrUnauthorized.Error())
			return
		}

		token, err = shared.CreateToken(*accountId)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(wr, model.ErrUnauthorized.Error())
			return
		}
		shared.AssignTokenToResponse(wr, token)

		ctx := context.WithValue(r.Context(), shared.ContextAccount, *accountId)
		next.ServeHTTP(wr, r.WithContext(ctx))
	})
}

// Returns the request token present in the request
func GetRequestToken(r *http.Request) string {
	jwtToken := r.Header.Get("Authorization")
	if strings.Trim(jwtToken, " ") == "" {
		cookies := r.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				jwtToken = cookie.Value
				break
			}
		}
	}

	return jwtToken
}
