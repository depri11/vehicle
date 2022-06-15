package users

import (
	"github.com/depri11/vehicle/src/database/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*models.Users, error) {
	var users models.Users

	err := r.db.Order("id desc").Find(&users).Error
	// err := r.db.Order("id desc").Preload("Historys").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *repository) Save(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) GetUserID(ID int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(user *models.User) (*models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Delete(ID uint) error {
	var user models.User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
