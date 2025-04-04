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

	ginR.POST("/api/sso", r.SSO)
	ginR.Run(":8080")
}
