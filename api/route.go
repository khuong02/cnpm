package api

import (
	"cnpm/pkg/services/booking"
	"cnpm/pkg/services/center"
	"cnpm/pkg/services/user"
	"github.com/gin-gonic/gin"
)

type Route struct {
	BookingSvc booking.IBooking
	CenterSvc  center.ICenter
	UserSvc    user.IUser
}

func (r Route) Run() {
	ginR := gin.Default()

	// User
	ginR.POST("/api/sso", r.sso)

	// booking
	ginR.POST("/api/booking", r.booking)

	// center
	ginR.GET("/api/center/:id", r.getByID)
	ginR.GET("/api/centers", r.list)
	ginR.POST("/api/centers", r.insert)
	ginR.PUT("/api/centers/:id", r.update)

	ginR.Run(":8080")
}
