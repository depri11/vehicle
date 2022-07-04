package interfaces

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
)

type VehicleRepository interface {
	FindAll() (*models.Vehicles, error)
	Save(vehicle *models.Vehicle) (*models.Vehicle, error)
	GetID(ID int) (*models.Vehicle, error)
	GetByType(search string) (*models.Vehicles, error)
	Update(vehicle *models.Vehicle) (*models.Vehicle, error)
	Delete(ID int) error
	Popular() (*models.Vehicles, error)
	Sort(sort string) (*models.Vehicles, error)
	Search(search string) (*models.Vehicles, error)
}

type VehicleService interface {
	FindAll() (*helper.Res, error)
	FindByID(id int) (*helper.Res, error)
	FindByType(id string) (*helper.Res, error)
	Create(user *models.Vehicle) (*helper.Res, error)
	Update(id int, vehicle *models.Vehicle) (*helper.Res, error)
	Delete(id int) (*helper.Res, error)
	Popular() (*helper.Res, error)
	Sort(sort string) (*helper.Res, error)
	Search(search string) (*helper.Res, error)
}
