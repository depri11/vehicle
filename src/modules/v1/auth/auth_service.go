package auth

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
	"github.com/depri11/vehicle/src/interfaces"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type service struct {
	auth interfaces.UserRepository
}

func NewService(auth interfaces.UserRepository) *service {
	return &service{auth}
}

func (s *service) Login(user models.User) *helper.Res {
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

	token := helper.NewToken(int(data.ID), data.Email, data.Role)
	tokens, err := token.Create()
	if err != nil {
		response := helper.ResponseJSON("Internal Server Error", 500, "error", nil)
		return response
	}

	response := helper.ResponseJSON("Success", 200, "OK", tokenResponse{Token: tokens})
	return response
}
