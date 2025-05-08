package presenter

import (
	"fmt"
	"log"
	"newretail-homework/view"
	"time"

	"gorm.io/gorm"
)


func ClaimCouponWithRetry(db *gorm.DB, userID int, couponID int, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        err, _, _ := TryClaimCoupon(db, userID, couponID)
        if err == nil {
            return nil // 成功
        }

        

        log.Printf("第 %d 次嘗試失敗：%v，重試中...\n", i+1, err)
        time.Sleep(100 * time.Millisecond) // 可用 Exponential Backoff
    }

    return fmt.Errorf("領取失敗：超過最大重試次數")
}

func TryClaimCoupon(db *gorm.DB, userID int, couponID int) ([]view.CouponResponse, map[uint]float64, error) {
    var responses []view.CouponResponse
	resultMap := make(map[uint]float64)

	err := db.Transaction(func(tx *gorm.DB) error {
		res := tx.Exec(`
			UPDATE coupon
			SET quantity = quantity - 1
			WHERE id = ? AND quantity > 0 AND end_time > now()
		`, couponID)

		if res.Error != nil {
			return fmt.Errorf("update coupon failed: %w", res.Error)
		}

		if res.RowsAffected == 0 {
			return fmt.Errorf("coupon already claimed or expired")
		}

		err := tx.Exec(`
			INSERT INTO user_coupon (user_id, coupon_id, claimed_at, status)
			VALUES (?, ?, now(), 'used')
		`, userID, couponID).Error
		if err != nil {
			return fmt.Errorf("insert user_coupon failed: %w", err)
		}

		// 你也可以這裡查詢資料、組合返回結果
		responses = append(responses, view.CouponResponse{
			UserId:   userID,
			CouponId: couponID,
		})
		resultMap[uint(couponID)] = 0.85

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return responses, resultMap, nil
}




