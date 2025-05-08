package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"newretail-homework/presenter"
	"newretail-homework/view"
)

func Customer(r *gin.Engine, db *gorm.DB) {
    r.POST("/customer/send", func(c *gin.Context) {

		var req view.CustomerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
		}

        minAmount := req.Amount
        sinceDays := req.SinceDays
        customerResponse, _, err := presenter.FindTargetCustomers(db, sinceDays, minAmount)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

		smsTemplate := "親愛的 {{name}}，您在最近的消費金額為 {{amount}} 元，感謝您的支持！"
		sentMessages  := presenter.SendMarketingSMS(customerResponse, smsTemplate)

		c.JSON(http.StatusOK, gin.H{
			"sms_logs":  sentMessages,
		})
    })
}