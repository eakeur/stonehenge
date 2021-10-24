package shared

import (
	"net/http"
	model "stonehenge/model"
	"time"

	"github.com/golang-jwt/jwt"
)

// Retrieves the account id by the authenticated token
func GetAccountIdByToken(token string) (*string, error) {
	parsed, err := retrieveToken(token)
	if err != nil {
		return nil, model.ErrUnauthorized
	}

	if parsed.Claims.Valid() != nil {
		return nil, model.ErrUnauthorized
	}
	claims := parsed.Claims.(*TokenDetails)
	return &claims.AccountId, nil
}

func retrieveToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})
}

func CreateToken(userId string) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("HS256"))

	t.Claims = TokenDetails{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * TOKEN_VALID_MINUTES).Unix(),
		},
		userId,
	}

	return t.SignedString([]byte(JWT_KEY))
}

func AssignTokenToResponse(rw http.ResponseWriter, token string) {
	rw.Header().Add("Authorization", "Bearer "+token)
	http.SetCookie(rw, &http.Cookie{
		Name:    "access_token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(time.Minute * 15),
	})
}

// A struct containing the type of data in the parsed token
type TokenDetails struct {
	*jwt.StandardClaims

	// The id of the current account
	AccountId string
}
