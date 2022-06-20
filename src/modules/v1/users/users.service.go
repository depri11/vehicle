package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
	"github.com/depri11/vehicle/src/interfaces"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() (*helper.Res, error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (s *service) FindByEmail(email string) (*helper.Res, error) {
	data, err := s.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (s *service) FindByID(id int) (*helper.Res, error) {
	data, err := s.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}

	data.Password = ""

	response := helper.ResponseJSON("Success", 200, "OK", &data)
	return response, nil
}

func (r *service) RegisterUser(user *models.User) (*helper.Res, error) {
	hashPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPass
	user.Role = "user"
	data, err := r.repository.Save(user)
	if err != nil {
		return nil, err
	}

	data.Password = ""

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (s *service) Delete(id int, r *http.Request) (*helper.Res, error) {
	data, err := s.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}

	reqUserId := r.Header.Get("user_id")
	setId, err := strconv.ParseUint(reqUserId, 0, 0)
	if err != nil {
		return nil, err
	}

	if uint(setId) != data.ID {
		return nil, errors.New("access danied")
	}

	err = s.repository.Delete(data.ID)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", nil)
	return response, nil
}

func (s *service) UpdateUser(id int, user *models.User, r *http.Request) (*helper.Res, error) {
	data, err := s.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}

	reqUserId := r.Header.Get("user_id")
	setId, err := strconv.ParseUint(reqUserId, 0, 0)
	if err != nil {
		return nil, err
	}

	if uint(setId) != data.ID {
		return nil, errors.New("access danied")
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

	res, err := s.repository.Update(data)
	data.Password = ""
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", res)
	return response, nil

}

func (s *service) Sort(sort string) (*helper.Res, error) {
	vehicles, err := s.repository.Sort(sort)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("success", 200, "OK", vehicles)
	return response, nil
}

func (s *service) Search(search string) (*helper.Res, error) {
	vehicles, err := s.repository.Search(search)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("success", 200, "OK", vehicles)
	return response, nil
}
