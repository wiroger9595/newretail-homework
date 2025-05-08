package model

import "time"


type CouponType string

const (
	Discount CouponType = "D"
	Fill     CouponType = "F"
)

type Coupon struct {
	ID        uint      
	Name      string
	Type      CouponType 
	Value     float64
	Quantity  int
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time 
}