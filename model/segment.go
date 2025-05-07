package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerWithTotal struct {
    ID         uint
    Name       string
    Phone      string
    Email      string
    TotalSpent float64
}

func FindSegmentedCustomers(db *gorm.DB, minAmount float64, days int) ([]CustomerWithTotal, error) {
    var result []CustomerWithTotal
    cutoff := time.Now().AddDate(0, 0, -days)

    err := db.Table("customers").
        Select("customers.id, customers.name, customers.phone, customers.email, SUM(purchases.amount) as total_spent").
        Joins("JOIN purchases ON purchases.customer_id = customers.id").
        Where("purchases.purchased_at >= ?", cutoff).
        Group("customers.id").
        Having("SUM(purchases.amount) > ?", minAmount).
        Scan(&result).Error

    return result, err
}