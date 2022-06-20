package mocks

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindAll() (*models.Users, error) {
	args := pr.Mock.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr *RepoMock) Save(User *models.User) (*models.User, error) {
	args := pr.Mock.Called(User)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) GetUserID(ID int) (*models.User, error) {
	args := pr.Mock.Called(ID)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) GetByEmail(email string) (*models.User, error) {
	args := pr.Mock.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) Update(User *models.User) (*models.User, error) {
	args := pr.Mock.Called(User)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) Delete(ID uint) error {
	args := pr.Mock.Called(ID)
	return args.Error(0)
}

func (pr *RepoMock) Search(search string) (*models.Users, error) {
	args := pr.Mock.Called(search)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr *RepoMock) Sort(sort string) (*models.Users, error) {
	args := pr.Mock.Called(sort)
	return args.Get(0).(*models.Users), args.Error(1)
}
