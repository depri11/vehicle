package history

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*Historyss, error)
	GetID(ID int) (*Historys, error)
	Save(history *Historys) (*Historys, error)
	Delete(ID int) error
	Query(sort string) (*Historyss, error)
	Search(search string) (*Historyss, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Historyss, error) {
	var historys Historyss
	err := r.db.Order("id desc").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) GetID(ID int) (*Historys, error) {
	var historys Historys
	err := r.db.Preload("User").First(&historys, ID).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) Save(history *Historys) (*Historys, error) {
	err := r.db.Create(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
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

	err := r.db.Order(fmt.Sprintf("id %v", sort)).Preload("User").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil
}

func (r *repository) Search(search string) (*Historyss, error) {
	var historys Historyss
	err := r.db.Where("LOWER(name) LIKE ?", "%"+search+"%").Preload("User").Find(&historys).Error
	if err != nil {
		return nil, err
	}

	return &historys, nil

}
