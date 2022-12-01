package entity

import "time"

type TransactionHistory struct {
	ID         int `gorm:"primaryKey"`
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product `gorm:"foreignKey:ProductID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User       User    `gorm:"foreignKey:UserID;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
