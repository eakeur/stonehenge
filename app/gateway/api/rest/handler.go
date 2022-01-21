package rest

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(*http.Request) Response

type Handler struct {
	Func HandlerFunc
}

func (h Handler) Handle(rw http.ResponseWriter, r *http.Request) {

	res := h.Func(r)

	payload, err := json.Marshal(res)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(res.HTTPStatus)
	_, err = rw.Write(payload)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
