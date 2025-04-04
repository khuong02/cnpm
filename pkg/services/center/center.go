package center

import (
	"cnpm/pkg/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type center struct {
	mongoClient *mongo.Client
}

func New(mongoClient *mongo.Client) ICenter {
	return &center{
		mongoClient: mongoClient,
	}
}

func (c center) List(ctx context.Context, req model.GetListCenter) ([]model.Center, error) {
	//TODO implement me
	panic("implement me")
}

func (c center) Insert(ctx context.Context, req model.UpsertCenter) (model.Center, error) {
	//TODO implement me
	panic("implement me")
}

func (c center) GetByID(ctx context.Context, id string) (model.Center, error) {
	//TODO implement me
	panic("implement me")
}

func (c center) Update(ctx context.Context, id string, req model.UpsertCenter) (model.Center, error) {
	//TODO implement me
	panic("implement me")
}
