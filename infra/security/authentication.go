package security

import (
	m "stonehenge/core/model"
	"time"

	"github.com/golang-jwt/jwt"
)

const JWT_KEY string = "EDAF12D5D997C58B1962FD8350E8B1C158447B5D1002DABA4F551BC3CD38F236"

const TOKEN_VALID_MINUTES = 10

// Creates a JWT token string containing the account id of the applicant
func CreateToken(userId string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = TokenDetails{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * TOKEN_VALID_MINUTES).Unix(),
		},
		&userId,
	}
	return t.SignedString([]byte(JWT_KEY))
}

func ExtractToken(token string) (*TokenDetails, error) {
	parsed, err := jwt.ParseWithClaims(token, &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, m.ErrUnauthorized
	}

	if parsed.Claims.Valid() != nil {
		return nil, m.ErrUnauthorized
	}
	claims := parsed.Claims.(*TokenDetails)
	return claims, nil
}

// A struct containing the type of data in the parsed token
type TokenDetails struct {
	*jwt.StandardClaims

	// The id of the current account
	AccountId *string
}

type key int

const (
	ContextAccount key = iota
)
