package users

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Nickname  string `json:"nickname"`
	Birthday  string `json:"birthday"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User
