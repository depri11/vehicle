package middleware

import (
	"net/http"
	"strconv"
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

		r.Header.Set("user_id", strconv.Itoa(checkToken.Id))
		r.Header.Set("role", checkToken.Role)

		next.ServeHTTP(w, r)
	}
}

func CheckRoleAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		if role != "admin" {
			http.Error(w, "you are not admin", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
