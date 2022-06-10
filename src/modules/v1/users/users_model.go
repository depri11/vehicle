package users

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Nickname  string    `json:"nickname"`
	Birthday  string    `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

type UserInput struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	Nickname string `json:"nickname"`
	Birthday string `json:"birthday"`
}
