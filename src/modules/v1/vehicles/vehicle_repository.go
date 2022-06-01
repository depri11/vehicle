package vehicle

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Vehicles, error) {
	var vehicles Vehicles

	err := r.db.Preload("Images", "vehicle_images.is_primary = 1").Find(&vehicles).Error
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
	err := r.db.Preload("Images", "vehicle_images.is_primary = 1").First(&vehicle, ID).Error
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
