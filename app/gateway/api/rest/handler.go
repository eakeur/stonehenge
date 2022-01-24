package rest

import (
	"encoding/json"
	"net/http"
)

type Handler func(*http.Request) Response

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	res := h(r)

	if res.headers != nil {
		header := rw.Header()
		for k, v := range res.headers {
			header.Add(k, v)
		}
	}

	if res.HTTPStatus != 0 {
		rw.WriteHeader(res.HTTPStatus)
	}

	if res.Error != nil || res.Content != nil {
		payload, err := json.Marshal(res)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = rw.Write(payload)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
