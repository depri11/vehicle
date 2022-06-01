package vehicle

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Vehicles, error) {
	var vehicles Vehicles

	err := r.db.Find(&vehicles).Error
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
	err := r.db.First(&vehicle, ID).Error
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

func (r *repository) GetImage(VehicleID int) (VehicleImages, error) {
	var image VehicleImages
	err := r.db.Where("vehicle_id = ?", VehicleID).Find(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *repository) CreateImage(image *VehicleImage) (*VehicleImage, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return nil, err
	}

	return image, nil
}
