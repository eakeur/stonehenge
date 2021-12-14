package controllers

import (
	"encoding/json"
	"net/http"
	"stonehenge/app"
	"stonehenge/core/model"
	"stonehenge/infra/security"
	"time"

	"github.com/go-chi/chi/v5"
)

type AccountController struct {
	accounts app.AccountApp
	logins   app.IdentityApp
}

func NewAccountController(accounts *app.AccountApp, logins *app.IdentityApp) AccountController {
	return AccountController{
		accounts: *accounts, logins: *logins,
	}
}

// Gets all accounts existing
func (ac *AccountController) GetAllAcounts(rw http.ResponseWriter, r *http.Request) {
	accounts, err := ac.accounts.GetAll()
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}
	responseBody, err := json.Marshal(accounts)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}
	SendResponse(rw, responseBody, http.StatusOK)

}

// Gets the account by id
func (ac *AccountController) GetAccountById(rw http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	if len(accountId) == 0 {
		SendErrorResponse(rw, model.ErrAccountInvalid)
		return
	}

	if accountId == r.Context().Value(security.ContextAccount).(string) {
		acc, err := ac.accounts.GetById(accountId)
		if err != nil {
			SendErrorResponse(rw, err)
			return
		}
		responseBody, err := json.Marshal(acc)
		if err != nil {
			SendErrorResponse(rw, err)
			return
		}
		SendResponse(rw, responseBody, http.StatusOK)
	} else {
		SendErrorResponse(rw, model.ErrForbidden)
	}

}

// Gets the balance of the account of the id passed
func (ac *AccountController) GetAccountBalance(rw http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	if len(accountId) == 0 {
		SendErrorResponse(rw, model.ErrAccountInvalid)
		return
	}

	if accountId == r.Context().Value(security.ContextAccount).(string) {
		bal, err := ac.accounts.GetBalanceById(accountId)
		if err != nil {
			SendErrorResponse(rw, err)
			return
		}
		responseBody, err := json.Marshal(bal)
		if err != nil {
			SendErrorResponse(rw, err)
			return
		}
		SendResponse(rw, responseBody, http.StatusOK)
	} else {
		SendErrorResponse(rw, model.ErrForbidden)
	}
}

// Creates an account based on the data parsed
func (ac *AccountController) CreateAccount(rw http.ResponseWriter, r *http.Request) {
	account := model.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		SendErrorResponse(rw, model.ErrInvalidBody)
		return
	}

	res, err := ac.accounts.Add(&account)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}

	rw.Header().Add("Location", r.URL.Path+"/"+*res)
	token, err := security.CreateToken(*res)
	if err != nil {
		SendErrorResponse(rw, model.ErrUnauthorized)
	}

	rw.Header().Add("Authorization", "Bearer "+token)
	http.SetCookie(rw, &http.Cookie{
		Name:    "access_token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(time.Minute * 15),
	})

	SendResponse(rw, nil, http.StatusCreated)
}
