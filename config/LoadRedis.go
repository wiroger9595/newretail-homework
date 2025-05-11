package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)





func LoadCouponsToRedis(rdb *redis.Client, db *gorm.DB) error {
	ctx := context.Background()

	var coupons []struct {
		ID       int
		Quantity int
	}
	
	if err := db.Raw("SELECT id, quantity FROM coupon WHERE end_time > NOW()").Scan(&coupons).Error; err != nil {
		return err
	}

	for _, c := range coupons {
		key := fmt.Sprintf("coupon:%d", c.ID)
		err := rdb.Set(ctx, key, c.Quantity, 0).Err()
		if err != nil {
			return fmt.Errorf("failed to set coupon %d to redis: %w", c.ID, err)
		}
	}

	return nil
}