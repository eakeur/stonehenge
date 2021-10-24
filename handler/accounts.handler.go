package handler

import (
	"encoding/json"
	"net/http"
	"stonehenge/domain"
	model "stonehenge/model"
	"time"

	"github.com/go-chi/chi/v5"
)

// Gets all accounts existing
func GetAllAcounts(rw http.ResponseWriter, r *http.Request) {
	accounts, err := domain.GetAllAccounts()
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
func GetAccountById(rw http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	if len(accountId) > 0 {
		acc, err := domain.GetAccountById(accountId)
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
		SendErrorResponse(rw, model.ErrAccountInvalid)
	}

}

// Gets the balance of the account of the id passed
func GetAccountBalance(rw http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	if len(accountId) > 0 {
		bal, err := domain.GetAccountBalance(accountId)
		if err != nil {
			SendErrorResponse(rw, err)
			return
		}
		responseBody, err := json.Marshal(map[string]interface{}{"balance": bal})
		if err != nil {
			SendErrorResponse(rw, err)
			return
		}
		SendResponse(rw, responseBody, http.StatusOK)
	} else {
		SendErrorResponse(rw, model.ErrAccountInvalid)
	}
}

// Creates an account based on the data parsed
func CreateAccount(rw http.ResponseWriter, r *http.Request) {
	account := model.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		SendErrorResponse(rw, model.ErrInvalidBody)
		return
	}

	res, err := domain.AddNewAccount(account)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}

	rw.Header().Add("Location", r.URL.Path+"/"+*res)
	token, err := domain.Authenticate(model.Login{
		Cpf: account.Cpf, Secret: account.Secret,
	})
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

	SendResponse(rw, nil, http.StatusCreated)
}
