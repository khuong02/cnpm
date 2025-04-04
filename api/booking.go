package api

import (
	"cnpm/pkg/model"
	"github.com/gin-gonic/gin"
)

func (r Route) booking(c *gin.Context) {
	var payload model.BookingRequest
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Sai định dạng dữ liệu đầu vào!",
		})
		return
	}
	if err := r.BookingSvc.Booking(c, payload); err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Đặt chỗ thất bại!",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Booking API",
	})
}
