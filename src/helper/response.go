package helper

import (
	"encoding/json"
	"net/http"
)

// type Response struct {
// 	Meta Meta        `json:"meta"`
// 	Data interface{} `json:"data"`
// }

// type Meta struct {
// 	Message string `json:"message"`
// 	Status  int    `json:"status"`
// }

// func ResponseJSON(message string, code int, status string, data interface{}) Response {
// 	meta := Meta{
// 		Message: message,
// 		Code:    code,
// 		Status:  status,
// 	}

// 	response := Response{
// 		Meta: meta,
// 		Data: data,
// 	}

// 	return response
// }

func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)	
	w.Write([]byte(response))
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJSON(w, code, map[string]string{"error": message})
}
