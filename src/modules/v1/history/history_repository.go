package history

import (
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Historyss, error) {
	var historys Historyss
	err := r.db.Order("id desc").Preload("Vehicle").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) GetID(ID int) (*Historys, error) {
	var historys Historys
	err := r.db.Preload("Vehicle").First(&historys, ID).Error
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

func (r *repository) Query(sort string) (*Historyss, error) {
	var historys Historyss

	err := r.db.Order(fmt.Sprintf("id %v", sort)).Preload("Vehicle").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}
