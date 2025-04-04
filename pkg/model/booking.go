package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingRequest struct {
	CenterID  primitive.ObjectID `json:"center_id" binding:"required"`
	StartTime string             `json:"start_time" binding:"required"`
	EndTime   string             `json:"end_time" binding:"required"`
}

type Booking struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	CenterID primitive.ObjectID `json:"center_id" bson:"center_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
}

func (Booking) CollectionName() string {
	return "booking"
}
