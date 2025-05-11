package presenter

import (
	"context"
	"fmt"
	"newretail-homework/view"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


func TryClaimCoupon(db *gorm.DB, rdb *redis.Client, ctx context.Context, userID int, couponID int, userLevel string) ([]view.CouponResponse, map[uint]bool, error) {
    var responses []view.CouponResponse
	resultMap := make(map[uint]bool)

	err := db.Transaction(func(tx *gorm.DB) error {
		res := tx.Exec(`
			UPDATE coupon
			SET quantity = quantity - 1
			WHERE id = ? AND quantity > 0 AND end_time > now() AND coupon_level = ?
		`, couponID, userLevel)

		if res.Error != nil {
			return fmt.Errorf("update coupon failed: %w", res.Error)
		}

		if res.RowsAffected == 0 {
			return fmt.Errorf("coupon already claimed or expired")
		}

		err := tx.Exec(`
			INSERT INTO user_coupon (user_id, coupon_id, claimed_at, status)
			VALUES (?, ?, now(), 'unused')
		`, userID, couponID).Error
		if err != nil {
			return fmt.Errorf("insert user_coupon failed: %w", err)
		}

		responses = append(responses, view.CouponResponse{
			UserId:   userID,
			CouponId: couponID,
		})
		resultMap[uint(couponID)] = true

		redisKey := fmt.Sprintf("user_coupon:%d:%d", userID, couponID)
		if err := rdb.Set(ctx, redisKey, true, 0).Err(); err != nil {
			return fmt.Errorf("failed to update redis user_coupon: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return responses, resultMap, nil
}




