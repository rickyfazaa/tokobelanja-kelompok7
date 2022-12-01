package entity

import "time"

type Category struct {
	ID                int `gorm:"primaryKey"`
	Type              string
	SoldProductAmount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
