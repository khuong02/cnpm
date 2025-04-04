package center

import (
	"cnpm/pkg/model"
	"context"
)

type ICenter interface {
	List(ctx context.Context, req model.GetListCenter) ([]model.Center, error)
	Insert(ctx context.Context, req model.UpsertCenter) (model.Center, error)
	GetByID(ctx context.Context, id string) (model.Center, error)
	Update(ctx context.Context, id string, req model.UpsertCenter) (model.Center, error)
}
