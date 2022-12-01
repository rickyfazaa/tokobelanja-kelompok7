package entity

import "time"

type Product struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	Price      int
	Stock      int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Category   Category `gorm:"foreignKey:CategoryID; Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
