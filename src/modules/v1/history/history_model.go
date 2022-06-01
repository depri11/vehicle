package history

import (
	"time"

	vehicle "github.com/depri11/vehicle/src/modules/v1/vehicles"
)

type Historys struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	VehicleID int             `json:"vehicle_id"`
	Return    bool            `json:"return"`
	Vehicle   vehicle.Vehicle `json:"vehicle"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Historyss []Historys
