package history

import (
	"time"

	"github.com/depri11/vehicle/src/modules/v1/users"
)

type Historys struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	UserID     int         `json:"user_id" validate:"required"`
	Name       string      `json:"name" validate:"required"`
	Duration   string      `json:"duration" validate:"required"`
	Prepayment string      `json:"prepayment" validate:"required"`
	Returned   bool        `json:"returned"`
	User       *users.User `json:"user"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Historyss []Historys
