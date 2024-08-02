package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	ID      *int64 `json:"id,omitempty"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	var id *int64
	if payload != nil {
		switch v := payload.(type) {
		case map[string]interface{}:
			if val, ok := v["id"].(int64); ok {
				id = &val
			}
		case struct{ ID int64 }:
			id = &v.ID
		case *struct{ ID int64 }:
			id = &v.ID
		}

		response = ResponseWithData{
			Status:  status,
			Message: message,
			Data:    payload,
			ID:      id,
		}
	} else {
		response = ResponseWithoutData{
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
