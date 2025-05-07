package controller

import (
	"net/http"
	"strconv"

	"github.com/wiroger9595/newretail-homework/presenter"
	"github.com/wiroger9595/newretail-homework/view"

	"github.com/gin-gonic/gin"
)

type SegmentRequest struct {
    MinAmount float64 `json:"min_amount"`
    Days      int     `json:"days"`
    Message   string  `json:"message"` // 支援 {{Name}} {{TotalSpent}}
}

func SendMarketingSMS(c *gin.Context) {
    var req SegmentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // 呼叫 Presenter 邏輯
    results, err := presenter.SendSegmentedSMS(req.MinAmount, req.Days, req.Message)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 回傳 View 結構
    c.JSON(http.StatusOK, view.MarketingResponse{
        SentCount: len(results),
        Customers: results,
        Message:   "SMS sent successfully to " + strconv.Itoa(len(results)) + " customers",
    })
}