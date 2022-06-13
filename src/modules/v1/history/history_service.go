package history

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/depri11/vehicle/src/helper"
)

type Service interface {
	FindAll() (*helper.Res, error)
	FindByID(id int) (*helper.Res, error)
	Create(user *Historys, r *http.Request) (*helper.Res, error)
	Update(id int, vehicle *Historys, r *http.Request) (*helper.Res, error)
	Delete(id int, r *http.Request) (*helper.Res, error)
	Sort(sort string) (*helper.Res, error)
	Search(search string) (*helper.Res, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() (*helper.Res, error) {
	vehicles, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("success", 200, "OK", vehicles)
	return response, nil
}

func (s *service) FindByID(id int) (*helper.Res, error) {
	vehicle, err := s.repository.GetID(id)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("success", 200, "OK", vehicle)
	return response, nil
}

func (s *service) Create(history *Historys, r *http.Request) (*helper.Res, error) {
	reqUserId := r.Header.Get("user_id")
	setId, err := strconv.Atoi(reqUserId)
	if err != nil {
		return nil, err
	}

	history.UserID = setId

	data, err := s.repository.Save(history)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (s *service) Update(id int, history *Historys, r *http.Request) (*helper.Res, error) {
	data, err := s.repository.GetID(id)
	if err != nil {
		return nil, err
	}

	reqUserId := r.Header.Get("user_id")
	setId, err := strconv.Atoi(reqUserId)
	if err != nil {
		return nil, err
	}

	if setId != data.UserID {
		return nil, err
	}

	data.UserID = history.ID
	data.Name = history.Name
	data.Duration = history.Duration
	data.Prepayment = history.Prepayment
	data.Returned = history.Returned

	res, err := s.repository.Update(data)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", res)
	return response, nil
}

func (s *service) Delete(id int, r *http.Request) (*helper.Res, error) {
	data, err := s.repository.GetID(id)
	if err != nil {
		return nil, err
	}

	reqUserId := r.Header.Get("user_id")
	setId, err := strconv.Atoi(reqUserId)
	if err != nil {
		return nil, err
	}

	if setId != data.UserID {
		return nil, errors.New("access danied")
	}
	err = s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", nil)
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
