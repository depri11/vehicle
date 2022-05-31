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

func (r *repository) FindAll() (Users, error) {
	var users Users

	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, errors.New("failed to retrieve data")
	}

	return users, nil
}
