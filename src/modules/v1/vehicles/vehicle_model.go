package vehicle

import "time"

type Vehicle struct {
	ID          uint   `json:"id" gorm:"primariKey"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Available   bool   `json:"available"`
	Prepayment  bool   `json:"prepayment"`
	Capacity    int    `json:"capacity"`
	Type        string `json:"type"`
	Reservation string `json:"reservation"`
	Price       string `json:"price"`
	Likes       int    `json:"like"`
	Quantity    int    `json:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Vehicles []Vehicle
