package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*Users, error)
	Save(user *User) (*User, error)
	GetUserID(ID int) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) (*User, error)
	Delete(ID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() (*Users, error) {
	var users Users

	err := r.db.Order("id desc").Find(&users).Error
	// err := r.db.Order("id desc").Preload("Historys").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *repository) Save(user *User) (*User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) GetUserID(ID int) (*User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(user *User) (*User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Delete(ID uint) error {
	var user User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
