package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"stonehenge/domain"
	"stonehenge/model"
	"strings"
	"time"
)

// A middleware that listens to the JWT of the request and validates it
func TokenValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/accounts" && r.Method == http.MethodPost {
			next.ServeHTTP(wr, r)
			return
		}
		token := getRequestToken(r)
		accountId, err := domain.GetAccountIdByToken(token)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(wr, model.ErrUnauthorized.Error())
			return
		}

		token, err = domain.CreateToken(*accountId)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(wr, model.ErrUnauthorized.Error())
			return
		}
		assignTokenToResponse(wr, token)

		ctx := context.WithValue(r.Context(), model.ContextAccount, *accountId)
		next.ServeHTTP(wr, r.WithContext(ctx))
	})
}

// Returns the request token present in the request
func getRequestToken(r *http.Request) string {
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

// Assing the authorization a Set Cookie header to a response object
func assignTokenToResponse(rw http.ResponseWriter, token string) {
	rw.Header().Add("Authorization", "Bearer "+token)
	http.SetCookie(rw, &http.Cookie{
		Name:    "access_token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(time.Minute * 15),
	})
}
