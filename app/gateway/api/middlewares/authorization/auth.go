package authorization

import (
	"net/http"
	"stonehenge/app/gateway/api/common"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		token := getToken(r)
		account, err := common.ExtractToken(token)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err = common.CreateToken(*account.AccountId)
		if err != nil {
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		common.AssignToken(wr, token)

		ctx := common.AssignUserToContext(r.Context(), *account.AccountId)
		next.ServeHTTP(wr, r.WithContext(ctx))
	})
}

// getToken returns the request token present in the request
func getToken(r *http.Request) string {
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
