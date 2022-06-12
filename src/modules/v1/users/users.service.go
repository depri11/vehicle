package users

import (
	"github.com/depri11/vehicle/src/helper"
)

type Service interface {
	FindAll() (*helper.Res, error)
	FindByEmail(email string) (*helper.Res, error)
	FindByID(id int) (*helper.Res, error)
	RegisterUser(user *User) (*helper.Res, error)
	UpdateUser(id int, user *User) (*helper.Res, error)
	Delete(id int) (*helper.Res, error)
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

func (r *service) FindByID(id int) (*helper.Res, error) {
	data, err := r.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}

	data.Password = ""

	response := helper.ResponseJSON("Success", 200, "OK", &data)
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

	data.Password = ""

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (r *service) Delete(id int) (*helper.Res, error) {
	data, err := r.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}

	err = r.repository.Delete(data.ID)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", nil)
	return response, nil
}

func (r *service) UpdateUser(id int, user *User) (*helper.Res, error) {
	data, err := r.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}
	
	data.Fullname = user.Fullname
	data.Address = user.Address
	data.Birthday = user.Birthday
	data.Email = user.Email
	data.Gender = user.Gender
	data.Phone = user.Phone
	data.Nickname = user.Nickname
	if user.Password != "" {
		hashPass, err := helper.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		data.Password = hashPass
	}

	res, err := r.repository.Update(data)
	data.Password = ""
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", res)
	return response, nil

}
