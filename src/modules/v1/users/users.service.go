package users

import (
	"github.com/depri11/vehicle/src/helper"
)

type Service interface {
	FindAll() (*helper.Res, error)
	FindByEmail(email string) (*helper.Res, error)
	RegisterUser(user *User) (*helper.Res, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (r *service) FindAll() (*helper.Res, error) {
	data, err := r.repository.FindAll()
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (r *service) FindByEmail(email string) (*helper.Res, error) {
	data, err := r.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (r *service) RegisterUser(user *User) (*helper.Res, error) {
	hashPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPass
	data, err := r.repository.Save(user)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}
