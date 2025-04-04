package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetListCenter struct {
	Page    int       `json:"page"`
	Limit   int       `json:"limit"`
	Date    time.Time `json:"date"`
	Address string    `json:"address" binding:"required"`
	Block   string    `json:"block"`
	Room    string    `json:"room"`
}

type UpsertCenter struct {
	Block   string `json:"block" binding:"required"`
	Address string `json:"address" binding:"required"`
	Room    string `json:"room" binding:"required"`
}

type Center struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Block     string             `json:"block" binding:"required"`
	Address   string             `json:"address" binding:"required"`
	Room      string             `json:"room" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (Center) CollectionName() string {
	return "center"
}
