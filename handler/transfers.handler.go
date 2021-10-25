package handler

import (
	"encoding/json"
	"net/http"
	"stonehenge/domain"
	model "stonehenge/model"
)

// Gets all transfers of this actual account
func GetAllTransfers(rw http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(model.ContextAccount).(string)
	//parameter that sets if the transfers being requested are the ones made to the applicant or by them
	toMe := r.URL.Query().Get("toMe")

	transfers, err := domain.GetAllTransfers(accountId, toMe == "true")
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}
	responseBody, err := json.Marshal(transfers)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}
	SendResponse(rw, responseBody, http.StatusOK)
}

// Transfers an amount of money from an account to another
func RequestTransfer(rw http.ResponseWriter, r *http.Request) {
	transfer := model.Transfer{}
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		SendErrorResponse(rw, model.ErrInvalidBody)
		return
	}

	transfer.AccountOriginId = r.Context().Value(model.ContextAccount).(string)
	id, err := domain.TransferMoney(transfer)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}

	rw.Header().Add("Location", r.URL.Path+"/"+*id)
	SendResponse(rw, nil, http.StatusCreated)
}
