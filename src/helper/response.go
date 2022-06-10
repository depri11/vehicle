package helper

import (
	"encoding/json"
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

func (r *Res) Send(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func ResponseJSON(message string, code int, status string, data interface{}) *Res {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	response := Res{
		Meta: meta,
		Data: data,
	}

	return &response
}
