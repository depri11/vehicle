package history

import (
	"time"
)

type Historys struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	Duration   string `json:"duration"`
	Prepayment string `json:"prepayment"`
	Returned   bool   `json:"returned"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Historyss []Historys
