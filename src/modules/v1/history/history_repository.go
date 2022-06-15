package history

import (
	"fmt"

	"github.com/depri11/vehicle/src/database/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*models.Historyss, error)
	GetID(ID int) (*models.Historys, error)
	Save(history *models.Historys) (*models.Historys, error)
	Update(history *models.Historys) (*models.Historys, error)
	Delete(ID int) error
	Sort(sort string) (*models.Historyss, error)
	Search(search string) (*models.Historyss, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() (*models.Historyss, error) {
	var historys models.Historyss
	err := r.db.Order("id desc").Preload("User").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) GetID(ID int) (*models.Historys, error) {
	var historys models.Historys
	err := r.db.Preload("User").First(&historys, ID).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) Save(history *models.Historys) (*models.Historys, error) {
	err := r.db.Create(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (r *repository) Update(history *models.Historys) (*models.Historys, error) {
	err := r.db.Save(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (r *repository) Delete(ID int) error {
	var history models.Historys
	err := r.db.Where("id = ?", ID).Delete(&history).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Sort(sort string) (*models.Historyss, error) {
	var historys models.Historyss

	err := r.db.Order(fmt.Sprintf("id %v", sort)).Preload("User").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) Search(search string) (*models.Historyss, error) {
	var historys models.Historyss
	err := r.db.Where("LOWER(name) LIKE ?", "%"+search+"%").Preload("User").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil

}
