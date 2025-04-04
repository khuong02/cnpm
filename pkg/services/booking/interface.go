package booking

import (
	"cnpm/pkg/model"
	"context"
)

type IBooking interface {
	Booking(ctx context.Context, req model.BookingRequest) error
}
