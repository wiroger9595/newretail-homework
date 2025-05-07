package models

import "time"

type CouponStatus string

const (
	Unused  CouponStatus = "unused"
	Used    CouponStatus = "used"
	Expired CouponStatus = "expired"
)

type UserCoupon struct {
	ID        uint         
	UserID    uint         
	CouponID  uint         
	Status    CouponStatus 
	ClaimedAt time.Time
	UsedAt    *time.Time 
}