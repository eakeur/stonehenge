package handler

import (
	"encoding/json"
	"net/http"
	"stonehenge/domain"
	model "stonehenge/model"
	"time"
)

// Authenticates the user requesting access to the API. Verifies if the password is correct
// and returns a token
func Authenticate(rw http.ResponseWriter, r *http.Request) {
	login := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		SendErrorResponse(rw, model.ErrInvalidBody)
		return
	}

	token, err := domain.Authenticate(login)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}

	rw.Header().Add("Authorization", "Bearer "+*token)
	http.SetCookie(rw, &http.Cookie{
		Name:    "access_token",
		Value:   *token,
		Path:    "/",
		Expires: time.Now().Add(time.Minute * 15),
	})
}
