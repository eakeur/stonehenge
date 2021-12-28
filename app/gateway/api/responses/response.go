package responses

import (
	"encoding/json"
	"net/http"
)

func WriteSuccessfulJSON(w http.ResponseWriter, status int, content interface{}) error {
	if content != nil {
		body, err := json.Marshal(content)
		if err != nil {
			return err
		}
		_, err = w.Write(body)
		if err != nil {
			return err
		}
	}
	w.WriteHeader(status)
	return nil
}

func WriteErrorResponse(w http.ResponseWriter, status int, message error) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(message.Error()))
	if err != nil {
		return
	}
}
