package controllers

import (
	"encoding/json"
	"net/http"
	"stonehenge/app"
	"stonehenge/core/model"
	"stonehenge/infra/security"
)

type TransferController struct {
	transfers app.TransferApp
}

func NewTransferController(transfers *app.TransferApp) TransferController {
	return TransferController{
		transfers: *transfers,
	}
}

// Gets all transfers of this actual account
func (tc *TransferController) GetAllTransfers(rw http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(security.ContextAccount).(string)
	//parameter that sets if the transfers being requested are the ones made to the applicant or by them
	toMe := r.URL.Query().Get("toMe")

	transfers, err := tc.transfers.GetAll(accountId, toMe == "true")
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
func (tc *TransferController) RequestTransfer(rw http.ResponseWriter, r *http.Request) {
	transfer := model.Transfer{}
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		SendErrorResponse(rw, model.ErrInvalidBody)
		return
	}

	transfer.AccountOriginId = r.Context().Value(security.ContextAccount).(string)
	id, err := tc.transfers.Transact(&transfer)
	if err != nil {
		SendErrorResponse(rw, err)
		return
	}

	rw.Header().Add("Location", r.URL.Path+"/"+*id)
	SendResponse(rw, nil, http.StatusCreated)
}
