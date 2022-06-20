package users

import (
	"testing"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/modules/v1/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var modelMocks = models.Users{
	models.User{
		ID: 1,
	},
	models.User{
		ID: 2,
	},
}

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("FindAll").Return(&modelMocks, nil)
	data, err := service.FindAll()

	Users := data.Data.(*models.Users)

	expectedModelMocks := &models.Users{
		models.User{ID: 1},
		models.User{ID: 2},
	}

	assert.Equal(t, expectedModelMocks, Users, "Expect id = 1 & 2")
	assert.Nil(t, err)
}

var modelMockGetID = models.User{
	ID: 1,
}

// func TestGetUserID(t *testing.T) {
// 	var repo = mocks.RepoMock{Mock: mock.Mock{}}
// 	var service = service{&repo}

// 	repo.Mock.On("GetUserID", 1).Return(&modelMockGetID, nil)
// 	data, err := service.FindByID(1)

// 	users := data.Data.(*models.User)
// 	assert.Equal(t, 1, users.ID, "Expect id = 1")
// 	assert.Nil(t, err)
// }

var modelMockSave = models.User{
	ID:       1,
	Fullname: "dev",
	Email:    "dev@gmail.com",
	Password: "123456",
	Phone:    "081234567890",
	Gender:   "Man",
	Address:  "Jl",
	Nickname: "dev",
	Birthday: "2-9-1990",
	Role:     "user",
}

func TestSave(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("Save", &modelMockSave).Return(&modelMockSave, nil)
	data, err := service.RegisterUser(&modelMockSave)

	Users := data.Data.(*models.User)
	assert.Equal(t, "dev", Users.Fullname, "Expect name = Mobil")
	assert.Nil(t, err)
}
