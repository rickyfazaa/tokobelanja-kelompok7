package entity

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	FullName  string
	Email     string `gorm:"unique"`
	Password  string
	Role      string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
