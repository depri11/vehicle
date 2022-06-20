package models

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password,omitempty" validate:"required,min=6"`
	Phone     string    `json:"phone" validate:"required,min=12"`
	Gender    string    `json:"gender" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Nickname  string    `json:"nickname" validate:"required"`
	Birthday  string    `json:"birthday" validate:"required"`
	Role      string    `json:"role" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
