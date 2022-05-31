package users

import (
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Users, error) {
	var users Users

	err := r.db.Find(&users).Error

	if err != nil {
		return nil, errors.New("failed to retrieve data")
	}

	return &users, nil
}

func (r *repository) Save(user *User) (*User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, errors.New("gagal menyimpan data")
	}

	return user, nil
}

func (r *repository) GetUserID(ID int) (*User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	if err != nil {
		return nil, errors.New("failed get user")
	}

	return &user, nil
}

func (r *repository) Update(user *User) (*User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, errors.New("failed update data user")
	}

	return user, nil
}

func (r *repository) Delete(ID int) error {
	var user User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
