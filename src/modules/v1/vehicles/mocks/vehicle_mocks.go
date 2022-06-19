package mocks

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindByUserId(id int) (*models.Vehicle, error) {
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (pr *RepoMock) Add(data *models.Vehicle) (*models.Vehicle, error) {
	args := pr.Mock.Called(data)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}
