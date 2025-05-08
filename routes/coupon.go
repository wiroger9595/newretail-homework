package route

import (
	"net/http"
	"newretail-homework/presenter"
	"newretail-homework/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Coupon(r *gin.Engine, db *gorm.DB) {

    r.POST("/coupon/claim", func(c *gin.Context) {
        var req view.CouponRequest

        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }


        couponResponse, _, err := presenter.TryClaimCoupon(db, req.UserId, req.CouponId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        
		



		c.JSON(http.StatusOK, gin.H{
			"response":  couponResponse,
		})
    })



}