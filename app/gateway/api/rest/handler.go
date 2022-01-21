package rest

import (
	"encoding/json"
	"net/http"
)

type Handler func(*http.Request) Response

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	res := h(r)

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
