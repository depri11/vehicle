package vehicle

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
	"github.com/depri11/vehicle/src/interfaces"
)

type service struct {
	repository interfaces.VehicleRepository
}

func NewService(repository interfaces.VehicleRepository) *service {
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

func (s *service) Create(vehicle *models.Vehicle) (*helper.Res, error) {
	data, err := s.repository.Save(vehicle)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}

func (s *service) Update(id int, vehicle *models.Vehicle) (*helper.Res, error) {
	data, err := s.repository.GetID(id)
	if err != nil {
		return nil, err
	}

	data.Name = vehicle.Name
	data.City = vehicle.City
	data.Available = vehicle.Available
	data.Prepayment = vehicle.Prepayment
	data.Capacity = vehicle.Capacity
	data.Type = vehicle.Type
	data.Reservation = vehicle.Reservation
	data.Price = vehicle.Price
	data.Quantity = vehicle.Quantity

	res, err := s.repository.Update(data)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", res)
	return response, nil
}

func (s *service) Delete(id int) (*helper.Res, error) {
	_, err := s.repository.GetID(id)
	if err != nil {
		return nil, err
	}

	err = s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", nil)
	return response, nil
}

func (s *service) Popular() (*helper.Res, error) {
	vehicles, err := s.repository.Popular()
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("success", 200, "OK", vehicles)
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
