package booking

import (
	"cnpm/pkg/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type booking struct {
	mongoClient *mongo.Client
}

func New(mongoClient *mongo.Client) IBooking {
	return &booking{
		mongoClient: mongoClient,
	}
}

func (b booking) Booking(ctx context.Context, req model.BookingRequest) error {
	//TODO implement me
	panic("implement me")
}
