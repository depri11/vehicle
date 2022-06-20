package mocks

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindAll() (*models.Vehicles, error) {
	args := pr.Mock.Called()
	return args.Get(0).(*models.Vehicles), args.Error(1)
}

func (pr *RepoMock) Save(vehicle *models.Vehicle) (*models.Vehicle, error) {
	args := pr.Mock.Called(vehicle)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (pr *RepoMock) GetID(ID int) (*models.Vehicle, error) {
	args := pr.Mock.Called(ID)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (pr *RepoMock) Update(vehicle *models.Vehicle) (*models.Vehicle, error) {
	args := pr.Mock.Called(vehicle)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (pr *RepoMock) Delete(ID int) error {
	args := pr.Mock.Called(ID)
	return args.Error(1)
}

func (pr *RepoMock) Popular() (*models.Vehicles, error) {
	args := pr.Mock.Called()
	return args.Get(0).(*models.Vehicles), args.Error(1)
}

func (pr *RepoMock) Sort(sort string) (*models.Vehicles, error) {
	args := pr.Mock.Called(sort)
	return args.Get(0).(*models.Vehicles), args.Error(1)
}

func (pr *RepoMock) Search(search string) (*models.Vehicles, error) {
	args := pr.Mock.Called(search)
	return args.Get(0).(*models.Vehicles), args.Error(1)
}
