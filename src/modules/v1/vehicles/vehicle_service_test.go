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

func TestGetID(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("GetID", 1).Return(&modelMock, nil)
	data, err := service.FindByID(1)

	vehicles := data.Data.(*models.Vehicle)
	assert.Equal(t, 1, vehicles.ID, "Expect id = 1")
	assert.Nil(t, err)
}
