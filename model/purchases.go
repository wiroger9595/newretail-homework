package models

import "time"

type Purchases struct {
	ID            uint      `gorm:"primaryKey"`
	Customer_Id   uint    `gorm:"size:100"`
	Amount        float64    `gorm:"size:100"`
	PurchasedAt   time.Time `gorm:"autoCreateTime"`
}