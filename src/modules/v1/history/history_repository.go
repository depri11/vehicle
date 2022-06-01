package history

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*[]Historys, error) {
	var data []Historys

	err := r.db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetID(ID int) (*Historys, error) {
	var historys Historys
	err := r.db.First(&historys, ID).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) Delete(ID int) error {
	var history Historys
	err := r.db.Where("id = ?", ID).Delete(&history).Error
	if err != nil {
		return err
	}

	return nil
}
