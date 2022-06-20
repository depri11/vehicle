package auth

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/interfaces"
)

type controller struct {
	service interfaces.AuthService
}

func NewController(service interfaces.AuthService) *controller {
	return &controller{service}
}

func (c *controller) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	res := c.service.Login(user)

	res.Send(w)
}
