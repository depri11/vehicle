package vehicle

import (
	"testing"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/modules/v1/vehicles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var modelMocks = models.Vehicles{
	models.Vehicle{
		ID: 1,
	},
	models.Vehicle{
		ID: 2,
	},
}

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("FindAll").Return(&modelMocks, nil)
	data, err := service.FindAll()

	vehicles := data.Data.(*models.Vehicles)

	expectedModelMocks := &models.Vehicles{
		models.Vehicle{ID: 1},
		models.Vehicle{ID: 2},
	}

	assert.Equal(t, expectedModelMocks, vehicles, "Expect id = 1 & 2")
	assert.Nil(t, err)
}

var modelMockGetID = models.Vehicle{
	ID: 1,
}

func TestGetID(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("GetID", 1).Return(&modelMockGetID, nil)
	data, err := service.FindByID(1)

	vehicles := data.Data.(*models.Vehicle)
	assert.Equal(t, 1, vehicles.ID, "Expect id = 1")
	assert.Nil(t, err)
}

var modelMockSave = models.Vehicle{
	ID:          2,
	Name:        "Mobil",
	City:        "Indonesia",
	Available:   true,
	Prepayment:  false,
	Capacity:    2,
	Type:        "Bike",
	Reservation: "20-05-2022",
	Price:       "Rp1.000.000",
	Likes:       5,
	Quantity:    2,
}

func TestSave(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("Save", &modelMockSave).Return(&modelMockSave, nil)
	data, err := service.Create(&modelMockSave)

	vehicles := data.Data.(*models.Vehicle)
	assert.Equal(t, "Mobil", vehicles.Name, "Expect name = Mobil")
	assert.Nil(t, err)
}

var modelMockUpdate = models.Vehicle{
	ID:   2,
	Name: "Mobil",
}

func TestUpdate(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("GetID", 2).Return(&modelMockUpdate, nil)
	repo.Mock.On("Update", &modelMockUpdate).Return(&modelMockUpdate, nil)

	modelMockAfterUpdate := models.Vehicle{
		Name: "Motor",
	}

	data, err := service.Update(2, &modelMockAfterUpdate)

	vehicles := data.Data.(*models.Vehicle)
	assert.Equal(t, modelMockAfterUpdate.Name, vehicles.Name, "Expect Name = Motor")
	assert.Nil(t, err)
}

var PopularModelMocks = models.Vehicles{
	models.Vehicle{
		ID:    1,
		Name:  "Mobil",
		Likes: 5,
	},
	models.Vehicle{
		ID:    2,
		Name:  "Mobil",
		Likes: 7,
	},
	models.Vehicle{
		ID:    3,
		Name:  "Mobil",
		Likes: 8,
	},
}

func TestPopular(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	popular := models.Vehicles{
		models.Vehicle{
			ID:    3,
			Name:  "Mobil",
			Likes: 8,
		},
		models.Vehicle{
			ID:    2,
			Name:  "Mobil",
			Likes: 7,
		},
		models.Vehicle{
			ID:    1,
			Name:  "Mobil",
			Likes: 5,
		},
	}

	repo.Mock.On("Popular").Return(&popular, nil)
	data, err := service.Popular()

	expected := &models.Vehicles{
		models.Vehicle{
			ID:    3,
			Name:  "Mobil",
			Likes: 8,
		},
		models.Vehicle{
			ID:    2,
			Name:  "Mobil",
			Likes: 7,
		},
		models.Vehicle{
			ID:    1,
			Name:  "Mobil",
			Likes: 5,
		},
	}

	vehicles := data.Data.(*models.Vehicles)
	assert.Equal(t, expected, vehicles)
	assert.Nil(t, err)
}

// func TestSort(t *testing.T) {
// 	var repo = mocks.RepoMock{Mock: mock.Mock{}}
// 	var service = service{&repo}

// 	repo.Mock.On("GetID", 1).Return(&modelMock, nil)
// 	data, err := service.FindByID(1)

// 	vehicles := data.Data.(*models.Vehicle)
// 	assert.Equal(t, 1, vehicles.ID, "Expect id = 1")
// 	assert.Nil(t, err)
// }

var SearchmodelMocks = models.Vehicles{
	models.Vehicle{
		ID:   1,
		Name: "Mobil",
	},
	models.Vehicle{
		ID:   2,
		Name: "Mobil",
	},
}

func TestSearch(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = service{&repo}

	repo.Mock.On("Search", "Mobil").Return(&SearchmodelMocks, nil)
	data, err := service.Search("Mobil")

	expected := &models.Vehicles{
		models.Vehicle{
			ID:   1,
			Name: "Mobil",
		},
		models.Vehicle{
			ID:   2,
			Name: "Mobil",
		},
	}

	vehicles := data.Data.(*models.Vehicles)
	assert.Equal(t, expected, vehicles, "Expect Name = Mobil")
	assert.Nil(t, err)
}
