package history

import (
	"time"

	"github.com/depri11/vehicle/src/modules/v1/users"
)

type Historys struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	UserID     int         `json:"user_id"`
	Name       string      `json:"name"`
	Duration   string      `json:"duration"`
	Prepayment string      `json:"prepayment"`
	Returned   bool        `json:"returned"`
	User       *users.User `json:"user"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateHistorys struct {
	ID         int        `json:"id" gorm:"primaryKey"`
	UserID     int        `json:"user_id"`
	Name       string     `json:"name"`
	Duration   string     `json:"duration"`
	Prepayment string     `json:"prepayment"`
	Returned   bool       `json:"returned"`
	User       users.User `json:"user"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Historyss []Historys
