package users

import (
	"time"

	"github.com/depri11/vehicle/src/modules/v1/history"
)

type User struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	Fullname  string            `json:"fullname"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	Phone     string            `json:"phone"`
	Gender    string            `json:"gender"`
	Address   string            `json:"address"`
	Nickname  string            `json:"nickname"`
	Birthday  string            `json:"birthday"`
	Historys  history.Historyss `json:"historys"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User
