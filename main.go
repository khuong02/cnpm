package main

import (
	"cnpm/api"
	"cnpm/pkg/database"
	"cnpm/pkg/services/booking"
	"cnpm/pkg/services/center"
	"cnpm/pkg/services/user"
)

func main() {
	client := database.ConnectMongo("mongodb://root:example@localhost:27017")

	route := api.Route{
		BookingSvc: booking.New(client),
		CenterSvc:  center.New(client),
		UserSvc:    user.New(client),
	}

	route.Run()
}
