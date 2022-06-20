package vehicle

import (
	"fmt"

	"github.com/depri11/vehicle/src/database/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*models.Vehicles, error) {
	var vehicles models.Vehicles

	err := r.db.Order("id desc").Preload("Images", "vehicle_images.is_primary = true").Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	return &vehicles, nil
}

func (r *repository) Save(vehicle *models.Vehicle) (*models.Vehicle, error) {
	err := r.db.Create(vehicle).Error
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (r *repository) GetID(ID int) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.Preload("Images", "vehicle_images.is_primary = true").First(&vehicle, ID).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Update(vehicle *models.Vehicle) (*models.Vehicle, error) {
	err := r.db.Save(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (r *repository) Delete(ID int) error {
	var vehicle models.Vehicle
	err := r.db.Where("id = ?", ID).Delete(&vehicle).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Popular() (*models.Vehicles, error) {
	var vehicle models.Vehicles
	err := r.db.Order("likes desc").Preload("Images", "vehicle_images.is_primary = true").Find(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Sort(sort string) (*models.Vehicles, error) {
	var vehicle models.Vehicles
	err := r.db.Order(fmt.Sprintf("id %v", sort)).Preload("Images", "vehicle_images.is_primary = true").Find(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *repository) Search(search string) (*models.Vehicles, error) {
	var vehicle *models.Vehicles
	err := r.db.Where("LOWER(name) LIKE ?", "%"+search+"%").Preload("Images", "vehicle_images.is_primary = true").Find(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return vehicle, nil

}
