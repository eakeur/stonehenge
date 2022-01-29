package rest

import (
	"context"
	"encoding/json"
	"net/http"
	logger "stonehenge/app/core/types/logger"
)

type Handler func(*http.Request) Response

func (handler Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	ch := make(chan Response)
	go func() { ch <- handler(r) }()
	res := <-ch

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

func GetRequestID(ctx context.Context) string {
	return ctx.Value(logger.RequestTracerContextKey).(string)
}
