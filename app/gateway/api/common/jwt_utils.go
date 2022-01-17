package common

import (
	"net/http"
	"stonehenge/app/core/types/id"
	"time"

	"github.com/golang-jwt/jwt"
)

const JwtKey string = "EDAF12D5D997C58B1962FD8350E8B1C158447B5D1002DABA4F551BC3CD38F236"

const TokenValidMinutes = 10

// CreateToken creates a JWT token string containing the account id of the applicant
func CreateToken(userId id.External) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = TokenDetails{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * TokenValidMinutes).Unix(),
		},
		&userId,
	}
	return t.SignedString([]byte(JwtKey))
}

// ExtractToken extracts data from a JWT Token
func ExtractToken(token string) (*TokenDetails, error) {
	parsed, err := jwt.ParseWithClaims(string(token), &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if parsed.Claims.Valid() != nil {
		return nil, err
	}
	claims := parsed.Claims.(*TokenDetails)
	return claims, nil
}

// AssignToken assigns the authorization and Set Cookie header to a response object
func AssignToken(rw http.ResponseWriter, token string) {
	rw.Header().Add("Authorization", "Bearer "+token)
	http.SetCookie(rw, &http.Cookie{
		Name:    "access_token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(time.Minute * TokenValidMinutes),
	})
}

// TokenDetails is a struct containing the type of data in the parsed token
type TokenDetails struct {
	*jwt.StandardClaims

	// The id of the current account
	AccountId *id.External
}
