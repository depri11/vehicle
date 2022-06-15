package models

import (
	"time"
)

type Historys struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	UserID     int    `json:"user_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Duration   string `json:"duration" validate:"required"`
	Prepayment string `json:"prepayment" validate:"required"`
	Returned   bool   `json:"returned"`
	User       *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Historyss []Historys
