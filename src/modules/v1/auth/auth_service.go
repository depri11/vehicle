package auth

import (
	"github.com/depri11/vehicle/src/helper"
	"github.com/depri11/vehicle/src/modules/v1/users"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type Service interface {
	Login(user users.User) *helper.Res
}

type service struct {
	auth users.Repository
}

func NewService(auth users.Repository) *service {
	return &service{auth}
}

func (s *service) Login(user users.User) *helper.Res {
	data, err := s.auth.GetByEmail(user.Email)
	if err != nil {
		response := helper.ResponseJSON("email your incorrect", 401, "Bad Request", nil)
		return response
	}

	if data.ID == 0 {
		response := helper.ResponseJSON("email your incorrect", 401, "Bad Request", nil)
		return response
	}

	if !helper.CheckPassword(data.Password, user.Password) {
		response := helper.ResponseJSON("Internal Server Error", 500, "error", nil)
		return response
	}

	token := helper.NewToken(data.Email)
	tokens, err := token.Create()
	if err != nil {
		response := helper.ResponseJSON("Internal Server Error", 500, "error", nil)
		return response
	}

	response := helper.ResponseJSON("Success", 200, "OK", tokenResponse{Token: tokens})
	return response
}
