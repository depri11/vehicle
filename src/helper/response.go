package helper

import (
	"net/http"
)

type Res struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseJSON(w http.ResponseWriter, message string, code int, status string, data interface{}) Res {
	w.Header().Set("Content-Type", "application/json")
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	response := Res{
		Meta: meta,
		Data: data,
	}

	return response
}
