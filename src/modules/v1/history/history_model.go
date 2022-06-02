package history

import (
	"time"

	"github.com/depri11/vehicle/src/modules/v1/users"
	vehicle "github.com/depri11/vehicle/src/modules/v1/vehicles"
)

type Historys struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	UserID    int             `json:"user_id"`
	VehicleID int             `json:"vehicle_id"`
	Return    bool            `json:"return"`
	User      users.User      `json:"user"`
	Vehicle   vehicle.Vehicle `json:"vehicle"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Historyss []Historys
