package inputport

import (
	"context"

	"golang-trainning-frontend/pkg/domain/collection"
	"golang-trainning-frontend/pkg/domain/entity"
	"golang-trainning-frontend/pkg/usecase/input"
)

type WomanUsecase interface {
	GetList(ctx context.Context, i input.GetWomanListInput) (collection.Collection[entity.WomanEntity], error)
	GetStoreWomanList(ctx context.Context, i input.GetStoreWomanListInput) (collection.Collection[entity.WomanEntity], error)
	GetDetail(ctx context.Context, i input.GetWomanDetailInput) (entity.WomanEntity, error)
}
