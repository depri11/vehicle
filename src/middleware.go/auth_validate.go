package middleware

import (
	"net/http"
	"strings"

	"github.com/depri11/vehicle/src/helper"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helper.ResponseJSON("invalid header type", 401, "error", nil).Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := helper.CheckToken(token)
		if err != nil {
			helper.ResponseJSON("invalid token", 401, "error", nil).Send(w)
			return
		}

		if !checkToken {
			helper.ResponseJSON("please login again", 401, "error", nil).Send(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
