package vehicle

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*Vehicles, error)
	Save(vehicle *Vehicle) (*Vehicle, error)
	GetID(ID int) (*Vehicle, error)
	Update(vehicle *Vehicle) (*Vehicle, error)
	Delete(ID int) error
	Popular() (*Vehicles, error)
	Sort(sort string) (*Vehicles, error)
	Search(search string) (*Vehicles, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Vehicles, error) {
	var vehicles Vehicles

	err := r.db.Order("id desc").Preload("Images", "vehicle_images.is_primary = true").Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	return &vehicles, nil
}

func (r *repository) Save(vehicle *Vehicle) (*Vehicle, error) {
	err := r.db.Create(vehicle).Error
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (r *repository) GetID(ID int) (*Vehicle, error) {
	var vehicle Vehicle
	err := r.db.Preload("Images", "vehicle_images.is_primary = true").First(&vehicle, ID).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Update(vehicle *Vehicle) (*Vehicle, error) {
	err := r.db.Save(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (r *repository) Delete(ID int) error {
	var vehicle Vehicle
	err := r.db.Where("id = ?", ID).Delete(&vehicle).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Popular() (*Vehicles, error) {
	var vehicle Vehicles
	err := r.db.Order("likes desc").Preload("Images", "vehicle_images.is_primary = true").Find(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Sort(sort string) (*Vehicles, error) {
	var vehicle Vehicles
	err := r.db.Order(fmt.Sprintf("id %v", sort)).Preload("Images", "vehicle_images.is_primary = true").Find(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Search(search string) (*Vehicles, error) {
	var vehicle Vehicles
	err := r.db.Where("LOWER(name) LIKE ?", "%"+search+"%").Preload("Images", "vehicle_images.is_primary = true").Find(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil

}
