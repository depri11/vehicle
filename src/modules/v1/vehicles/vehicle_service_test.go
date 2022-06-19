package vehicle

import (
	"testing"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/modules/v1/vehicles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var modelMock = models.Vehicle{
	ID: 1,
}

func TestFindById(t testing.T) {
	var repos = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repos}

	repos.Mock.On("FindByUserId", 1).Return(modelMock, nil)
	data, err := service.FindByID(1)

	orders := data.Data.(models.Vehicle)
	assert.Equal(t, 1, orders, "Expect id = 1")
	assert.Nil(t, err)
}
