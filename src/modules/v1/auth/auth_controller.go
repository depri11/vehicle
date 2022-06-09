package auth

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/vehicle/src/modules/v1/users"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (c *controller) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user users.User

	json.NewDecoder(r.Body).Decode(&user)

	response := c.service.Login(user)

	json.NewEncoder(w).Encode(response)
}
