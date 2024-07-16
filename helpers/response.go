package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithouthData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	if payload != nil {
		response = ResponseWithData{
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = ResponseWithouthData{
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
