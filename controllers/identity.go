package controllers

import (
	"encoding/json"
	"net/http"
	"stonehenge/app"
	"stonehenge/core/model"
	"time"
)

type IdentityController struct {
	logins app.IdentityApp
}

func NewIdentityController(logins *app.IdentityApp) IdentityController {
	return IdentityController{
		logins: *logins,
	}
}

// Authenticates the user requesting access to the API. Verifies if the password is correct
// and returns a token
func (lg *IdentityController) Authenticate(rw http.ResponseWriter, r *http.Request) {
	login := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		SendErrorResponse(rw, model.ErrInvalidBody)
		return
	}

	token, err := lg.logins.Authenticate(login.Identity)
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
