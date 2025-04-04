package user

import "go.mongodb.org/mongo-driver/mongo"

type user struct {
	mongoClient *mongo.Client
}

func New(mongoClient *mongo.Client) IUser {
	return &user{
		mongoClient: mongoClient,
	}
}
