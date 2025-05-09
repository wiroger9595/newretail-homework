package route

import (
	"context"
	"fmt"
	"net/http"
	"newretail-homework/presenter"
	"newretail-homework/view"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx = context.Background()

func Coupon(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {

    r.POST("/coupon/claim", func(c *gin.Context) {
        var req view.CouponRequest

        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        userKey := fmt.Sprintf("user_claimed:%d:%d", req.CouponId, req.UserId)
		exists, err := rdb.Exists(ctx, userKey).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			return
		}
		if exists > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You have already claimed this coupon"})
			return
		}

		// 先從 Redis 扣數量
		couponKey := fmt.Sprintf("coupon:%d", req.CouponId)
		remaining, err := rdb.Decr(ctx, couponKey).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			return
		}

        if remaining < 0 {
			// 把 Redis 補回去
			rdb.Incr(ctx, couponKey)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Coupon sold out"})
			return
		}


        couponResponse, _, err := presenter.TryClaimCoupon(db, rdb, ctx, req.UserId, req.CouponId, req.UserLevel)
        if err != nil {
            rdb.Incr(ctx, couponKey)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        rdb.Set(ctx, userKey, true, 0)
		
		c.JSON(http.StatusOK, gin.H{
            "message": "get free coupon success",
			"response":  couponResponse,
		})
    })



}