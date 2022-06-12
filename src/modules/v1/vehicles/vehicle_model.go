package vehicle

import "time"

type Vehicle struct {
	ID          int           `json:"id" gorm:"primariKey"`
	Name        string        `json:"name" validate:"required,max=20"`
	City        string        `json:"city" validate:"required"`
	Available   bool          `json:"available"`
	Prepayment  bool          `json:"prepayment"`
	Capacity    int           `json:"capacity" validate:"required"`
	Type        string        `json:"type" validate:"required"`
	Reservation string        `json:"reservation" validate:"required"`
	Price       string        `json:"price" validate:"required"`
	Likes       int           `json:"likes"`
	Quantity    int           `json:"quantity" validate:"required"`
	Images      *VehicleImage `json:"images"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Vehicles []Vehicle

type VehicleImage struct {
	ID        int    `json:"id" gorm:"primariKey"`
	VehicleID uint   `json:"vehicle_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Path      string `json:"path"`
	IsPrimary bool   `json:"is_primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VehicleImages []VehicleImage
